package descriptor

import (
	"fmt"
	"testing"
)

func TestRegistry_SetupInput(t *testing.T) {
	type args struct {
		inputDirPath  string
		pkgName       string
		dirOrFilePath string
	}
	type want struct {
		numFiles int
		path     string
		err      error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Dir Input",
			args: args{
				pkgName:       "testpkg",
				dirOrFilePath: "../examples",
			},
			want: want{
				numFiles: 3,
				path:     "../examples",
				err:      nil,
			},
		},
		{
			name: "File Input",
			args: args{
				pkgName:       "testpkg",
				dirOrFilePath: "../examples/family_list.json",
			},
			want: want{
				numFiles: 1,
				path:     "../examples",
				err:      nil,
			},
		},
		{
			name: "Invalid Path Error",
			args: args{
				pkgName:       "testpkg",
				dirOrFilePath: "../example",
			},
			want: want{
				numFiles: 0,
				path:     "",
				err:      fmt.Errorf("path [%s] does not exist", "../example"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := NewRegistry(tt.args.pkgName)
			err := reg.SetupInput(tt.args.dirOrFilePath)
			if err != nil {
				if err.Error() != tt.want.err.Error() {
					t.Errorf("SetupInput() error = %v, wantErr %v", err, tt.want.err)
				}
			}
			if reg.inputDirPath != tt.want.path {
				t.Errorf("SetupInput() inputDirPath: %s, want %v", reg.inputDirPath, tt.want.path)
			}
			if len(reg.files) != tt.want.numFiles {
				t.Errorf("SetupInput() len(reg.files): %d, want %d", len(reg.files), tt.want.numFiles)
			}
		})
	}
}

func TestRegistry_Load(t *testing.T) {
	type args struct {
		inputDirPath  string
		pkgName       string
		dirOrFilePath string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "Dir Input",
			args: args{
				pkgName:       "testpkg",
				dirOrFilePath: "../examples",
			},
			want:    3,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := NewRegistry(tt.args.pkgName)
			_ = reg.SetupInput(tt.args.dirOrFilePath)
			actual, err := reg.Load()
			if err != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(actual) != tt.want {
				t.Errorf("Load() load %v files, want %v", len(actual), tt.want)
			}
		})
	}
}
