// Code generated by DarwinKit. DO NOT EDIT.

package cloudkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [FetchDatabaseChangesOperation] class.
var FetchDatabaseChangesOperationClass = _FetchDatabaseChangesOperationClass{objc.GetClass("CKFetchDatabaseChangesOperation")}

type _FetchDatabaseChangesOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchDatabaseChangesOperation] class.
type IFetchDatabaseChangesOperation interface {
	IDatabaseOperation
	RecordZoneWithIDWasPurgedBlock() func(zoneID RecordZoneID)
	SetRecordZoneWithIDWasPurgedBlock(value func(zoneID RecordZoneID))
	RecordZoneWithIDWasDeletedBlock() func(zoneID RecordZoneID)
	SetRecordZoneWithIDWasDeletedBlock(value func(zoneID RecordZoneID))
	ChangeTokenUpdatedBlock() func(serverChangeToken ServerChangeToken)
	SetChangeTokenUpdatedBlock(value func(serverChangeToken ServerChangeToken))
	FetchAllChanges() bool
	SetFetchAllChanges(value bool)
	RecordZoneWithIDChangedBlock() func(zoneID RecordZoneID)
	SetRecordZoneWithIDChangedBlock(value func(zoneID RecordZoneID))
	FetchDatabaseChangesCompletionBlock() func(serverChangeToken ServerChangeToken, moreComing bool, operationError foundation.Error)
	SetFetchDatabaseChangesCompletionBlock(value func(serverChangeToken ServerChangeToken, moreComing bool, operationError foundation.Error))
	RecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock() func(zoneID RecordZoneID)
	SetRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock(value func(zoneID RecordZoneID))
	ResultsLimit() uint
	SetResultsLimit(value uint)
	PreviousServerChangeToken() ServerChangeToken
	SetPreviousServerChangeToken(value IServerChangeToken)
}

// An operation that fetches database changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation?language=objc
type FetchDatabaseChangesOperation struct {
	DatabaseOperation
}

func FetchDatabaseChangesOperationFrom(ptr unsafe.Pointer) FetchDatabaseChangesOperation {
	return FetchDatabaseChangesOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (f_ FetchDatabaseChangesOperation) Init() FetchDatabaseChangesOperation {
	rv := objc.Call[FetchDatabaseChangesOperation](f_, objc.Sel("init"))
	return rv
}

func (f_ FetchDatabaseChangesOperation) InitWithPreviousServerChangeToken(previousServerChangeToken IServerChangeToken) FetchDatabaseChangesOperation {
	rv := objc.Call[FetchDatabaseChangesOperation](f_, objc.Sel("initWithPreviousServerChangeToken:"), previousServerChangeToken)
	return rv
}

// Creates an operation for fetching database changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640502-initwithpreviousserverchangetoke?language=objc
func NewFetchDatabaseChangesOperationWithPreviousServerChangeToken(previousServerChangeToken IServerChangeToken) FetchDatabaseChangesOperation {
	instance := FetchDatabaseChangesOperationClass.Alloc().InitWithPreviousServerChangeToken(previousServerChangeToken)
	instance.Autorelease()
	return instance
}

func (fc _FetchDatabaseChangesOperationClass) Alloc() FetchDatabaseChangesOperation {
	rv := objc.Call[FetchDatabaseChangesOperation](fc, objc.Sel("alloc"))
	return rv
}

func (fc _FetchDatabaseChangesOperationClass) New() FetchDatabaseChangesOperation {
	rv := objc.Call[FetchDatabaseChangesOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchDatabaseChangesOperation() FetchDatabaseChangesOperation {
	return FetchDatabaseChangesOperationClass.New()
}

// The block to execute when CloudKit purges a record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/2866207-recordzonewithidwaspurgedblock?language=objc
func (f_ FetchDatabaseChangesOperation) RecordZoneWithIDWasPurgedBlock() func(zoneID RecordZoneID) {
	rv := objc.Call[func(zoneID RecordZoneID)](f_, objc.Sel("recordZoneWithIDWasPurgedBlock"))
	return rv
}

// The block to execute when CloudKit purges a record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/2866207-recordzonewithidwaspurgedblock?language=objc
func (f_ FetchDatabaseChangesOperation) SetRecordZoneWithIDWasPurgedBlock(value func(zoneID RecordZoneID)) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordZoneWithIDWasPurgedBlock:"), value)
}

// The block to execute when a record zone no longer exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640428-recordzonewithidwasdeletedblock?language=objc
func (f_ FetchDatabaseChangesOperation) RecordZoneWithIDWasDeletedBlock() func(zoneID RecordZoneID) {
	rv := objc.Call[func(zoneID RecordZoneID)](f_, objc.Sel("recordZoneWithIDWasDeletedBlock"))
	return rv
}

// The block to execute when a record zone no longer exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640428-recordzonewithidwasdeletedblock?language=objc
func (f_ FetchDatabaseChangesOperation) SetRecordZoneWithIDWasDeletedBlock(value func(zoneID RecordZoneID)) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordZoneWithIDWasDeletedBlock:"), value)
}

// The block to execute when the change token updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640467-changetokenupdatedblock?language=objc
func (f_ FetchDatabaseChangesOperation) ChangeTokenUpdatedBlock() func(serverChangeToken ServerChangeToken) {
	rv := objc.Call[func(serverChangeToken ServerChangeToken)](f_, objc.Sel("changeTokenUpdatedBlock"))
	return rv
}

// The block to execute when the change token updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640467-changetokenupdatedblock?language=objc
func (f_ FetchDatabaseChangesOperation) SetChangeTokenUpdatedBlock(value func(serverChangeToken ServerChangeToken)) {
	objc.Call[objc.Void](f_, objc.Sel("setChangeTokenUpdatedBlock:"), value)
}

// A Boolean value that indicates whether to send repeated requests to the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640473-fetchallchanges?language=objc
func (f_ FetchDatabaseChangesOperation) FetchAllChanges() bool {
	rv := objc.Call[bool](f_, objc.Sel("fetchAllChanges"))
	return rv
}

// A Boolean value that indicates whether to send repeated requests to the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640473-fetchallchanges?language=objc
func (f_ FetchDatabaseChangesOperation) SetFetchAllChanges(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchAllChanges:"), value)
}

// The block to execute with a single record zone change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640391-recordzonewithidchangedblock?language=objc
func (f_ FetchDatabaseChangesOperation) RecordZoneWithIDChangedBlock() func(zoneID RecordZoneID) {
	rv := objc.Call[func(zoneID RecordZoneID)](f_, objc.Sel("recordZoneWithIDChangedBlock"))
	return rv
}

// The block to execute with a single record zone change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640391-recordzonewithidchangedblock?language=objc
func (f_ FetchDatabaseChangesOperation) SetRecordZoneWithIDChangedBlock(value func(zoneID RecordZoneID)) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordZoneWithIDChangedBlock:"), value)
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640434-fetchdatabasechangescompletionbl?language=objc
func (f_ FetchDatabaseChangesOperation) FetchDatabaseChangesCompletionBlock() func(serverChangeToken ServerChangeToken, moreComing bool, operationError foundation.Error) {
	rv := objc.Call[func(serverChangeToken ServerChangeToken, moreComing bool, operationError foundation.Error)](f_, objc.Sel("fetchDatabaseChangesCompletionBlock"))
	return rv
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640434-fetchdatabasechangescompletionbl?language=objc
func (f_ FetchDatabaseChangesOperation) SetFetchDatabaseChangesCompletionBlock(value func(serverChangeToken ServerChangeToken, moreComing bool, operationError foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchDatabaseChangesCompletionBlock:"), value)
}

// The block to execute when a user-invoked account reset deletes a record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/3746820-recordzonewithidwasdeletedduetou?language=objc
func (f_ FetchDatabaseChangesOperation) RecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock() func(zoneID RecordZoneID) {
	rv := objc.Call[func(zoneID RecordZoneID)](f_, objc.Sel("recordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock"))
	return rv
}

// The block to execute when a user-invoked account reset deletes a record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/3746820-recordzonewithidwasdeletedduetou?language=objc
func (f_ FetchDatabaseChangesOperation) SetRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock(value func(zoneID RecordZoneID)) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock:"), value)
}

// The maximum number of results that the operation fetches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640520-resultslimit?language=objc
func (f_ FetchDatabaseChangesOperation) ResultsLimit() uint {
	rv := objc.Call[uint](f_, objc.Sel("resultsLimit"))
	return rv
}

// The maximum number of results that the operation fetches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640520-resultslimit?language=objc
func (f_ FetchDatabaseChangesOperation) SetResultsLimit(value uint) {
	objc.Call[objc.Void](f_, objc.Sel("setResultsLimit:"), value)
}

// The server change token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640522-previousserverchangetoken?language=objc
func (f_ FetchDatabaseChangesOperation) PreviousServerChangeToken() ServerChangeToken {
	rv := objc.Call[ServerChangeToken](f_, objc.Sel("previousServerChangeToken"))
	return rv
}

// The server change token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/1640522-previousserverchangetoken?language=objc
func (f_ FetchDatabaseChangesOperation) SetPreviousServerChangeToken(value IServerChangeToken) {
	objc.Call[objc.Void](f_, objc.Sel("setPreviousServerChangeToken:"), value)
}