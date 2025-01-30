/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// FunctionResourcesApplyConfiguration represents an declarative configuration of the FunctionResources type for use
// with apply.
type FunctionResourcesApplyConfiguration struct {
	Memory *string `json:"memory,omitempty"`
	CPU    *string `json:"cpu,omitempty"`
}

// FunctionResourcesApplyConfiguration constructs an declarative configuration of the FunctionResources type for use with
// apply.
func FunctionResources() *FunctionResourcesApplyConfiguration {
	return &FunctionResourcesApplyConfiguration{}
}

// WithMemory sets the Memory field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Memory field is set to the value of the last call.
func (b *FunctionResourcesApplyConfiguration) WithMemory(value string) *FunctionResourcesApplyConfiguration {
	b.Memory = &value
	return b
}

// WithCPU sets the CPU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CPU field is set to the value of the last call.
func (b *FunctionResourcesApplyConfiguration) WithCPU(value string) *FunctionResourcesApplyConfiguration {
	b.CPU = &value
	return b
}
