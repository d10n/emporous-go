package defaultmanager

import (
	"context"

	empspec "github.com/emporous/collection-spec/specs-go/v1alpha1"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/emporous/emporous-go/attributes"
	"github.com/emporous/emporous-go/nodes/descriptor/v2"
	"github.com/emporous/emporous-go/registryclient"
)

func (d DefaultManager) ReadLayer(ctx context.Context, source string, title string, remote registryclient.Remote) ([]byte, error) {
	graph, err := remote.LoadCollection(ctx, source)
	if err != nil {
		return nil, err
	}
	var target ocispec.Descriptor
	titleAttribute := attributes.NewString(ocispec.AnnotationTitle, title)
	for _, node := range graph.Nodes() {
		// Check that this is a descriptor node and the blob is
		// not a config or schema resource.
		desc, ok := node.(*v2.Node)
		if !ok {
			continue
		}
		switch desc.Descriptor().MediaType {
		case empspec.MediaTypeSchemaDescriptor:
			continue
		case ocispec.MediaTypeImageConfig:
			continue
		case empspec.MediaTypeConfiguration:
			continue
		}
		exists, err := desc.Attributes().Exists(titleAttribute)
		if err != nil {
			return nil, err
		}
		if exists {
			target = desc.Descriptor()
			break
		}
	}
	bytes, err := remote.GetContent(ctx, source, target)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

//func (d DefaultManager) ReadLayerStream(ctx context.Context, source string, title string, remote registryclient.Remote) (io.ReadCloser, error) {
//	// TODO
//}