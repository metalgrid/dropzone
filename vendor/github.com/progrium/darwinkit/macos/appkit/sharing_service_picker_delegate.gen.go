// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// An interface for managing content for the macOS share sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate?language=objc
type PSharingServicePickerDelegate interface {
	// optional
	SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []SharingService
	HasSharingServicePickerSharingServicesForItemsProposedSharingServices() bool

	// optional
	SharingServicePickerDidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService)
	HasSharingServicePickerDidChooseSharingService() bool

	// optional
	SharingServicePickerDelegateForSharingService(sharingServicePicker SharingServicePicker, sharingService SharingService) SharingServiceDelegateObject
	HasSharingServicePickerDelegateForSharingService() bool
}

// A delegate implementation builder for the [PSharingServicePickerDelegate] protocol.
type SharingServicePickerDelegate struct {
	_SharingServicePickerSharingServicesForItemsProposedSharingServices func(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []SharingService
	_SharingServicePickerDidChooseSharingService                        func(sharingServicePicker SharingServicePicker, service SharingService)
	_SharingServicePickerDelegateForSharingService                      func(sharingServicePicker SharingServicePicker, sharingService SharingService) SharingServiceDelegateObject
}

func (di *SharingServicePickerDelegate) HasSharingServicePickerSharingServicesForItemsProposedSharingServices() bool {
	return di._SharingServicePickerSharingServicesForItemsProposedSharingServices != nil
}

// Asks the delegate to specify which services to make available from the sharing service picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402664-sharingservicepicker?language=objc
func (di *SharingServicePickerDelegate) SetSharingServicePickerSharingServicesForItemsProposedSharingServices(f func(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []SharingService) {
	di._SharingServicePickerSharingServicesForItemsProposedSharingServices = f
}

// Asks the delegate to specify which services to make available from the sharing service picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402664-sharingservicepicker?language=objc
func (di *SharingServicePickerDelegate) SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []SharingService {
	return di._SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker, items, proposedServices)
}
func (di *SharingServicePickerDelegate) HasSharingServicePickerDidChooseSharingService() bool {
	return di._SharingServicePickerDidChooseSharingService != nil
}

// Tells the delegate that the person selected a sharing service for the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402610-sharingservicepicker?language=objc
func (di *SharingServicePickerDelegate) SetSharingServicePickerDidChooseSharingService(f func(sharingServicePicker SharingServicePicker, service SharingService)) {
	di._SharingServicePickerDidChooseSharingService = f
}

// Tells the delegate that the person selected a sharing service for the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402610-sharingservicepicker?language=objc
func (di *SharingServicePickerDelegate) SharingServicePickerDidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService) {
	di._SharingServicePickerDidChooseSharingService(sharingServicePicker, service)
}
func (di *SharingServicePickerDelegate) HasSharingServicePickerDelegateForSharingService() bool {
	return di._SharingServicePickerDelegateForSharingService != nil
}

// Asks your delegate to provide an object that the selected sharing service can use as its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402608-sharingservicepicker?language=objc
func (di *SharingServicePickerDelegate) SetSharingServicePickerDelegateForSharingService(f func(sharingServicePicker SharingServicePicker, sharingService SharingService) SharingServiceDelegateObject) {
	di._SharingServicePickerDelegateForSharingService = f
}

// Asks your delegate to provide an object that the selected sharing service can use as its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402608-sharingservicepicker?language=objc
func (di *SharingServicePickerDelegate) SharingServicePickerDelegateForSharingService(sharingServicePicker SharingServicePicker, sharingService SharingService) SharingServiceDelegateObject {
	return di._SharingServicePickerDelegateForSharingService(sharingServicePicker, sharingService)
}

// ensure impl type implements protocol interface
var _ PSharingServicePickerDelegate = (*SharingServicePickerDelegateObject)(nil)

// A concrete type for the [PSharingServicePickerDelegate] protocol.
type SharingServicePickerDelegateObject struct {
	objc.Object
}

func (s_ SharingServicePickerDelegateObject) HasSharingServicePickerSharingServicesForItemsProposedSharingServices() bool {
	return s_.RespondsToSelector(objc.Sel("sharingServicePicker:sharingServicesForItems:proposedSharingServices:"))
}

// Asks the delegate to specify which services to make available from the sharing service picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402664-sharingservicepicker?language=objc
func (s_ SharingServicePickerDelegateObject) SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []SharingService {
	rv := objc.Call[[]SharingService](s_, objc.Sel("sharingServicePicker:sharingServicesForItems:proposedSharingServices:"), sharingServicePicker, items, proposedServices)
	return rv
}

func (s_ SharingServicePickerDelegateObject) HasSharingServicePickerDidChooseSharingService() bool {
	return s_.RespondsToSelector(objc.Sel("sharingServicePicker:didChooseSharingService:"))
}

// Tells the delegate that the person selected a sharing service for the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402610-sharingservicepicker?language=objc
func (s_ SharingServicePickerDelegateObject) SharingServicePickerDidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService) {
	objc.Call[objc.Void](s_, objc.Sel("sharingServicePicker:didChooseSharingService:"), sharingServicePicker, service)
}

func (s_ SharingServicePickerDelegateObject) HasSharingServicePickerDelegateForSharingService() bool {
	return s_.RespondsToSelector(objc.Sel("sharingServicePicker:delegateForSharingService:"))
}

// Asks your delegate to provide an object that the selected sharing service can use as its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402608-sharingservicepicker?language=objc
func (s_ SharingServicePickerDelegateObject) SharingServicePickerDelegateForSharingService(sharingServicePicker SharingServicePicker, sharingService SharingService) SharingServiceDelegateObject {
	rv := objc.Call[SharingServiceDelegateObject](s_, objc.Sel("sharingServicePicker:delegateForSharingService:"), sharingServicePicker, sharingService)
	return rv
}