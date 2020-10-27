package s3client

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Mock struct {
	manager.UploadAPIClient
	putObjectErr error
	output       *s3.PutObjectOutput
}

func (s s3Mock) PutObject(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return s.output, s.putObjectErr
}

func Test_s3PutObject_Upload(t *testing.T) {
	type fields struct {
		s3Client manager.UploadAPIClient
	}
	type args struct {
		buffer   []byte
		fileName string
		bucket   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				s3Client: &s3Mock{},
			},
			args: args{
				buffer:   []byte{},
				fileName: "test.txt",
				bucket:   "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &s3PutObject{
				s3Client: tt.fields.s3Client,
			}
			if err := s.Upload(tt.args.buffer, tt.args.fileName, tt.args.bucket); (err != nil) != tt.wantErr {
				t.Errorf("s3PutObject.Upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
