package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"fmt"
	"os"
	"bytes"
	"launchpad.net/gnuflag"
)

var (
	bucketName, fileName, localFile string
)

func main() {

	sesh, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	gnuflag.StringVar(&bucketName,"bucket","","Enter the name of the s3 bucket")
	gnuflag.StringVar(&fileName,"filename","","Enter the name or path of the file")
	gnuflag.StringVar(&localFile,"localfile","","Optional. Use to define the file you want to upload if the name is different from what's declared with the 'filename' option ")

	gnuflag.Parse(true)


	var fileToUpload string

	if len(localFile) == 0 {
		fileToUpload = fileName
	}

	file, err := os.Open(fileToUpload)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)

	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer)

	// Uploader Instance
	UploadInst := s3manager.NewUploader(sesh)

	// Upload input params
	uploadParams := &s3manager.UploadInput{
		Bucket: &bucketName,      // Name of s3 bucket
		Key:    &fileName,        // Name of uploaded file when in bucket
		Body:   fileBytes,        // Local file to upload
	}

	result, err := UploadInst.Upload(uploadParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The upload resulted in:\n%v", result)
	}
}
