package defaultmanager

import (
	"context"

	"github.com/emporous/emporous-go/nodes/collection"
	"github.com/emporous/emporous-go/registryclient"
)

func (d DefaultManager) List(ctx context.Context, reference string, remote registryclient.Remote) (*collection.Collection, error) {
	loadCollection, err := remote.LoadCollection(ctx, reference)
	if err != nil {
		return nil, err
	}
	return &loadCollection, nil
}
