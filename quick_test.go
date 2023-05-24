package sort

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestQuick(t *testing.T) {
	type args[V constraints.Ordered] struct {
		data      Slice[V]
		direction direction
	}

	type testCase[V constraints.Ordered] struct {
		name string
		args args[V]
		want Slice[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			args: args[int]{data: []int{}, direction: Asc},
			want: Slice[int]([]int{}),
		},
		{
			name: "one element",
			args: args[int]{data: []int{1}, direction: Asc},
			want: Slice[int]([]int{1}),
		},
		{
			name: "ascending",
			args: args[int]{data: []int{74, 59, 238}, direction: Asc},
			want: Slice[int]([]int{59, 74, 238}),
		},
		{
			name: "descending",
			args: args[int]{data: []int{74, 59, 238}, direction: Desc},
			want: Slice[int]([]int{238, 74, 59}),
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			Quick(tt.args.data, tt.args.direction)
			if !reflect.DeepEqual(tt.args.data, tt.want) {
				t.Errorf("Quick() = %v, want %v", tt.args.data, tt.want)
			}
		})
	}

	stringTests := []testCase[string]{
		{
			name: "empty",
			args: args[string]{data: []string{}, direction: Asc},
			want: Slice[string]([]string{}),
		},
		{
			name: "one element",
			args: args[string]{data: []string{"a"}, direction: Asc},
			want: Slice[string]([]string{"a"}),
		},
		{
			name: "ascending",
			args: args[string]{data: []string{"c", "a", "b"}, direction: Asc},
			want: Slice[string]([]string{"a", "b", "c"}),
		},
		{
			name: "descending",
			args: args[string]{data: []string{"c", "a", "b"}, direction: Desc},
			want: Slice[string]([]string{"c", "b", "a"}),
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			Quick(tt.args.data, tt.args.direction)
			if !reflect.DeepEqual(tt.args.data, tt.want) {
				t.Errorf("Quick() = %v, want %v", tt.args.data, tt.want)
			}
		})
	}
}

func TestQuickStruct(t *testing.T) {
	type args[T constraints.Ordered, V Element[T]] struct {
		data      []Person
		direction direction
	}
	type testCase[T constraints.Ordered, V Element[T]] struct {
		name string
		args args[T, V]
		want []Person
	}
	tests := []testCase[int, Element[int]]{
		{
			name: "empty",
			args: args[int, Element[int]]{data: []Person{}, direction: Asc},
			want: []Person{},
		},
		{
			name: "one element",
			args: args[int, Element[int]]{data: []Person{{"Jenny", 26}}, direction: Asc},
			want: []Person{{"Jenny", 26}},
		},
		{
			name: "structure Asc",
			args: args[int, Element[int]]{
				data: []Person{
					{"Jenny", 26},
					{"John", 42},
					{"Michael", 17},
					{"Bob", 31},
				},
				direction: Asc,
			},
			want: []Person{
				{"Michael", 17},
				{"Jenny", 26},
				{"Bob", 31},
				{"John", 42},
			},
		},
		{
			name: "structure Desc",
			args: args[int, Element[int]]{
				data: []Person{
					{"Jenny", 26},
					{"John", 42},
					{"Michael", 17},
					{"Bob", 31},
				},
				direction: Desc,
			},
			want: []Person{
				{"John", 42},
				{"Bob", 31},
				{"Jenny", 26},
				{"Michael", 17},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickStruct[int](tt.args.data, tt.args.direction)
			if !reflect.DeepEqual(tt.args.data, tt.want) {
				t.Errorf("QuickStruct() = %v, want %v", tt.args.data, tt.want)
			}
		})
	}
}

// Benchmarks
func benchmarkQuick[V constraints.Ordered](data Slice[V], direction direction, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Quick(data, direction)
	}
}

func BenchmarkQuickInt(b *testing.B) {
	benchmarkQuick([]int{4, 5, 8, 7}, Asc, b)
}
func BenchmarkQuickString(b *testing.B) {
	benchmarkQuick([]string{"c", "a", "b"}, Asc, b)
}

func BenchmarkQuickStruct(b *testing.B) {
	benchmarks := []struct {
		name      string
		data      []Person
		direction direction
	}{
		{"structure",
			[]Person{
				{"John", 42},
				{"Bob", 31},
				{"Jenny", 26},
				{"Michael", 17},
			},
			Asc,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				QuickStruct[int](bm.data, bm.direction)
			}
		})
	}
}
