package gaws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
	"archive/tar"

	"io"
	"os"
	"path/filepath"
	
)

func sessionInit() (*s3.S3, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})
	if err != nil {
		return nil, err
	}

	s3Client:= s3.New(sess)
	
	return s3Client, nil
}

func BucketList() (*s3.ListBucketsOutput, error) {

	client, err := sessionInit()
	if err != nil {
		return nil, err
	}
	
	buckets, err := listBuckets(client)
	if err != nil {
		return nil, err
	}

	println("\nThe following buckets were found:\n")

	for _, bucket := range buckets.Buckets {
		fmt.Printf("%s\n", *bucket.Name)
	}
	return buckets, nil


}

func listBuckets(client *s3.S3) (*s3.ListBucketsOutput, error) {
	res, err := client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateBucket(bucket string) () {
	client, err := sessionInit()
	if err != nil {
		return 
	}
        	
	_, err = client.CreateBucket(&s3.CreateBucketInput{
        Bucket: aws.String(bucket),

	})
	if err != nil {
		fmt.Printf("Unable to create bucket %q, %v", bucket, err)
		return
	}

	    // Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be created...\n", bucket)
	
	err = client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		fmt.Printf("Error occurred while waiting for bucket to be created, %v", bucket)
		return
	}
	
	fmt.Printf("Bucket %q successfully created\n", bucket)
}

func checkerror(err error) {
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
}
func Tarfunc() () {
 
   destinationfile := "tarball.tar"
   sourcedir := "/home/jim/git/go/src/gaws"
 
   dir, err := os.Open(sourcedir)
   checkerror(err)
   defer dir.Close()
 
   // get list of files
   files, err := dir.Readdir(0)
   checkerror(err)
 
   // create tar file
   tarfile, err := os.Create(destinationfile)
   checkerror(err)
   defer tarfile.Close()
 
   var fileWriter io.WriteCloser = tarfile
 
   tarfileWriter := tar.NewWriter(fileWriter)
   defer tarfileWriter.Close()
 
   for _, fileInfo := range files {
 
	        if fileInfo.IsDir() {
	   continue
	   }
 
      file, err := os.Open(dir.Name() + string(filepath.Separator) + fileInfo.Name())
      checkerror(err)
      defer file.Close()
 
      // prepare the tar header
      header := new(tar.Header)
      header.Name = file.Name()
      header.Size = fileInfo.Size()
      header.Mode = int64(fileInfo.Mode())
      header.ModTime = fileInfo.ModTime()
 
      err = tarfileWriter.WriteHeader(header)
      checkerror(err)
 
      _, err = io.Copy(tarfileWriter, file)
      checkerror(err)
   }
}
