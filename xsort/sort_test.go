package xsort

import (
	"fmt"
	"testing"
)

type SortableObj struct {
	Name string
}

func (c SortableObj) Less(other SortableObj) bool {
	return c.Name < other.Name
}

func TestSort(t *testing.T) {
	type args struct {
		slices []SortableObj
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{
				slices: []SortableObj{
					{
						Name: "Name1",
					},
					{
						Name: "Name2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceDesc(tt.args.slices)
			fmt.Printf("%#v", tt.args.slices)
		})
	}
}
