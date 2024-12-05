// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PDFPanel] class.
var PDFPanelClass = _PDFPanelClass{objc.GetClass("NSPDFPanel")}

type _PDFPanelClass struct {
	objc.Class
}

// An interface definition for the [PDFPanel] class.
type IPDFPanel interface {
	objc.IObject
	BeginSheetWithPDFInfoModalForWindowCompletionHandler(pdfInfo IPDFInfo, docWindow IWindow, completionHandler func(arg0 int))
	Options() PDFPanelOptions
	SetOptions(value PDFPanelOptions)
	DefaultFileName() string
	SetDefaultFileName(value string)
	AccessoryController() ViewController
	SetAccessoryController(value IViewController)
}

// A Save or Export as PDF panel that’s consistent with the macOS user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel?language=objc
type PDFPanel struct {
	objc.Object
}

func PDFPanelFrom(ptr unsafe.Pointer) PDFPanel {
	return PDFPanel{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PDFPanelClass) Alloc() PDFPanel {
	rv := objc.Call[PDFPanel](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PDFPanelClass) New() PDFPanel {
	rv := objc.Call[PDFPanel](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPDFPanel() PDFPanel {
	return PDFPanelClass.New()
}

func (p_ PDFPanel) Init() PDFPanel {
	rv := objc.Call[PDFPanel](p_, objc.Sel("init"))
	return rv
}

// Returns a new NSPDFPanel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/1577141-panel?language=objc
func (pc _PDFPanelClass) Panel() PDFPanel {
	rv := objc.Call[PDFPanel](pc, objc.Sel("panel"))
	return rv
}

// Returns a new NSPDFPanel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/1577141-panel?language=objc
func PDFPanel_Panel() PDFPanel {
	return PDFPanelClass.Panel()
}

// Presents a document-modal PDF panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/1529098-beginsheetwithpdfinfo?language=objc
func (p_ PDFPanel) BeginSheetWithPDFInfoModalForWindowCompletionHandler(pdfInfo IPDFInfo, docWindow IWindow, completionHandler func(arg0 int)) {
	objc.Call[objc.Void](p_, objc.Sel("beginSheetWithPDFInfo:modalForWindow:completionHandler:"), pdfInfo, docWindow, completionHandler)
}

// A set of configuration options that determine the accessory views the PDF panel should display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/1532479-options?language=objc
func (p_ PDFPanel) Options() PDFPanelOptions {
	rv := objc.Call[PDFPanelOptions](p_, objc.Sel("options"))
	return rv
}

// A set of configuration options that determine the accessory views the PDF panel should display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/1532479-options?language=objc
func (p_ PDFPanel) SetOptions(value PDFPanelOptions) {
	objc.Call[objc.Void](p_, objc.Sel("setOptions:"), value)
}

// The initial value for the user-editable filename shown in the name field of the PDF panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/1532720-defaultfilename?language=objc
func (p_ PDFPanel) DefaultFileName() string {
	rv := objc.Call[string](p_, objc.Sel("defaultFileName"))
	return rv
}

// The initial value for the user-editable filename shown in the name field of the PDF panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/1532720-defaultfilename?language=objc
func (p_ PDFPanel) SetDefaultFileName(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setDefaultFileName:"), value)
}

// A view controller for the accessory view that the panel can present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/1524637-accessorycontroller?language=objc
func (p_ PDFPanel) AccessoryController() ViewController {
	rv := objc.Call[ViewController](p_, objc.Sel("accessoryController"))
	return rv
}

// A view controller for the accessory view that the panel can present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/1524637-accessorycontroller?language=objc
func (p_ PDFPanel) SetAccessoryController(value IViewController) {
	objc.Call[objc.Void](p_, objc.Sel("setAccessoryController:"), value)
}