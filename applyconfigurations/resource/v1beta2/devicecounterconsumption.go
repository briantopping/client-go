/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta2

// DeviceCounterConsumptionApplyConfiguration represents a declarative configuration of the DeviceCounterConsumption type for use
// with apply.
type DeviceCounterConsumptionApplyConfiguration struct {
	CounterSet *string                              `json:"counterSet,omitempty"`
	Counters   map[string]CounterApplyConfiguration `json:"counters,omitempty"`
}

// DeviceCounterConsumptionApplyConfiguration constructs a declarative configuration of the DeviceCounterConsumption type for use with
// apply.
func DeviceCounterConsumption() *DeviceCounterConsumptionApplyConfiguration {
	return &DeviceCounterConsumptionApplyConfiguration{}
}

// WithCounterSet sets the CounterSet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CounterSet field is set to the value of the last call.
func (b *DeviceCounterConsumptionApplyConfiguration) WithCounterSet(value string) *DeviceCounterConsumptionApplyConfiguration {
	b.CounterSet = &value
	return b
}

// WithCounters puts the entries into the Counters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Counters field,
// overwriting an existing map entries in Counters field with the same key.
func (b *DeviceCounterConsumptionApplyConfiguration) WithCounters(entries map[string]CounterApplyConfiguration) *DeviceCounterConsumptionApplyConfiguration {
	if b.Counters == nil && len(entries) > 0 {
		b.Counters = make(map[string]CounterApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Counters[k] = v
	}
	return b
}
