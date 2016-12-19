package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"fmt"
	"os"
	"bytes"
	"flag"
)

var (
	bucketName, fileName string
)

func main() {

	session, err := session.NewSession()
	if err != nil {
		fmt.Println("Failed to create AWS session,", err)
	}

	flag.StringVar(&bucketName,"bucket","","Enter the name of the s3 bucket")
	flag.StringVar(&fileName,"filename","","Enter the name of the file")

	flag.Parse()


	var fileToUpload  string = fileName

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
	UploadInst := s3manager.NewUploader(session)

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
		fmt.Printf("The upload resulted in: %v", result)
	}



}

