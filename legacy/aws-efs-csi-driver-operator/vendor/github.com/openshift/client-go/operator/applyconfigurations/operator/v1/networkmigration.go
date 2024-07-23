// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
)

// NetworkMigrationApplyConfiguration represents an declarative configuration of the NetworkMigration type for use
// with apply.
type NetworkMigrationApplyConfiguration struct {
	NetworkType *string                              `json:"networkType,omitempty"`
	MTU         *MTUMigrationApplyConfiguration      `json:"mtu,omitempty"`
	Features    *FeaturesMigrationApplyConfiguration `json:"features,omitempty"`
	Mode        *operatorv1.NetworkMigrationMode     `json:"mode,omitempty"`
}

// NetworkMigrationApplyConfiguration constructs an declarative configuration of the NetworkMigration type for use with
// apply.
func NetworkMigration() *NetworkMigrationApplyConfiguration {
	return &NetworkMigrationApplyConfiguration{}
}

// WithNetworkType sets the NetworkType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkType field is set to the value of the last call.
func (b *NetworkMigrationApplyConfiguration) WithNetworkType(value string) *NetworkMigrationApplyConfiguration {
	b.NetworkType = &value
	return b
}

// WithMTU sets the MTU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MTU field is set to the value of the last call.
func (b *NetworkMigrationApplyConfiguration) WithMTU(value *MTUMigrationApplyConfiguration) *NetworkMigrationApplyConfiguration {
	b.MTU = value
	return b
}

// WithFeatures sets the Features field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Features field is set to the value of the last call.
func (b *NetworkMigrationApplyConfiguration) WithFeatures(value *FeaturesMigrationApplyConfiguration) *NetworkMigrationApplyConfiguration {
	b.Features = value
	return b
}

// WithMode sets the Mode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mode field is set to the value of the last call.
func (b *NetworkMigrationApplyConfiguration) WithMode(value operatorv1.NetworkMigrationMode) *NetworkMigrationApplyConfiguration {
	b.Mode = &value
	return b
}