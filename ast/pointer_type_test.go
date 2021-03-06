package ast

import (
	"testing"
)

func TestPointerType(t *testing.T) {
	nodes := map[string]Node{
		`0x7fa3b88bbb30 'struct _opaque_pthread_t *'`: &PointerType{
			Address:  "0x7fa3b88bbb30",
			Type:     "struct _opaque_pthread_t *",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
