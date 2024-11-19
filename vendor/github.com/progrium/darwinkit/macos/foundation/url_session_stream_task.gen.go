// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [URLSessionStreamTask] class.
var URLSessionStreamTaskClass = _URLSessionStreamTaskClass{objc.GetClass("NSURLSessionStreamTask")}

type _URLSessionStreamTaskClass struct {
	objc.Class
}

// An interface definition for the [URLSessionStreamTask] class.
type IURLSessionStreamTask interface {
	IURLSessionTask
	StartSecureConnection()
	CloseWrite()
	CaptureStreams()
	ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler(minBytes uint, maxBytes uint, timeout TimeInterval, completionHandler func(data []byte, atEOF bool, error Error))
	CloseRead()
	WriteDataTimeoutCompletionHandler(data []byte, timeout TimeInterval, completionHandler func(error Error))
}

// A URL session task that is stream-based. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamtask?language=objc
type URLSessionStreamTask struct {
	URLSessionTask
}

func URLSessionStreamTaskFrom(ptr unsafe.Pointer) URLSessionStreamTask {
	return URLSessionStreamTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}

func (uc _URLSessionStreamTaskClass) Alloc() URLSessionStreamTask {
	rv := objc.Call[URLSessionStreamTask](uc, objc.Sel("alloc"))
	return rv
}

func (uc _URLSessionStreamTaskClass) New() URLSessionStreamTask {
	rv := objc.Call[URLSessionStreamTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionStreamTask() URLSessionStreamTask {
	return URLSessionStreamTaskClass.New()
}

func (u_ URLSessionStreamTask) Init() URLSessionStreamTask {
	rv := objc.Call[URLSessionStreamTask](u_, objc.Sel("init"))
	return rv
}

// Completes any enqueued reads and writes, and establishes a secure connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamtask/1411567-startsecureconnection?language=objc
func (u_ URLSessionStreamTask) StartSecureConnection() {
	objc.Call[objc.Void](u_, objc.Sel("startSecureConnection"))
}

// Completes any enqueued reads and writes, and then closes the write side of the underlying socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamtask/1411347-closewrite?language=objc
func (u_ URLSessionStreamTask) CloseWrite() {
	objc.Call[objc.Void](u_, objc.Sel("closeWrite"))
}

// Completes any already enqueued reads and writes, and then invokes the [foundation/nsurlsessionstreamdelegate/urlsession] delegate message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamtask/1410132-capturestreams?language=objc
func (u_ URLSessionStreamTask) CaptureStreams() {
	objc.Call[objc.Void](u_, objc.Sel("captureStreams"))
}

// Asynchronously reads a number of bytes from the stream, and calls a handler upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamtask/1411604-readdataofminlength?language=objc
func (u_ URLSessionStreamTask) ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler(minBytes uint, maxBytes uint, timeout TimeInterval, completionHandler func(data []byte, atEOF bool, error Error)) {
	objc.Call[objc.Void](u_, objc.Sel("readDataOfMinLength:maxLength:timeout:completionHandler:"), minBytes, maxBytes, timeout, completionHandler)
}

// Completes any enqueued reads and writes, and then closes the read side of the underlying socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamtask/1411558-closeread?language=objc
func (u_ URLSessionStreamTask) CloseRead() {
	objc.Call[objc.Void](u_, objc.Sel("closeRead"))
}

// Asynchronously writes the specified data to the stream, and calls a handler upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionstreamtask/1411602-writedata?language=objc
func (u_ URLSessionStreamTask) WriteDataTimeoutCompletionHandler(data []byte, timeout TimeInterval, completionHandler func(error Error)) {
	objc.Call[objc.Void](u_, objc.Sel("writeData:timeout:completionHandler:"), data, timeout, completionHandler)
}