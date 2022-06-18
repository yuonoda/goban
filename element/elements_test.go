package element

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestElements_Len(t *testing.T) {
	type fields struct {
		elements Elements[NumericElement]
	}
	tests := []struct {
		name string
		fields
		want int
	}{
		{
			name: "pass",
			fields: fields{
				elements: Elements[NumericElement]{
					NumericElement(1),
					NumericElement(2),
					NumericElement(3),
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.elements.Len()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}

}
