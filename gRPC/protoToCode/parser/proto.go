package parser

import "github.com/emicklei/proto"

// Proto describes a proto file,
type Proto struct {
	Src       string
	Name      string
	Package   Package
	PbPackage string
	GoPackage string
	Import    []Import
	Message   []Message
	Service   Services
}

type (
	// Message embeds proto.Message
	Message struct {
		*proto.Message
	}

	// Import embeds proto.Import
	Import struct {
		*proto.Import
	}

	// Package defines the protobuf package.
	Package struct {
		*proto.Package
	}
)
