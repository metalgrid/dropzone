// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ButtonCell] class.
var ButtonCellClass = _ButtonCellClass{objc.GetClass("NSButtonCell")}

type _ButtonCellClass struct {
	objc.Class
}

// An interface definition for the [ButtonCell] class.
type IButtonCell interface {
	IActionCell
	MouseExited(event IEvent)
	SetPeriodicDelayInterval(delay float32, interval float32)
	MouseEntered(event IEvent)
	SetButtonType(type_ ButtonType)
	DrawBezelWithFrameInView(frame foundation.Rect, controlView IView)
	DrawTitleWithFrameInView(title foundation.IAttributedString, frame foundation.Rect, controlView IView) foundation.Rect
	DrawImageWithFrameInView(image IImage, frame foundation.Rect, controlView IView)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	AlternateTitle() string
	SetAlternateTitle(value string)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	SetKeyEquivalent(value string)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	AlternateImage() Image
	SetAlternateImage(value IImage)
	ImageDimsWhenDisabled() bool
	SetImageDimsWhenDisabled(value bool)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	Sound() Sound
	SetSound(value ISound)
	HighlightsBy() CellStyleMask
	SetHighlightsBy(value CellStyleMask)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	IsTransparent() bool
	SetTransparent(value bool)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.IAttributedString)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	ShowsStateBy() CellStyleMask
	SetShowsStateBy(value CellStyleMask)
}

// An object that defines the user interface of a button or other clickable region of a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell?language=objc
type ButtonCell struct {
	ActionCell
}

func ButtonCellFrom(ptr unsafe.Pointer) ButtonCell {
	return ButtonCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (b_ ButtonCell) InitImageCell(image IImage) ButtonCell {
	rv := objc.Call[ButtonCell](b_, objc.Sel("initImageCell:"), image)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1639152-initimagecell?language=objc
func NewButtonCellImageCell(image IImage) ButtonCell {
	instance := ButtonCellClass.Alloc().InitImageCell(image)
	instance.Autorelease()
	return instance
}

func (b_ ButtonCell) InitTextCell(string_ string) ButtonCell {
	rv := objc.Call[ButtonCell](b_, objc.Sel("initTextCell:"), string_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1639134-inittextcell?language=objc
func NewButtonCellTextCell(string_ string) ButtonCell {
	instance := ButtonCellClass.Alloc().InitTextCell(string_)
	instance.Autorelease()
	return instance
}

func (bc _ButtonCellClass) Alloc() ButtonCell {
	rv := objc.Call[ButtonCell](bc, objc.Sel("alloc"))
	return rv
}

func (bc _ButtonCellClass) New() ButtonCell {
	rv := objc.Call[ButtonCell](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewButtonCell() ButtonCell {
	return ButtonCellClass.New()
}

func (b_ ButtonCell) Init() ButtonCell {
	rv := objc.Call[ButtonCell](b_, objc.Sel("init"))
	return rv
}

// Erases the button’s border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1530776-mouseexited?language=objc
func (b_ ButtonCell) MouseExited(event IEvent) {
	objc.Call[objc.Void](b_, objc.Sel("mouseExited:"), event)
}

// Sets the message delay and interval for the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1525725-setperiodicdelay?language=objc
func (b_ ButtonCell) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Call[objc.Void](b_, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}

// Draws the button’s border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1524997-mouseentered?language=objc
func (b_ ButtonCell) MouseEntered(event IEvent) {
	objc.Call[objc.Void](b_, objc.Sel("mouseEntered:"), event)
}

// Sets how the button highlights while pressed and how it shows its state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527474-setbuttontype?language=objc
func (b_ ButtonCell) SetButtonType(type_ ButtonType) {
	objc.Call[objc.Void](b_, objc.Sel("setButtonType:"), type_)
}

// Draws the border of the button using the current bezel style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1524939-drawbezelwithframe?language=objc
func (b_ ButtonCell) DrawBezelWithFrameInView(frame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](b_, objc.Sel("drawBezelWithFrame:inView:"), frame, controlView)
}

// Draws the button’s title centered vertically in a specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528861-drawtitle?language=objc
func (b_ ButtonCell) DrawTitleWithFrameInView(title foundation.IAttributedString, frame foundation.Rect, controlView IView) foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("drawTitle:withFrame:inView:"), title, frame, controlView)
	return rv
}

// Draws the image associated with the button’s current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1531792-drawimage?language=objc
func (b_ ButtonCell) DrawImageWithFrameInView(image IImage, frame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](b_, objc.Sel("drawImage:withFrame:inView:"), image, frame, controlView)
}

// The mask that identifies the modifier keys for the button's key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528315-keyequivalentmodifiermask?language=objc
func (b_ ButtonCell) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Call[EventModifierFlags](b_, objc.Sel("keyEquivalentModifierMask"))
	return rv
}

// The mask that identifies the modifier keys for the button's key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528315-keyequivalentmodifiermask?language=objc
func (b_ ButtonCell) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Call[objc.Void](b_, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

// The string displayed by the button when it’s in its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1535382-alternatetitle?language=objc
func (b_ ButtonCell) AlternateTitle() string {
	rv := objc.Call[string](b_, objc.Sel("alternateTitle"))
	return rv
}

// The string displayed by the button when it’s in its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1535382-alternatetitle?language=objc
func (b_ ButtonCell) SetAlternateTitle(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setAlternateTitle:"), value)
}

// The position of the button’s image relative to its title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529593-imageposition?language=objc
func (b_ ButtonCell) ImagePosition() CellImagePosition {
	rv := objc.Call[CellImagePosition](b_, objc.Sel("imagePosition"))
	return rv
}

// The position of the button’s image relative to its title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529593-imageposition?language=objc
func (b_ ButtonCell) SetImagePosition(value CellImagePosition) {
	objc.Call[objc.Void](b_, objc.Sel("setImagePosition:"), value)
}

// The button’s key-equivalent character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529476-keyequivalent?language=objc
func (b_ ButtonCell) SetKeyEquivalent(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setKeyEquivalent:"), value)
}

// A Boolean value that indicates if the button displays its border only when the pointer is over it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527903-showsborderonlywhilemouseinside?language=objc
func (b_ ButtonCell) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Call[bool](b_, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}

// A Boolean value that indicates if the button displays its border only when the pointer is over it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527903-showsborderonlywhilemouseinside?language=objc
func (b_ ButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}

// The image the button displays in its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527064-alternateimage?language=objc
func (b_ ButtonCell) AlternateImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("alternateImage"))
	return rv
}

// The image the button displays in its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527064-alternateimage?language=objc
func (b_ ButtonCell) SetAlternateImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setAlternateImage:"), value)
}

// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1534152-imagedimswhendisabled?language=objc
func (b_ ButtonCell) ImageDimsWhenDisabled() bool {
	rv := objc.Call[bool](b_, objc.Sel("imageDimsWhenDisabled"))
	return rv
}

// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1534152-imagedimswhendisabled?language=objc
func (b_ ButtonCell) SetImageDimsWhenDisabled(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setImageDimsWhenDisabled:"), value)
}

// The appearance of the button’s border, if it has one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528696-bezelstyle?language=objc
func (b_ ButtonCell) BezelStyle() BezelStyle {
	rv := objc.Call[BezelStyle](b_, objc.Sel("bezelStyle"))
	return rv
}

// The appearance of the button’s border, if it has one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528696-bezelstyle?language=objc
func (b_ ButtonCell) SetBezelStyle(value BezelStyle) {
	objc.Call[objc.Void](b_, objc.Sel("setBezelStyle:"), value)
}

// The background color of the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529743-backgroundcolor?language=objc
func (b_ ButtonCell) BackgroundColor() Color {
	rv := objc.Call[Color](b_, objc.Sel("backgroundColor"))
	return rv
}

// The background color of the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529743-backgroundcolor?language=objc
func (b_ ButtonCell) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](b_, objc.Sel("setBackgroundColor:"), value)
}

// The sound that’s played when the user presses the button (that is during a mouse-down event). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1525955-sound?language=objc
func (b_ ButtonCell) Sound() Sound {
	rv := objc.Call[Sound](b_, objc.Sel("sound"))
	return rv
}

// The sound that’s played when the user presses the button (that is during a mouse-down event). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1525955-sound?language=objc
func (b_ ButtonCell) SetSound(value ISound) {
	objc.Call[objc.Void](b_, objc.Sel("setSound:"), value)
}

// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528459-highlightsby?language=objc
func (b_ ButtonCell) HighlightsBy() CellStyleMask {
	rv := objc.Call[CellStyleMask](b_, objc.Sel("highlightsBy"))
	return rv
}

// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528459-highlightsby?language=objc
func (b_ ButtonCell) SetHighlightsBy(value CellStyleMask) {
	objc.Call[objc.Void](b_, objc.Sel("setHighlightsBy:"), value)
}

// The scale factor for the button’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1535104-imagescaling?language=objc
func (b_ ButtonCell) ImageScaling() ImageScaling {
	rv := objc.Call[ImageScaling](b_, objc.Sel("imageScaling"))
	return rv
}

// The scale factor for the button’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1535104-imagescaling?language=objc
func (b_ ButtonCell) SetImageScaling(value ImageScaling) {
	objc.Call[objc.Void](b_, objc.Sel("setImageScaling:"), value)
}

// A Boolean value that indicates if the button is transparent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1530887-transparent?language=objc
func (b_ ButtonCell) IsTransparent() bool {
	rv := objc.Call[bool](b_, objc.Sel("isTransparent"))
	return rv
}

// A Boolean value that indicates if the button is transparent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1530887-transparent?language=objc
func (b_ ButtonCell) SetTransparent(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setTransparent:"), value)
}

// The title displayed by the button when it’s in its alternate state, as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1526922-attributedalternatetitle?language=objc
func (b_ ButtonCell) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](b_, objc.Sel("attributedAlternateTitle"))
	return rv
}

// The title displayed by the button when it’s in its alternate state, as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1526922-attributedalternatetitle?language=objc
func (b_ ButtonCell) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](b_, objc.Sel("setAttributedAlternateTitle:"), value)
}

// The title displayed by the button when it’s in its normal state as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529303-attributedtitle?language=objc
func (b_ ButtonCell) AttributedTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](b_, objc.Sel("attributedTitle"))
	return rv
}

// The title displayed by the button when it’s in its normal state as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529303-attributedtitle?language=objc
func (b_ ButtonCell) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](b_, objc.Sel("setAttributedTitle:"), value)
}

// The flags that indicate how the button cell shows its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1533225-showsstateby?language=objc
func (b_ ButtonCell) ShowsStateBy() CellStyleMask {
	rv := objc.Call[CellStyleMask](b_, objc.Sel("showsStateBy"))
	return rv
}

// The flags that indicate how the button cell shows its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1533225-showsstateby?language=objc
func (b_ ButtonCell) SetShowsStateBy(value CellStyleMask) {
	objc.Call[objc.Void](b_, objc.Sel("setShowsStateBy:"), value)
}