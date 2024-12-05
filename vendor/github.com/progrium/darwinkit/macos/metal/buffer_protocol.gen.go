// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A resource that stores data in a format defined by your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer?language=objc
type PBuffer interface {
	// optional
	Contents() unsafe.Pointer
	HasContents() bool

	// optional
	NewRemoteBufferViewForDevice(device DeviceObject) BufferObject
	HasNewRemoteBufferViewForDevice() bool

	// optional
	AddDebugMarkerRange(marker string, range_ foundation.Range)
	HasAddDebugMarkerRange() bool

	// optional
	NewTextureWithDescriptorOffsetBytesPerRow(descriptor TextureDescriptor, offset uint, bytesPerRow uint) TextureObject
	HasNewTextureWithDescriptorOffsetBytesPerRow() bool

	// optional
	RemoveAllDebugMarkers()
	HasRemoveAllDebugMarkers() bool

	// optional
	DidModifyRange(range_ foundation.Range)
	HasDidModifyRange() bool

	// optional
	RemoteStorageBuffer() BufferObject
	HasRemoteStorageBuffer() bool

	// optional
	Length() uint
	HasLength() bool
}

// ensure impl type implements protocol interface
var _ PBuffer = (*BufferObject)(nil)

// A concrete type for the [PBuffer] protocol.
type BufferObject struct {
	objc.Object
}

func (b_ BufferObject) HasContents() bool {
	return b_.RespondsToSelector(objc.Sel("contents"))
}

// Gets the system address of the buffer’s storage allocation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1515716-contents?language=objc
func (b_ BufferObject) Contents() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](b_, objc.Sel("contents"))
	return rv
}

func (b_ BufferObject) HasNewRemoteBufferViewForDevice() bool {
	return b_.RespondsToSelector(objc.Sel("newRemoteBufferViewForDevice:"))
}

// Creates a remote view of the buffer for another GPU in the same peer group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/2967415-newremotebufferviewfordevice?language=objc
func (b_ BufferObject) NewRemoteBufferViewForDevice(device DeviceObject) BufferObject {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[BufferObject](b_, objc.Sel("newRemoteBufferViewForDevice:"), po0)
	return rv
}

func (b_ BufferObject) HasAddDebugMarkerRange() bool {
	return b_.RespondsToSelector(objc.Sel("addDebugMarker:range:"))
}

// Adds a debug marker string to a specific buffer range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1779576-adddebugmarker?language=objc
func (b_ BufferObject) AddDebugMarkerRange(marker string, range_ foundation.Range) {
	objc.Call[objc.Void](b_, objc.Sel("addDebugMarker:range:"), marker, range_)
}

func (b_ BufferObject) HasNewTextureWithDescriptorOffsetBytesPerRow() bool {
	return b_.RespondsToSelector(objc.Sel("newTextureWithDescriptor:offset:bytesPerRow:"))
}

// Creates a texture that shares its storage with the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1613852-newtexturewithdescriptor?language=objc
func (b_ BufferObject) NewTextureWithDescriptorOffsetBytesPerRow(descriptor TextureDescriptor, offset uint, bytesPerRow uint) TextureObject {
	rv := objc.Call[TextureObject](b_, objc.Sel("newTextureWithDescriptor:offset:bytesPerRow:"), descriptor, offset, bytesPerRow)
	return rv
}

func (b_ BufferObject) HasRemoveAllDebugMarkers() bool {
	return b_.RespondsToSelector(objc.Sel("removeAllDebugMarkers"))
}

// Removes all debug marker strings from the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1779577-removealldebugmarkers?language=objc
func (b_ BufferObject) RemoveAllDebugMarkers() {
	objc.Call[objc.Void](b_, objc.Sel("removeAllDebugMarkers"))
}

func (b_ BufferObject) HasDidModifyRange() bool {
	return b_.RespondsToSelector(objc.Sel("didModifyRange:"))
}

// Informs the GPU that the CPU has modified a section of the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1516121-didmodifyrange?language=objc
func (b_ BufferObject) DidModifyRange(range_ foundation.Range) {
	objc.Call[objc.Void](b_, objc.Sel("didModifyRange:"), range_)
}

func (b_ BufferObject) HasRemoteStorageBuffer() bool {
	return b_.RespondsToSelector(objc.Sel("remoteStorageBuffer"))
}

// The buffer on another GPU that the buffer was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/2967416-remotestoragebuffer?language=objc
func (b_ BufferObject) RemoteStorageBuffer() BufferObject {
	rv := objc.Call[BufferObject](b_, objc.Sel("remoteStorageBuffer"))
	return rv
}

func (b_ BufferObject) HasLength() bool {
	return b_.RespondsToSelector(objc.Sel("length"))
}

// The logical size of the buffer, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbuffer/1515373-length?language=objc
func (b_ BufferObject) Length() uint {
	rv := objc.Call[uint](b_, objc.Sel("length"))
	return rv
}