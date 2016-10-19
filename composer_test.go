package tusd_test

import (
	"github.com/nagae-memooff/tusd"
	"github.com/nagae-memooff/tusd/filestore"
	"github.com/nagae-memooff/tusd/limitedstore"
	"github.com/nagae-memooff/tusd/memorylocker"
)

func ExampleNewStoreComposer() {
	composer := tusd.NewStoreComposer()

	fs := filestore.New("./data")
	fs.UseIn(composer)

	ml := memorylocker.New()
	ml.UseIn(composer)

	ls := limitedstore.New(1024*1024*1024, composer.Core, composer.Terminater)
	ls.UseIn(composer)

	config := tusd.Config{
		StoreComposer: composer,
	}

	_, _ = tusd.NewHandler(config)
}
