// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"github.com/progrium/darwinkit/objc"
)

// A collection of model data for GPU-accelerated intersection of rays with the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructure?language=objc
type PAccelerationStructure interface {
	// optional
	Size() uint
	HasSize() bool
}

// ensure impl type implements protocol interface
var _ PAccelerationStructure = (*AccelerationStructureObject)(nil)

// A concrete type for the [PAccelerationStructure] protocol.
type AccelerationStructureObject struct {
	objc.Object
}

func (a_ AccelerationStructureObject) HasSize() bool {
	return a_.RespondsToSelector(objc.Sel("size"))
}

// The size of the acceleration structure’s memory allocation, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructure/3553858-size?language=objc
func (a_ AccelerationStructureObject) Size() uint {
	rv := objc.Call[uint](a_, objc.Sel("size"))
	return rv
}