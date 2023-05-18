package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args[V Element] struct {
		data      OrderedSlice[V]
		direction Direction
	}
	type testCase[V Element] struct {
		name string
		args args[V]
		want OrderedSlice[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			args: args[int]{data: []int{}, direction: Asc},
			want: OrderedSlice[int]([]int{}),
		},
		{
			name: "one element",
			args: args[int]{data: []int{1}, direction: Asc},
			want: OrderedSlice[int]([]int{1}),
		},
		{
			name: "ascending",
			args: args[int]{data: []int{74, 59, 238}, direction: Asc},
			want: OrderedSlice[int]([]int{59, 74, 238}),
		},
		{
			name: "descending",
			args: args[int]{data: []int{74, 59, 238}, direction: Desc},
			want: OrderedSlice[int]([]int{238, 74, 59}),
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.data, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}

	stringTests := []testCase[string]{
		{
			name: "empty",
			args: args[string]{data: []string{}, direction: Asc},
			want: OrderedSlice[string]([]string{}),
		},
		{
			name: "one element",
			args: args[string]{data: []string{"a"}, direction: Asc},
			want: OrderedSlice[string]([]string{"a"}),
		},
		{
			name: "ascending",
			args: args[string]{data: []string{"c", "a", "b"}, direction: Asc},
			want: OrderedSlice[string]([]string{"a", "b", "c"}),
		},
		{
			name: "descending",
			args: args[string]{data: []string{"c", "a", "b"}, direction: Desc},
			want: OrderedSlice[string]([]string{"c", "b", "a"}),
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.data, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkMerge[V Element](data OrderedSlice[V], direction Direction, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Merge(data, direction)
	}
}

func BenchmarkMergeInt(b *testing.B) {
	benchmarkMerge([]int{4, 5, 8, 7}, Asc, b)
}
func BenchmarkMergeString(b *testing.B) {
	benchmarkMerge([]string{"c", "a", "b"}, Asc, b)
}
