package mediaconvertclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/service/mediaconvert"
)

// Set the URL of a downstream client
var clientHost = "https://127.0.0.1:8080/"
var adminToken = "exampleToken"

// Set the names of all downstream services
var validServices = map[string]bool {
	"exampleServiceA": true,
    "exampleServiceB": true,
}
// CreateEncodeJob Create a new AWS MediaConvert encode job
func CreateEncodeJob(s3url string, body []byte) error {


	req, resp := client.CreateJobRequest(params)
	err := req.Send()
	if err == nil { // resp is now filled
		fmt.Println(resp)
	}
	func (c *MediaConvert) CreateJob(input *CreateJobInput) (*CreateJobOutput, error)

	if !validServices[service] {
		return error(fmt.Errorf("Service %s is not a valid Service", service))
	}
	url := fmt.Sprintf(clientHost + service)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("admintoken", adminToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return error(fmt.Errorf("Error sending request to %s due to: %v", url, err))
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return error(fmt.Errorf("Error reading response from %s due to %v", url, err))
	}
	return err
}
