// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type TargetStore string

// Enum values for TargetStore
const (
	TargetStoreOnlineStore  TargetStore = "OnlineStore"
	TargetStoreOfflineStore TargetStore = "OfflineStore"
)

// Values returns all known values for TargetStore. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetStore) Values() []TargetStore {
	return []TargetStore{
		"OnlineStore",
		"OfflineStore",
	}
}