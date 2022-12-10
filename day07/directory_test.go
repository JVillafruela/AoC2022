package day07

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDirectory(t *testing.T) {

	root := NewDirectory(nil, "/")
	want := Directory{Name: "/"}
	if !reflect.DeepEqual(root, want) {
		t.Errorf("NewDirectory(/) : got %v want %v", root, want)
	}

}

func TestDirectory_addChild(t *testing.T) {

	root := Directory{Name: "/"}
	dirA := Directory{Name: "A", Size: 1000}
	root.addChild(&dirA)

	want := Directory{Name: "/", Size: 0, Children: []*Directory{{Name: "A", Size: 1000, Parent: &root}}}
	if !reflect.DeepEqual(root, want) {
		t.Errorf("addChild(A) : got %v want %v", root, want)
	}

	dirE := Directory{Name: "E", Size: 2000}
	dirI := Directory{Name: "I", Size: 3000}

	dirA.addChild(&dirE)
	want = Directory{Name: "A", Size: 1000, Parent: &root, Children: []*Directory{{Name: "E", Size: 2000, Parent: &dirA}}}
	if !reflect.DeepEqual(dirA, want) {
		t.Errorf("addChildE) : got %v want %v", dirA, want)
	}

	dirE.addChild(&dirI)
	want = Directory{Name: "A", Size: 1000, Parent: &root,
		Children: []*Directory{{Name: "E", Size: 2000, Parent: &dirA,
			Children: []*Directory{{Name: "I", Size: 3000, Parent: &dirE}}}}}
	if !reflect.DeepEqual(dirA, want) {
		t.Errorf("addChildI) : got %v want %v", dirA, want)
	}

}

func TestDirectory_addSize(t *testing.T) {

	//root := Directory{Name: "/"}
	//root.addSize(1000)

	dirA := Directory{Name: "a", Size: 584}
	dirA.addSize(29116) // Size: 29116 + 2557 + 62596
	dirA.addSize(2557)
	dirA.addSize(62596)
	want := Directory{Name: "a", Size: 94853}
	if !reflect.DeepEqual(dirA, want) {
		t.Errorf("addSize(A) : got %v want %v", dirA, want)
	}

}

/*
- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
*/
func TestDirectory_size(t *testing.T) {

	root := Directory{Name: "/"}
	dirA := Directory{Name: "a", Size: 29116 + 2557 + 62596}
	dirE := Directory{Name: "e", Size: 0}
	dirI := Directory{Name: "i", Size: 584}

	root.addChild(&dirA)
	dirA.addChild(&dirE)
	dirE.addChild(&dirI)

	fmt.Printf("DirA %v \n", dirA)

	tests := []struct {
		name string
		dir  *Directory
		want int
	}{
		{"Dir i", &dirE, 584},
		{"Dir e", &dirE, 584},
		{"Dir a", &dirA, 94853},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dir.size(); got != tt.want {
				t.Errorf("Directory.size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectory_changeDir(t *testing.T) {
	root := Directory{Name: "/"}
	dirA := Directory{Name: "A", Size: 1000}
	root.addChild(&dirA)
	dirE := Directory{Name: "E", Size: 2000}
	dirA.addChild(&dirE)
	dirI := Directory{Name: "I", Size: 3000}
	dirE.addChild(&dirI)

	type args struct {
		name string
	}
	tests := []struct {
		name string
		dir  *Directory
		args args
		want *Directory
	}{ // up
		{"/ cd ..", &root, args{".."}, nil},
		{"A cd ..", &dirA, args{".."}, &root},
		{"E cd ..", &dirE, args{".."}, &dirA},
		{"I cd ..", &dirI, args{".."}, &dirE},
		// down
		{"/ cd A", &root, args{"A"}, &dirA},
		{"/ cd E", &root, args{"E"}, nil},
		{"A cd E", &dirA, args{"E"}, &dirE},
		{"E cd I", &dirE, args{"I"}, &dirI},
		{"E cd X", &dirE, args{"X"}, nil},
		// root
		{"/ cd /", &root, args{"/"}, &root},
		{"A cd /", &dirA, args{"/"}, &root},
		{"E cd /", &dirE, args{"/"}, &root},
		{"I cd /", &dirI, args{"/"}, &root},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dir.changeDir(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Directory.changeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
