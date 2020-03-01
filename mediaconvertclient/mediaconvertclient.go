package mediaconvertclient

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
)
var sess = session.Must(
		session.NewSessionWithOptions(
			session.Options{SharedConfigState: session.SharedConfigEnable,
}))

var client = mediaconvert.New(sess)

// CreateEncodeJob Create a new AWS MediaConvert encode job
func CreateEncodeJob(s3url string, rendition string) error {
	var i *mediaconvert.Input
	i.FileInput := s3url
	settings := mediaconvert.JobSettings(i)

	var input *mediaconvert.CreateJobInput
	input.Settings.Inputs = settings

	res, err := client.CreateJob(input)
	if err != nil {
		msg := fmt.Sprintf("Error creating encode job due to %v", err)
		log.Print(msg)
		return err
	}

	return res
}
