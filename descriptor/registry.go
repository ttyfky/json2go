package descriptor

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const defaultFileNum = 5

// Registry is a registry of information.
type Registry struct {
	// inputDirPath is path to directory where json is in.
	inputDirPath string

	// files is files to transform to Go file.
	files []string

	// pkgName is package Name of generating file.
	pkgName string

	// outputPath is where generated Go file is put.
	outputPath string

	// writer is a common writer of output. Null if outputPath is set.
	writer io.Writer
}

// NewRegistry returns a new Registry.
func NewRegistry(pkg string) *Registry {
	return &Registry{
		files:   make([]string, 0, defaultFileNum),
		pkgName: pkg,
	}
}

// SetInputDirPath set input dir path to Registry.
func (reg *Registry) SetInputDirPath(path string) {
	reg.inputDirPath = path
}

// SetFile set input file path to Registry.
// This method also extract input dir path which is prefix of path to the file.
func (reg *Registry) SetFile(filename string) {
	reg.files = append(reg.files, filename)
}

// SetFiles set files to read to Registry.
func (reg *Registry) SetFiles(f []string) {
	reg.files = f
}

// SetPackageName set package name of generated models.
func (reg *Registry) SetPackageName(pkg string) {
	reg.pkgName = pkg
}

// SetOutputPath set package name of generated models.
func (reg *Registry) SetOutputPath(path string) {
	reg.outputPath = path
}

// GetOutputPath set package name of generated models.
func (reg *Registry) GetOutputPath() string {
	return reg.outputPath
}

// SetWriter set writer applicable for all files.
func (reg *Registry) SetWriter(w io.Writer) {
	reg.writer = w
}

// GetWriter return writer.
func (reg *Registry) GetWriter() io.Writer {
	return reg.writer
}

// SetFilesInDir walks through files in inputDirPath and add JSON file to files list.
func (reg *Registry) SetFilesInDir(path string) error {
	list, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	files := make([]string, 0, len(list))
	for _, d := range list {
		if strings.HasSuffix(d.Name(), ".json") {
			files = append(files, d.Name())
		}
	}
	if len(files) == 0 {
		return fmt.Errorf("no json file was found on the path [%s]", path)
	}
	reg.SetInputDirPath(path)
	reg.SetFiles(files)
	return nil
}

// Load walks through all files having suffix ".json" in the
// directory specified by path and returns a list of descriptor.File.
func (reg *Registry) Load() ([]*File, error) {
	files := make([]*File, 0, len(reg.files))
	for _, fname := range reg.files {
		in, err := unmarshal(reg.inputDirPath + "/" + fname)
		if err != nil {
			return nil, err
		}
		dot := strings.LastIndex(fname, ".")
		files = append(files,
			&File{
				Name:      fname[:dot],
				PkgName:   reg.pkgName,
				RawFields: in,
			})
	}
	return files, nil
}

func unmarshal(filepath string) (map[string]interface{}, error) {
	var input map[string]interface{}
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &input)
	if err != nil {
		return nil, err
	}
	return input, nil
}

// SetupInput verify path of input and handle based on if it's file path or dir path.
func (reg *Registry) SetupInput(dirOrFilePath string) error {
	var fi os.FileInfo
	var err error
	if fi, err = os.Stat(dirOrFilePath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path [%s] does not exist", dirOrFilePath)
		}
		return err
	}
	if fi.IsDir() {
		err := reg.SetFilesInDir(dirOrFilePath)
		if err != nil {
			return err
		}
	} else {
		reg.SetInputDirPath(path.Dir(dirOrFilePath))
		reg.SetFile(path.Base(dirOrFilePath))
	}
	return nil
}
