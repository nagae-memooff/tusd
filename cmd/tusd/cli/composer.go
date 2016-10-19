package cli

import (
	"os"

	"github.com/nagae-memooff/tusd"
	"github.com/nagae-memooff/tusd/filestore"
	"github.com/nagae-memooff/tusd/limitedstore"
	"github.com/nagae-memooff/tusd/memorylocker"
	"github.com/nagae-memooff/tusd/s3store"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var Composer *tusd.StoreComposer

func CreateComposer() {
	// Attempt to use S3 as a backend if the -s3-bucket option has been supplied.
	// If not, we default to storing them locally on disk.
	Composer = tusd.NewStoreComposer()
	if Flags.S3Bucket == "" {
		dir := Flags.UploadDir

		stdout.Printf("Using '%s' as directory storage.\n", dir)
		if err := os.MkdirAll(dir, os.FileMode(0774)); err != nil {
			stderr.Fatalf("Unable to ensure directory exists: %s", err)
		}

		store := filestore.New(dir)
		store.UseIn(Composer)
	} else {
		stdout.Printf("Using 's3://%s' as S3 bucket for storage.\n", Flags.S3Bucket)

		// Derive credentials from AWS_SECRET_ACCESS_KEY, AWS_ACCESS_KEY_ID and
		// AWS_REGION environment variables.
		credentials := aws.NewConfig().WithCredentials(credentials.NewEnvCredentials())
		store := s3store.New(Flags.S3Bucket, s3.New(session.New(), credentials))
		store.UseIn(Composer)

		locker := memorylocker.New()
		locker.UseIn(Composer)
	}

	storeSize := Flags.StoreSize
	maxSize := Flags.MaxSize

	if storeSize > 0 {
		limitedstore.New(storeSize, Composer.Core, Composer.Terminater).UseIn(Composer)
		stdout.Printf("Using %.2fMB as storage size.\n", float64(storeSize)/1024/1024)

		// We need to ensure that a single upload can fit into the storage size
		if maxSize > storeSize || maxSize == 0 {
			Flags.MaxSize = storeSize
		}
	}

	stdout.Printf("Using %.2fMB as maximum size.\n", float64(Flags.MaxSize)/1024/1024)
}
