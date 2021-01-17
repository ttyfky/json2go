package gengo

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ttyfky/json2go/descriptor"
)

func Test_generationHandle_Generate(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		pathToWant string
	}{
		{
			name:       "Generate",
			args:       args{path: "./testdata/json/company.json"},
			pathToWant: "./testdata/generated/company.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reg *descriptor.Registry
			reg = descriptor.NewRegistry("testdata")
			actual := &bytes.Buffer{}
			reg.SetWriter(actual)
			_ = reg.SetupInput(tt.args.path)
			files, err := reg.Load()
			if err != nil {
				t.Errorf("unexpected error %s", err)
			}
			g := &generationHandle{
				reg: reg,
			}
			err = g.Generate(files)
			if err != nil {
				t.Errorf("unexpected error %s", err)
			}
			want, _ := load(tt.pathToWant)
			if diff := cmp.Diff(actual.String(), string(want)); diff != "" {
				t.Errorf("Diff in generated: %s", diff)
			}
		})
	}
}

func load(filepath string) (testInput []byte, err error) {
	testInput, err = ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return testInput, nil
}
