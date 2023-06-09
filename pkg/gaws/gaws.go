package gaws

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func removeFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println(err)
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func checkerror(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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
	s3Client := s3.New(sess)
	return s3Client, nil
}

func ListBuckets() (*s3.ListBucketsOutput, error) {

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

func DeleteObjects(bucket string) {
	client, err := sessionInit()
	if err != nil {
		return
	}
	// Setup BatchDeleteIterator to iterate through a list of objects.
	iter := s3manager.NewDeleteListIterator(client, &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})

	// Traverse iterator deleting each object
	if err := s3manager.NewBatchDeleteWithClient(client).Delete(aws.BackgroundContext(), iter); err != nil {
		exitErrorf("Unable to delete objects from bucket %q, %v", bucket, err)
	}

	fmt.Printf("Deleted object(s) from bucket: %s", bucket)

}

func DeleteBucket(bucket string) {
	client, err := sessionInit()
	if err != nil {
		return
	}
	_, err = client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Unable to delete bucket %q, %v", bucket, err)
	}

	// Wait until bucket is deleted before finishing
	fmt.Printf("Waiting for bucket %q to be deleted...\n", bucket)

	err = client.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be deleted, %v", bucket)
	}

	fmt.Printf("Bucket %q successfully deleted\n", bucket)

}

func CreateBucket(bucket string) {
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

func ArchiveDir(output string, input string) *os.File {

	//attempts to delete file if it already exists
	removeFile(output + ".zip")

	file, err := os.Create(output + ".zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}
		return nil
	}
	err = filepath.Walk(input, walker)
	if err != nil {
		panic(err)
	}
	return file
}

func S3Fileupload(filename string, file *os.File, bucket string) {
	filename = filename + ".zip"
	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", file, err)
	}

	defer file.Close()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)

	removeFile(filename)
}
