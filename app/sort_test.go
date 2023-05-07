package sort

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	type args[V Element] struct {
		data []V
		asc  bool
	}
	type testCase[V Element] struct {
		name string
		args args[V]
		want []V
	}
	intTests := []testCase[int]{
		{name: "int ascending", args: args[int]{data: []int{74, 59, 238}, asc: true}, want: []int{59, 74, 238}},
		{name: "int descending", args: args[int]{data: []int{74, 59, 238}, asc: false}, want: []int{238, 74, 59}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bubble(tt.args.data, tt.args.asc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bubble() = %v, want %v", got, tt.want)
			}
		})
	}

	stringTests := []testCase[string]{
		{name: "string ascending", args: args[string]{data: []string{"c", "a", "b"}, asc: true}, want: []string{"a", "b", "c"}},
		{name: "string descending", args: args[string]{data: []string{"c", "a", "b"}, asc: false}, want: []string{"c", "b", "a"}},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bubble(tt.args.data, tt.args.asc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertion(t *testing.T) {
	type args[V Element] struct {
		data []V
		asc  bool
	}
	type testCase[V Element] struct {
		name string
		args args[V]
		want []V
	}
	intTests := []testCase[int]{
		{name: "int ascending", args: args[int]{data: []int{74, 59, 238}, asc: true}, want: []int{59, 74, 238}},
		{name: "int descending", args: args[int]{data: []int{74, 59, 238}, asc: false}, want: []int{238, 74, 59}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insertion(tt.args.data, tt.args.asc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insertion() = %v, want %v", got, tt.want)
			}
		})
	}

	stringTests := []testCase[string]{
		{name: "string ascending", args: args[string]{data: []string{"c", "a", "b"}, asc: true}, want: []string{"a", "b", "c"}},
		{name: "string descending", args: args[string]{data: []string{"c", "a", "b"}, asc: false}, want: []string{"c", "b", "a"}},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insertion(tt.args.data, tt.args.asc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insertion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args[V Element] struct {
		data []V
		asc  bool
	}
	type testCase[V Element] struct {
		name string
		args args[V]
		want []V
	}
	intTests := []testCase[int]{
		{name: "int ascending", args: args[int]{data: []int{74, 59, 238}, asc: true}, want: []int{59, 74, 238}},
		{name: "int descending", args: args[int]{data: []int{74, 59, 238}, asc: false}, want: []int{238, 74, 59}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.data, tt.args.asc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}

	stringTests := []testCase[string]{
		{name: "string ascending", args: args[string]{data: []string{"c", "a", "b"}, asc: true}, want: []string{"a", "b", "c"}},
		{name: "string descending", args: args[string]{data: []string{"c", "a", "b"}, asc: false}, want: []string{"c", "b", "a"}},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.data, tt.args.asc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
