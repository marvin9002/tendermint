// Code generated by github.com/Liamsi/protogenerat0r DO NOT EDIT.

package types

import (
	"os"

	// we use dedis' protobuf library to generate proto files from go-structs
	// see: https://github.com/dedis/protobuf#generating-proto-files
	"github.com/dedis/protobuf"
)

var structTypes = []interface{}{
	PartSetHeader{},
	//SignedHeader{},
	Vote{},
	//VoteSet{},
}

// Call this method to generate protobuf messages:
func GenerateProtos() {
	// see: https://github.com/dedis/protobuf#generating-proto-files
	protobuf.GenerateProtobufDefinition(os.Stdout, structTypes, nil, nil)
}