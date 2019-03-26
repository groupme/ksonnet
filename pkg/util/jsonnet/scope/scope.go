package scope

import (
	"github.com/ksonnet/ksonnet/pkg/util/jsonnet"
	"github.com/spf13/afero"
)

var (
	componentJPaths  = make([]string, 0)
	componentExtVars = make(map[string]string)
	componentTlaVars = make(map[string]string)
)

// AddExtVar adds an ext var to a component evaluation.
func AddExtVar(key, value string) {
	componentExtVars[key] = value
}

// AddExtVarFile adds an ext var from a file to component evaluation.
func AddExtVarFile(fs afero.Fs, key, filePath string) error {
	data, err := afero.ReadFile(fs, filePath)
	if err != nil {
		return err
	}

	componentExtVars[key] = string(data)
	return nil
}

// AddTlaVar adds a tla var to a component evaluation.
func AddTlaVar(key, value string) {
	componentTlaVars[key] = value
}

// AddTlaVarFile adds a tla var from a file to component evaluation.
func AddTlaVarFile(fs afero.Fs, key, filePath string) error {
	data, err := afero.ReadFile(fs, filePath)
	if err != nil {
		return err
	}

	componentTlaVars[key] = string(data)
	return nil
}

func Apply(vm *jsonnet.VM) *jsonnet.VM {

	for k, v := range componentExtVars {
		vm.ExtVar(k, v)
	}

	for k, v := range componentTlaVars {
		vm.TLAVar(k, v)
	}

	return vm
}
