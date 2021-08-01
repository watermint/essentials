package fs

type FileSystem interface {
	// Exists tests whether the file or folder exists.
	Exists(path Path) ExistsResult

	// Delete a file or a folder includes sub-folders.
	Delete(path Path) DeleteResult

	// CreateFolder creates a folder including any necessary but nonexistent parent folders.
	CreateFolder(path Path) CreateFolderResult

	// CreateFile creates a file with an option.
	CreateFile(path Path, opts ...CreateFileOpt) CreateFileOpt

	// Entries returns entries of the folder of the path.
	Entries(path Path) EntriesResult
}

type CreateFileOpts struct {
}

type CreateFileOpt func(opts CreateFileOpts) CreateFileOpts

type ExistsResult interface {
	IfSuccess(f func(exists bool)) ExistsResult
	IfFailure(f func(err FileSystemError)) ExistsResult
}

type DeleteResult interface {
	IfSuccess(f func()) DeleteResult
	IfFailure(f func(err FileSystemError)) DeleteResult
}

type CreateFolderResult interface {
	IfSuccess(f func(folder Folder)) CreateFolderResult
	IfFailure(f func(err FileSystemError)) CreateFolderResult
}

type CreateFileResult interface {
	IfSuccess(f func(file File)) CreateFileResult
	IfFailure(f func(err FileSystemError)) CreateFileResult
}

type EntriesResult interface {
	// IfSuccess Callback per entry.
	IfSuccess(f func(entry Entry)) EntriesResult
	IfFailure(f func(err FileSystemError)) EntriesResult
}

type FileSystemAsync interface {
	Exists(path Path, onComplete func(c ExistsResult))

	Delete(path Path, onComplete func(c DeleteResult))

	CreateFolder(path Path, onComplete func(c CreateFolderResult))
}

type FileSystemError interface {
	error

	// True when permission is denied.
	IsPermissionDenied() bool

	// True when the operation timeout.
	IsTimeout() bool
}
