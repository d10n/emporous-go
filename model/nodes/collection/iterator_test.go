package collection

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uor-framework/client/model"
	"github.com/uor-framework/client/util/testutils"
)

var iteratorTests = []struct {
	nodes []model.Node
	want  []model.Node
}{
	{nodes: nil, want: nil},
	{
		nodes: []model.Node{
			&testutils.MockNode{
				I: "node1",
				A: testutils.MockAttributes{
					"kind": "txt",
					"name": "test",
				},
			},
		},
		want: []model.Node{
			&testutils.MockNode{
				I: "node1",
				A: testutils.MockAttributes{
					"kind": "txt",
					"name": "test",
				},
			},
		},
	},
	{
		nodes: []model.Node{
			&testutils.MockNode{
				I: "node1",
				A: testutils.MockAttributes{
					"kind": "txt",
					"name": "test",
				},
			},
			&testutils.MockNode{
				I: "node2",
				A: testutils.MockAttributes{
					"kind": "txt",
				},
			},
		},
		want: []model.Node{
			&testutils.MockNode{
				I: "node2",
				A: testutils.MockAttributes{
					"kind": "txt",
				},
			},
			&testutils.MockNode{
				I: "node1",
				A: testutils.MockAttributes{
					"kind": "txt",
					"name": "test",
				},
			},
		},
	},
}

func TestByAttributeIterator(t *testing.T) {
	for _, test := range iteratorTests {
		it := NewByAttributesIterator(test.nodes)
		for i := 0; i < 2; i++ {
			require.Equal(t, it.Len(), len(test.nodes))
			var got []model.Node
			for it.Next() {
				got = append(got, it.Node())
				require.Equal(t, len(got)+it.Len(), len(test.nodes))
			}
			require.Equal(t, test.want, got)
			it.Reset()
		}
	}
}
