package doubles

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
	"os"
)

// NewFileSystemStub returns a new fs.FS stub with provided files.
func NewFileSystemStub(files map[string][]byte) *FilesystemStub {
	return &FilesystemStub{
		files:   files,
		badFile: "badfile.txt",
	}
}

// FilesystemStub implements type FS interface
type FilesystemStub struct {
	files   map[string][]byte
	badFile string
}

// Open opens the named file
func (f *FilesystemStub) Open(name string) (fs.File, error) {
	if name == f.badFile {
		return &badFile{}, nil
	}

	tmpfile, err := ioutil.TempFile("", "fsdemo")
	if err != nil {
		log.Fatal(err)
	}

	contents, ok := f.files[name]
	if !ok {
		return nil, os.ErrNotExist
	}

	_, _ = tmpfile.Write([]byte(contents))
	_, _ = tmpfile.Seek(0, 0)

	return tmpfile, nil
}

type badFile struct{}

func (f *badFile) Stat() (fs.FileInfo, error) {
	return nil, errors.New("bad file")
}

func (f *badFile) Read(_ []byte) (int, error) {
	return 0, errors.New("bad file")
}

func (f *badFile) Close() error {
	return errors.New("bad file")
}
