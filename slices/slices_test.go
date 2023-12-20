package slices

import (
	"reflect"
	"strconv"
	"testing"
)

var filterIntTests = []struct {
	name string
	s    []int
	f    func(int) bool
	want []int
}{
	{
		name: "even",
		s:    []int{1, 2, 3, 4, 5},
		f:    func(i int) bool { return i%2 == 0 },
		want: []int{2, 4},
	},
	{
		name: "odd",
		s:    []int{1, 2, 3, 4, 5},
		f:    func(i int) bool { return i%2 != 0 },
		want: []int{1, 3, 5},
	},
}

var filterStringTests = []struct {
	name string
	s    []string
	f    func(string) bool
	want []string
}{
	{
		name: "even",
		s:    []string{"a", "2", "c", "4", "d"},
		f: func(s string) bool {
			_, err := strconv.Atoi(s)
			return err == nil
		},
		want: []string{"2", "4"},
	},
}

func TestFilter(t *testing.T) {
	for _, tt := range filterIntTests {
		got := Filter(tt.s, tt.f)
		if !reflect.DeepEqual(got, tt.want) {
			//if len(got) != len(tt.want) {
			t.Errorf("%s Filter : %v = %v, want %v", tt.name, tt.s, got, tt.want)
		}
	}

	for _, tt := range filterStringTests {
		got := Filter(tt.s, tt.f)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s Filter : %v = %v, want %v", tt.name, tt.s, got, tt.want)
		}
	}
}
