package fs

type Entry interface {
	IsFile() bool

	IsFolder() bool

	IfFile(e func(f File))

	IfFolder(e func(f Folder))
}
