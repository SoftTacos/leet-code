package findmedianfromdatastream

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder_FindMedian(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		want       []float64
	}{
		{
			name:       "single",
			operations: []string{"1", "get"},
			want:       []float64{1},
		},
		{
			name:       "two",
			operations: []string{"1", "2", "get"},
			want:       []float64{1.5},
		},
		{
			name:       "three",
			operations: []string{"1", "2", "3", "get"},
			want:       []float64{2},
		},
		// {
		// 	name:       "duplicates odd",
		// 	operations: []string{"1", "2", "2", "get"},
		// 	want:       []float64{2},
		// },
		{
			name:       "duplicates odd",
			operations: []string{"1", "1", "1", "10", "-1", "2", "123", "get"},
			want:       []float64{1},
		},
		{
			name:       "duplicates even",
			operations: []string{"2", "1", "10", "2", "get"},
			want:       []float64{2},
		},
		{
			name:       "four",
			operations: []string{"1", "2", "3", "4", "get"},
			want:       []float64{2.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				finder = Constructor()
				got    []float64
				num    int64
			)

			for _, op := range tt.operations {
				if op == "get" {
					got = append(got, finder.FindMedian())
				} else {
					num, _ = strconv.ParseInt(op, 10, 64)
					finder.AddNum(int(num))
				}
			}
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
