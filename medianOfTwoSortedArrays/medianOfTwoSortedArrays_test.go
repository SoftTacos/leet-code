package medianOfTwoSortedArrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name       string
		args       args
		wantMedian float64
	}{
		{
			name: "empty",
			args: args{
				nums1: []int{},
				nums2: []int{},
			},
			wantMedian: 0,
		},
		{
			name: "one",
			args: args{
				nums1: []int{1},
				nums2: []int{},
			},
			wantMedian: 1,
		},
		{
			name: "bunch in one",
			args: args{
				nums1: []int{1, 2, 3, 4},
				nums2: []int{},
			},
			wantMedian: 2.5,
		},
		{
			name: "bunch",
			args: args{
				nums1: []int{1, 2, 3, 4},
				nums2: []int{5, 6, 7, 8},
			},
			wantMedian: 4.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMedian := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); gotMedian != tt.wantMedian {
				t.Errorf("findMedianSortedArrays() = %v, want %v", gotMedian, tt.wantMedian)
			}
		})
	}
}
