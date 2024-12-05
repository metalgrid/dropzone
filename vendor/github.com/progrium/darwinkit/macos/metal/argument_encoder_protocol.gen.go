// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// An object used to encode data into an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder?language=objc
type PArgumentEncoder interface {
	// optional
	SetSamplerStatesWithRange(samplers unsafe.Pointer, range_ foundation.Range)
	HasSetSamplerStatesWithRange() bool

	// optional
	SetTexturesWithRange(textures unsafe.Pointer, range_ foundation.Range)
	HasSetTexturesWithRange() bool

	// optional
	SetRenderPipelineStatesWithRange(pipelines unsafe.Pointer, range_ foundation.Range)
	HasSetRenderPipelineStatesWithRange() bool

	// optional
	SetComputePipelineStatesWithRange(pipelines unsafe.Pointer, range_ foundation.Range)
	HasSetComputePipelineStatesWithRange() bool

	// optional
	SetSamplerStateAtIndex(sampler SamplerStateObject, index uint)
	HasSetSamplerStateAtIndex() bool

	// optional
	SetAccelerationStructureAtIndex(accelerationStructure AccelerationStructureObject, index uint)
	HasSetAccelerationStructureAtIndex() bool

	// optional
	SetComputePipelineStateAtIndex(pipeline ComputePipelineStateObject, index uint)
	HasSetComputePipelineStateAtIndex() bool

	// optional
	ConstantDataAtIndex(index uint) unsafe.Pointer
	HasConstantDataAtIndex() bool

	// optional
	SetBuffersOffsetsWithRange(buffers unsafe.Pointer, offsets *uint, range_ foundation.Range)
	HasSetBuffersOffsetsWithRange() bool

	// optional
	SetIndirectCommandBufferAtIndex(indirectCommandBuffer IndirectCommandBufferObject, index uint)
	HasSetIndirectCommandBufferAtIndex() bool

	// optional
	SetIndirectCommandBuffersWithRange(buffers unsafe.Pointer, range_ foundation.Range)
	HasSetIndirectCommandBuffersWithRange() bool

	// optional
	SetBufferOffsetAtIndex(buffer BufferObject, offset uint, index uint)
	HasSetBufferOffsetAtIndex() bool

	// optional
	SetIntersectionFunctionTablesWithRange(intersectionFunctionTables unsafe.Pointer, range_ foundation.Range)
	HasSetIntersectionFunctionTablesWithRange() bool

	// optional
	SetVisibleFunctionTableAtIndex(visibleFunctionTable VisibleFunctionTableObject, index uint)
	HasSetVisibleFunctionTableAtIndex() bool

	// optional
	SetRenderPipelineStateAtIndex(pipeline RenderPipelineStateObject, index uint)
	HasSetRenderPipelineStateAtIndex() bool

	// optional
	SetVisibleFunctionTablesWithRange(visibleFunctionTables unsafe.Pointer, range_ foundation.Range)
	HasSetVisibleFunctionTablesWithRange() bool

	// optional
	SetArgumentBufferOffset(argumentBuffer BufferObject, offset uint)
	HasSetArgumentBufferOffset() bool

	// optional
	SetIntersectionFunctionTableAtIndex(intersectionFunctionTable IntersectionFunctionTableObject, index uint)
	HasSetIntersectionFunctionTableAtIndex() bool

	// optional
	SetArgumentBufferStartOffsetArrayElement(argumentBuffer BufferObject, startOffset uint, arrayElement uint)
	HasSetArgumentBufferStartOffsetArrayElement() bool

	// optional
	NewArgumentEncoderForBufferAtIndex(index uint) ArgumentEncoderObject
	HasNewArgumentEncoderForBufferAtIndex() bool

	// optional
	SetTextureAtIndex(texture TextureObject, index uint)
	HasSetTextureAtIndex() bool

	// optional
	Alignment() uint
	HasAlignment() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	EncodedLength() uint
	HasEncodedLength() bool

	// optional
	Device() DeviceObject
	HasDevice() bool
}

// ensure impl type implements protocol interface
var _ PArgumentEncoder = (*ArgumentEncoderObject)(nil)

// A concrete type for the [PArgumentEncoder] protocol.
type ArgumentEncoderObject struct {
	objc.Object
}

func (a_ ArgumentEncoderObject) HasSetSamplerStatesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setSamplerStates:withRange:"))
}

// Encodes an array of samplers into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915778-setsamplerstates?language=objc
func (a_ ArgumentEncoderObject) SetSamplerStatesWithRange(samplers unsafe.Pointer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", samplers)
	objc.Call[objc.Void](a_, objc.Sel("setSamplerStates:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderObject) HasSetTexturesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setTextures:withRange:"))
}

// Encodes references to an array of textures into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915786-settextures?language=objc
func (a_ ArgumentEncoderObject) SetTexturesWithRange(textures unsafe.Pointer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLTexture", textures)
	objc.Call[objc.Void](a_, objc.Sel("setTextures:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderObject) HasSetRenderPipelineStatesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setRenderPipelineStates:withRange:"))
}

// Encodes references to an array of render pipeline states into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2966536-setrenderpipelinestates?language=objc
func (a_ ArgumentEncoderObject) SetRenderPipelineStatesWithRange(pipelines unsafe.Pointer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLRenderPipelineState", pipelines)
	objc.Call[objc.Void](a_, objc.Sel("setRenderPipelineStates:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderObject) HasSetComputePipelineStatesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setComputePipelineStates:withRange:"))
}

// Encodes references to an array of compute pipeline states into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2966534-setcomputepipelinestates?language=objc
func (a_ ArgumentEncoderObject) SetComputePipelineStatesWithRange(pipelines unsafe.Pointer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLComputePipelineState", pipelines)
	objc.Call[objc.Void](a_, objc.Sel("setComputePipelineStates:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderObject) HasSetSamplerStateAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setSamplerState:atIndex:"))
}

// Encodes a sampler into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915770-setsamplerstate?language=objc
func (a_ ArgumentEncoderObject) SetSamplerStateAtIndex(sampler SamplerStateObject, index uint) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", sampler)
	objc.Call[objc.Void](a_, objc.Sel("setSamplerState:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderObject) HasSetAccelerationStructureAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setAccelerationStructure:atIndex:"))
}

// Encodes a reference to an acceleration structure into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3553920-setaccelerationstructure?language=objc
func (a_ ArgumentEncoderObject) SetAccelerationStructureAtIndex(accelerationStructure AccelerationStructureObject, index uint) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", accelerationStructure)
	objc.Call[objc.Void](a_, objc.Sel("setAccelerationStructure:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderObject) HasSetComputePipelineStateAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setComputePipelineState:atIndex:"))
}

// Encodes a reference to a compute pipeline state into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2966533-setcomputepipelinestate?language=objc
func (a_ ArgumentEncoderObject) SetComputePipelineStateAtIndex(pipeline ComputePipelineStateObject, index uint) {
	po0 := objc.WrapAsProtocol("MTLComputePipelineState", pipeline)
	objc.Call[objc.Void](a_, objc.Sel("setComputePipelineState:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderObject) HasConstantDataAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("constantDataAtIndex:"))
}

// Returns a pointer for an inlined constant data argument in the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915771-constantdataatindex?language=objc
func (a_ ArgumentEncoderObject) ConstantDataAtIndex(index uint) unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](a_, objc.Sel("constantDataAtIndex:"), index)
	return rv
}

func (a_ ArgumentEncoderObject) HasSetBuffersOffsetsWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setBuffers:offsets:withRange:"))
}

// Encodes references to an array of buffers into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915772-setbuffers?language=objc
func (a_ ArgumentEncoderObject) SetBuffersOffsetsWithRange(buffers unsafe.Pointer, offsets *uint, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffers)
	objc.Call[objc.Void](a_, objc.Sel("setBuffers:offsets:withRange:"), po0, offsets, range_)
}

func (a_ ArgumentEncoderObject) HasSetIndirectCommandBufferAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setIndirectCommandBuffer:atIndex:"))
}

// Encodes a reference to an indirect command buffer into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2967410-setindirectcommandbuffer?language=objc
func (a_ ArgumentEncoderObject) SetIndirectCommandBufferAtIndex(indirectCommandBuffer IndirectCommandBufferObject, index uint) {
	po0 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", indirectCommandBuffer)
	objc.Call[objc.Void](a_, objc.Sel("setIndirectCommandBuffer:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderObject) HasSetIndirectCommandBuffersWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setIndirectCommandBuffers:withRange:"))
}

// Encodes an array of indirect command buffers into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2967411-setindirectcommandbuffers?language=objc
func (a_ ArgumentEncoderObject) SetIndirectCommandBuffersWithRange(buffers unsafe.Pointer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", buffers)
	objc.Call[objc.Void](a_, objc.Sel("setIndirectCommandBuffers:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderObject) HasSetBufferOffsetAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setBuffer:offset:atIndex:"))
}

// Encodes a reference to a buffer into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915785-setbuffer?language=objc
func (a_ ArgumentEncoderObject) SetBufferOffsetAtIndex(buffer BufferObject, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](a_, objc.Sel("setBuffer:offset:atIndex:"), po0, offset, index)
}

func (a_ ArgumentEncoderObject) HasSetIntersectionFunctionTablesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setIntersectionFunctionTables:withRange:"))
}

// Encodes references to an array of intersection function tables into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3608167-setintersectionfunctiontables?language=objc
func (a_ ArgumentEncoderObject) SetIntersectionFunctionTablesWithRange(intersectionFunctionTables unsafe.Pointer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLIntersectionFunctionTable", intersectionFunctionTables)
	objc.Call[objc.Void](a_, objc.Sel("setIntersectionFunctionTables:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderObject) HasSetVisibleFunctionTableAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setVisibleFunctionTable:atIndex:"))
}

// Encodes a reference to a function table into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3608168-setvisiblefunctiontable?language=objc
func (a_ ArgumentEncoderObject) SetVisibleFunctionTableAtIndex(visibleFunctionTable VisibleFunctionTableObject, index uint) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", visibleFunctionTable)
	objc.Call[objc.Void](a_, objc.Sel("setVisibleFunctionTable:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderObject) HasSetRenderPipelineStateAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setRenderPipelineState:atIndex:"))
}

// Encodes a reference to a render pipeline state into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2966535-setrenderpipelinestate?language=objc
func (a_ ArgumentEncoderObject) SetRenderPipelineStateAtIndex(pipeline RenderPipelineStateObject, index uint) {
	po0 := objc.WrapAsProtocol("MTLRenderPipelineState", pipeline)
	objc.Call[objc.Void](a_, objc.Sel("setRenderPipelineState:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderObject) HasSetVisibleFunctionTablesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setVisibleFunctionTables:withRange:"))
}

// Encodes references to an array of visible function tables into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3608169-setvisiblefunctiontables?language=objc
func (a_ ArgumentEncoderObject) SetVisibleFunctionTablesWithRange(visibleFunctionTables unsafe.Pointer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", visibleFunctionTables)
	objc.Call[objc.Void](a_, objc.Sel("setVisibleFunctionTables:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderObject) HasSetArgumentBufferOffset() bool {
	return a_.RespondsToSelector(objc.Sel("setArgumentBuffer:offset:"))
}

// Specifies the position in a buffer where the encoder writes argument data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915777-setargumentbuffer?language=objc
func (a_ ArgumentEncoderObject) SetArgumentBufferOffset(argumentBuffer BufferObject, offset uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", argumentBuffer)
	objc.Call[objc.Void](a_, objc.Sel("setArgumentBuffer:offset:"), po0, offset)
}

func (a_ ArgumentEncoderObject) HasSetIntersectionFunctionTableAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setIntersectionFunctionTable:atIndex:"))
}

// Encodes a reference to a ray-tracing intersection function table into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3608166-setintersectionfunctiontable?language=objc
func (a_ ArgumentEncoderObject) SetIntersectionFunctionTableAtIndex(intersectionFunctionTable IntersectionFunctionTableObject, index uint) {
	po0 := objc.WrapAsProtocol("MTLIntersectionFunctionTable", intersectionFunctionTable)
	objc.Call[objc.Void](a_, objc.Sel("setIntersectionFunctionTable:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderObject) HasSetArgumentBufferStartOffsetArrayElement() bool {
	return a_.RespondsToSelector(objc.Sel("setArgumentBuffer:startOffset:arrayElement:"))
}

// Specifies an array element within a buffer where the encoder writes argument data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915780-setargumentbuffer?language=objc
func (a_ ArgumentEncoderObject) SetArgumentBufferStartOffsetArrayElement(argumentBuffer BufferObject, startOffset uint, arrayElement uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", argumentBuffer)
	objc.Call[objc.Void](a_, objc.Sel("setArgumentBuffer:startOffset:arrayElement:"), po0, startOffset, arrayElement)
}

func (a_ ArgumentEncoderObject) HasNewArgumentEncoderForBufferAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("newArgumentEncoderForBufferAtIndex:"))
}

// Creates a new argument encoder for a nested argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915783-newargumentencoderforbufferatind?language=objc
func (a_ ArgumentEncoderObject) NewArgumentEncoderForBufferAtIndex(index uint) ArgumentEncoderObject {
	rv := objc.Call[ArgumentEncoderObject](a_, objc.Sel("newArgumentEncoderForBufferAtIndex:"), index)
	return rv
}

func (a_ ArgumentEncoderObject) HasSetTextureAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setTexture:atIndex:"))
}

// Encodes a reference to a texture into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915776-settexture?language=objc
func (a_ ArgumentEncoderObject) SetTextureAtIndex(texture TextureObject, index uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](a_, objc.Sel("setTexture:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderObject) HasAlignment() bool {
	return a_.RespondsToSelector(objc.Sel("alignment"))
}

// The alignment, in bytes, required for storing the encoded resources of an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915775-alignment?language=objc
func (a_ ArgumentEncoderObject) Alignment() uint {
	rv := objc.Call[uint](a_, objc.Sel("alignment"))
	return rv
}

func (a_ ArgumentEncoderObject) HasSetLabel() bool {
	return a_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915773-label?language=objc
func (a_ ArgumentEncoderObject) SetLabel(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setLabel:"), value)
}

func (a_ ArgumentEncoderObject) HasLabel() bool {
	return a_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915773-label?language=objc
func (a_ ArgumentEncoderObject) Label() string {
	rv := objc.Call[string](a_, objc.Sel("label"))
	return rv
}

func (a_ ArgumentEncoderObject) HasEncodedLength() bool {
	return a_.RespondsToSelector(objc.Sel("encodedLength"))
}

// The number of bytes required to store the encoded resources of an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915784-encodedlength?language=objc
func (a_ ArgumentEncoderObject) EncodedLength() uint {
	rv := objc.Call[uint](a_, objc.Sel("encodedLength"))
	return rv
}

func (a_ ArgumentEncoderObject) HasDevice() bool {
	return a_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the argument encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915779-device?language=objc
func (a_ ArgumentEncoderObject) Device() DeviceObject {
	rv := objc.Call[DeviceObject](a_, objc.Sel("device"))
	return rv
}