package day07

type Directory struct {
	Name string
	// size of files inside the directory
	Size     int
	Parent   *Directory
	Children []*Directory
}

func NewDirectory(parent *Directory, name string) Directory {
	dir := Directory{Name: name}
	if parent != nil {
		dir.Parent = parent
	}
	return dir
}

func (dir *Directory) addChild(d *Directory) {
	d.Parent = dir
	dir.Children = append(dir.Children, d)
}

func (dir *Directory) addSize(size int) {
	dir.Size += size
}

func (dir *Directory) size() int {
	size := dir.Size
	//fmt.Printf(">Dir %s Size %d %v \n", dir.Name, size, dir)
	for _, d := range dir.Children {
		//fmt.Printf("   Calc subdir %s of %s \n", d.Name, dir.Name)
		n := d.size()
		//fmt.Printf("   Subdir %s of %s  size =%d \n", d.Name, dir.Name, n)
		size += n
	}
	//fmt.Printf("<Dir %s size %d \n", dir.Name, size)
	return size
}

func (dir *Directory) Root() *Directory {
	if dir.Name == "/" || dir.Parent == nil {
		return dir
	}
	return dir.Parent.Root()
}

func (dir *Directory) changeDir(name string) *Directory {
	if name == ".." {
		return dir.Parent
	}
	if name == "/" {
		return dir.Root()
	}
	for _, d := range dir.Children {
		if d.Name == name {
			return d
		}
	}
	return nil
}

/*
func (dir *Directory) FullName() string {
	if dir.Name == "/" {
		return dir.Name
	}

}
*/
