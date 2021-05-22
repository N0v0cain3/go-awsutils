package awsutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	//REGION ..
	REGION = "ap-south-1"
)

//GetSession ..
func GetSession() *session.Session {
	newSession := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(REGION),
		}))
	return newSession
}
