package generator

import (
	"github.com/ttyfky/json2go/descriptor"
)

// Generator is an abstraction of code generators.
type Generator interface {
	// Generate generates output files from input json files.
	Generate(targets []*descriptor.File) error
}
