// Code generated by DarwinKit. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MultiArrayConstraint] class.
var MultiArrayConstraintClass = _MultiArrayConstraintClass{objc.GetClass("MLMultiArrayConstraint")}

type _MultiArrayConstraintClass struct {
	objc.Class
}

// An interface definition for the [MultiArrayConstraint] class.
type IMultiArrayConstraint interface {
	objc.IObject
	DataType() MultiArrayDataType
	ShapeConstraint() MultiArrayShapeConstraint
	Shape() []foundation.Number
}

// The shape and data type constraints for a multidimensional array feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint?language=objc
type MultiArrayConstraint struct {
	objc.Object
}

func MultiArrayConstraintFrom(ptr unsafe.Pointer) MultiArrayConstraint {
	return MultiArrayConstraint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MultiArrayConstraintClass) Alloc() MultiArrayConstraint {
	rv := objc.Call[MultiArrayConstraint](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MultiArrayConstraintClass) New() MultiArrayConstraint {
	rv := objc.Call[MultiArrayConstraint](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMultiArrayConstraint() MultiArrayConstraint {
	return MultiArrayConstraintClass.New()
}

func (m_ MultiArrayConstraint) Init() MultiArrayConstraint {
	rv := objc.Call[MultiArrayConstraint](m_, objc.Sel("init"))
	return rv
}

// The type for the multi array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/2921278-datatype?language=objc
func (m_ MultiArrayConstraint) DataType() MultiArrayDataType {
	rv := objc.Call[MultiArrayDataType](m_, objc.Sel("dataType"))
	return rv
}

// The constraint on the shape of the multiarray. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/2994311-shapeconstraint?language=objc
func (m_ MultiArrayConstraint) ShapeConstraint() MultiArrayShapeConstraint {
	rv := objc.Call[MultiArrayShapeConstraint](m_, objc.Sel("shapeConstraint"))
	return rv
}

// The shape of the multi array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/2921277-shape?language=objc
func (m_ MultiArrayConstraint) Shape() []foundation.Number {
	rv := objc.Call[[]foundation.Number](m_, objc.Sel("shape"))
	return rv
}