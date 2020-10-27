package registry

import (
	"context"
	"sync"
)

var once sync.Once
var metaService MetadataService

type MetadataService interface {
	Subscribe(ctx context.Context, metadata *ServiceMetadata) (chan struct{}, error)
}

func GetMetadataService() MetadataService {
	once.Do(func() {
		//todo
	})
	return nil
}

type discovery struct {

}
