package sort

import (
	"reflect"
	"sort"
	"testing"
)

func TestBubble(t *testing.T) {
	type args struct {
		data      sort.Interface
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want sort.Interface
	}{
		{
			name: "empty",
			args: args{data: OrderedSlice[int]([]int{}), direction: Asc},
			want: OrderedSlice[int]([]int{}),
		},
		{
			name: "one element",
			args: args{data: OrderedSlice[int]([]int{1}), direction: Asc},
			want: OrderedSlice[int]([]int{1}),
		},
		{
			name: "int ascending",
			args: args{data: OrderedSlice[int]([]int{4, 5, 8, 7}), direction: Asc},
			want: OrderedSlice[int]([]int{4, 5, 7, 8}),
		},
		{
			name: "int descending",
			args: args{data: OrderedSlice[int]([]int{4, 5, 8, 7}), direction: Desc},
			want: OrderedSlice[int]([]int{8, 7, 5, 4}),
		},
		{
			name: "string ascending",
			args: args{data: OrderedSlice[string]([]string{"c", "a", "b"}), direction: Asc},
			want: OrderedSlice[string]([]string{"a", "b", "c"}),
		},
		{
			name: "string descending",
			args: args{data: OrderedSlice[string]([]string{"c", "a", "b"}), direction: Desc},
			want: OrderedSlice[string]([]string{"c", "b", "a"}),
		},
		{
			name: "structure ascending",
			args: args{
				data: ByAge([]Person{
					{"Bob", 31},
					{"John", 42},
					{"Michael", 17},
					{"Jenny", 26},
				}),
				direction: Asc,
			},
			want: ByAge([]Person{
				{"Michael", 17},
				{"Jenny", 26},
				{"Bob", 31},
				{"John", 42},
			}),
		},
		{
			name: "structure descending",
			args: args{
				data: ByAge([]Person{
					{"Bob", 31},
					{"John", 42},
					{"Michael", 17},
					{"Jenny", 26},
				}),
				direction: Desc,
			},
			want: ByAge([]Person{
				{"John", 42},
				{"Bob", 31},
				{"Jenny", 26},
				{"Michael", 17},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bubble(tt.args.data, tt.args.direction)
			if !reflect.DeepEqual(tt.args.data, tt.want) {
				t.Errorf("Bubble() = %v, want %v", tt.args.data, tt.want)
			}
		})
	}
}
