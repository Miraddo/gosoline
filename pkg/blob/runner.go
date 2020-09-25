package blob

import (
	"context"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/kernel"
	"github.com/applike/gosoline/pkg/mon"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/hashicorp/go-multierror"
	"io"
	"sync"
)

const (
	metricName      = "BlobBatchRunner"
	operationRead   = "Read"
	operationWrite  = "Write"
	operationCopy   = "Copy"
	operationDelete = "Delete"
)

var br = struct {
	sync.Mutex
	instance *BatchRunner
}{}

func ProvideBatchRunner() *BatchRunner {
	br.Lock()
	defer br.Unlock()

	if br.instance != nil {
		return br.instance
	}

	br.instance = &BatchRunner{}

	return br.instance
}

type BatchRunner struct {
	kernel.ForegroundModule
	kernel.ServiceStage

	logger mon.Logger
	config *BatchRunnerConfig
	metric mon.MetricWriter
	client s3iface.S3API
	read   chan *Object
	write  chan *Object
	copy   chan *CopyObject
	delete chan *Object
}

type BatchRunnerConfig struct {
	ReaderRunnerCount int `cfg:"reader_runner_count" default:"10"`
	WriterRunnerCount int `cfg:"writer_runner_count" default:"10"`
	CopyRunnerCount   int `cfg:"copy_runner_count" default:"10"`
	DeleteRunnerCount int `cfg:"delete_runner_count" default:"10"`
}

func (r *BatchRunner) Boot(config cfg.Config, logger mon.Logger) error {
	defaultMetrics := getDefaultRunnerMetrics()

	cc := &BatchRunnerConfig{}
	config.UnmarshalKey("blob", cc)

	r.logger = logger
	r.config = cc
	r.client = ProvideS3Client(config)
	r.metric = mon.NewMetricDaemonWriter(defaultMetrics...)
	r.read = make(chan *Object, cc.ReaderRunnerCount)
	r.write = make(chan *Object, cc.WriterRunnerCount)
	r.copy = make(chan *CopyObject, cc.CopyRunnerCount)
	r.delete = make(chan *Object, cc.DeleteRunnerCount)

	return nil
}

func (r *BatchRunner) Run(ctx context.Context) error {
	for i := 0; i < r.config.ReaderRunnerCount; i++ {
		go r.executeRead()
	}

	for i := 0; i < r.config.WriterRunnerCount; i++ {
		go r.executeWrite()
	}

	for i := 0; i < r.config.CopyRunnerCount; i++ {
		go r.executeCopy()
	}

	for i := 0; i < r.config.DeleteRunnerCount; i++ {
		go r.executeDelete()
	}

	<-ctx.Done()

	return nil
}

func (r *BatchRunner) executeRead() {
	for object := range r.read {
		var body io.ReadCloser
		var err error

		key := object.GetFullKey()
		exists := true

		input := &s3.GetObjectInput{
			Bucket: object.bucket,
			Key:    aws.String(key),
		}

		out, err := r.client.GetObject(input)

		if err != nil {
			if awsErr, ok := err.(awserr.RequestFailure); ok && awsErr.StatusCode() == 404 {
				exists = false
				err = nil
			}
		} else {
			body = out.Body
		}

		r.writeMetric(operationRead)

		object.Body = StreamReader(body)
		object.Exists = exists
		object.Error = err
		object.wg.Done()
	}
}

func (r *BatchRunner) executeWrite() {
	for object := range r.write {
		key := object.GetFullKey()
		body := CloseOnce(object.Body.AsReader())

		input := &s3.PutObjectInput{
			ACL:    object.ACL,
			Body:   body,
			Bucket: object.bucket,
			Key:    aws.String(key),
		}

		if object.ContentEncoding != "" {
			input.ContentEncoding = aws.String(object.ContentEncoding)
		}

		if object.ContentType != "" {
			input.ContentType = aws.String(object.ContentType)
		}

		_, err := r.client.PutObject(input)

		if err != nil {
			object.Exists = false
			object.Error = err
		} else {
			object.Exists = true
		}

		if err := body.Close(); err != nil {
			object.Error = multierror.Append(object.Error, err)
		}

		r.writeMetric(operationWrite)

		object.wg.Done()
	}
}

func (r *BatchRunner) executeCopy() {
	for object := range r.copy {
		key := object.GetFullKey()
		source := object.getSource()

		input := &s3.CopyObjectInput{
			ACL:        object.ACL,
			Bucket:     object.bucket,
			Key:        aws.String(key),
			CopySource: aws.String(source),
		}

		_, err := r.client.CopyObject(input)

		if err != nil {
			object.Error = err
		}

		r.writeMetric(operationCopy)

		object.wg.Done()
	}
}

func (r *BatchRunner) executeDelete() {
	for object := range r.delete {
		key := object.GetFullKey()

		input := &s3.DeleteObjectInput{
			Bucket: object.bucket,
			Key:    aws.String(key),
		}

		_, err := r.client.DeleteObject(input)

		if err != nil {
			object.Error = err
		}

		r.writeMetric(operationDelete)

		object.wg.Done()
	}
}

func (r *BatchRunner) writeMetric(operation string) {
	r.metric.WriteOne(&mon.MetricDatum{
		MetricName: metricName,
		Priority:   mon.PriorityHigh,
		Dimensions: map[string]string{
			"Operation": operation,
		},
		Unit:  mon.UnitCount,
		Value: 1.0,
	})
}

func getDefaultRunnerMetrics() []*mon.MetricDatum {
	return []*mon.MetricDatum{
		{
			MetricName: metricName,
			Priority:   mon.PriorityHigh,
			Dimensions: map[string]string{
				"Operation": operationRead,
			},
			Unit:  mon.UnitCount,
			Value: 0.0,
		},
		{
			MetricName: metricName,
			Priority:   mon.PriorityHigh,
			Dimensions: map[string]string{
				"Operation": operationWrite,
			},
			Unit:  mon.UnitCount,
			Value: 0.0,
		},
		{
			MetricName: metricName,
			Priority:   mon.PriorityHigh,
			Dimensions: map[string]string{
				"Operation": operationCopy,
			},
			Unit:  mon.UnitCount,
			Value: 0.0,
		},
		{
			MetricName: metricName,
			Priority:   mon.PriorityHigh,
			Dimensions: map[string]string{
				"Operation": operationDelete,
			},
			Unit:  mon.UnitCount,
			Value: 0.0,
		},
	}
}
