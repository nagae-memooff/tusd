package cli

import (
	"os"

	"github.com/nagae-memooff/tusd"
	"github.com/nagae-memooff/tusd/filestore"
	"github.com/nagae-memooff/tusd/limitedstore"

)

var Composer *tusd.StoreComposer

func CreateComposer() {
	// Attempt to use S3 as a backend if the -s3-bucket option has been supplied.
	// If not, we default to storing them locally on disk.
	Composer = tusd.NewStoreComposer()
	dir := Flags.UploadDir

	stdout.Printf("Using '%s' as directory storage.\n", dir)
	if err := os.MkdirAll(dir, os.FileMode(0774)); err != nil {
		stderr.Fatalf("Unable to ensure directory exists: %s", err)
	}

	store := filestore.New(dir)
	store.UseIn(Composer)

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
