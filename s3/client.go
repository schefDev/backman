package s3

import (
	cfenv "github.com/cloudfoundry-community/go-cfenv"
	minio "github.com/minio/minio-go/v6"
	"gitlab.swisscloud.io/appc-cf-core/appcloud-backman-app/env"
	"gitlab.swisscloud.io/appc-cf-core/appcloud-backman-app/log"
)

// Client is used interact with S3 storage
type Client struct {
	Client     *minio.Client
	BucketName string
}

func New(app *cfenv.App) *Client {
	// read env
	s3ServiceLabel := env.Get("S3_SERVICE_LABEL", "dynstrg")

	// setup minio/s3 client
	s3Services, err := app.Services.WithLabel(s3ServiceLabel)
	if err != nil {
		log.Fatalf("could not get s3 service from VCAP environment: %v", err)
	}
	if len(s3Services) != 1 {
		log.Fatalf("there must be exactly one defined S3 service, but found %d instead", len(s3Services))

	}
	bucketName := env.Get("S3_BUCKET_NAME", s3Services[0].Name)
	if len(bucketName) == 0 {
		log.Fatalln("bucket name for S3 storage is not configured properly")
	}
	endpoint, _ := s3Services[0].CredentialString("accessHost")
	accessKeyID, _ := s3Services[0].CredentialString("accessKey")
	secretAccessKey, _ := s3Services[0].CredentialString("sharedSecret")
	useSSL := true

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// check if bucket exists and is accessible and if not create it, or fail
	exists, errBucketExists := minioClient.BucketExists(bucketName)
	if errBucketExists == nil && exists {
		log.Infof("S3 bucket [%s] found", bucketName)
	} else {
		if err := minioClient.MakeBucket(bucketName, ""); err != nil {
			log.Fatalf("S3 bucket [%s] could not be created: %v", bucketName, err)
			exists, errBucketExists := minioClient.BucketExists(bucketName)
			if errBucketExists != nil || exists {
				log.Fatalf("S3 bucket [%s] is not accessible: %v", bucketName, err)
			}
		} else {
			log.Infof("new S3 bucket [%s] was successfully created", bucketName)
		}
	}

	return &Client{
		Client:     minioClient,
		BucketName: bucketName,
	}
}

func (s *Client) ListObjects(folderPath string) ([]minio.ObjectInfo, error) {
	// read objects from S3
	doneCh := make(chan struct{})
	defer close(doneCh)

	isRecursive := true
	objects := make([]minio.ObjectInfo, 0)
	objectCh := s.Client.ListObjectsV2(s.BucketName, folderPath, isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			log.Errorf("could not read S3 object: %v", object.Err)
			return nil, object.Err
		}
		objects = append(objects, object)
	}
	return objects, nil
}

func (s *Client) DeleteObject(object string) error {
	// delete object from S3
	if err := s.Client.RemoveObject(s.BucketName, object); err != nil {
		log.Errorf("could not delete S3 object [%s]: %v", object, err)
		return err
	}
	return nil
}