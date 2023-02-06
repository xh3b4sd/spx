package apicliaws

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/xh3b4sd/tracer"
)

func (a *AWS) Upload(buc string, key string, rea bytes.Reader) error {
	{
		inp := &s3.PutObjectInput{
			Bucket: aws.String(buc),
			Key:    aws.String(key),
			Body: &Reader{
				rea: rea,
				siz: rea.Size(),
			},
			ACL: types.ObjectCannedACLPublicRead,
		}

		_, err := manager.NewUploader(a.S3, par).Upload(context.Background(), inp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		fmt.Printf("\n")
	}

	return nil
}

func par(u *manager.Uploader) {
	u.PartSize = 5 * 1024 * 1024
	u.LeavePartsOnError = true
}
