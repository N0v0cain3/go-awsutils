package awsutils

import (
	"errors"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

// UploadFile uploads a file to s3 bucket.
func UploadFile(file io.Reader, fileName string) (*s3manager.UploadOutput, error) {
	// Create an uploader with the session and default options
	sess := GetSession()
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(viper.GetString("BUCKET")),
		Key:    aws.String(fileName),
		ACL:    aws.String("public-read"),
		Body:   file,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

//DownloadFile
func DownloadFile(link string) {
	sess := GetSession()
	buff := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(sess)
	downloader.Download(buff, &s3.GetObjectInput{})
}

// DeleteFile deletes a file from s3 bucket.
func DeleteFile(key string) error {
	sess := GetSession()
	svc := s3.New(sess)
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(viper.GetString("BUCKET")), Key: aws.String(key)})
	if err != nil {
		return errors.New("s3 error: unable to delete file ")
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(viper.GetString("BUCKET")),
		Key:    aws.String(key),
	})
	return err
}
