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

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// ContainerRestartRuleOnExitCodesApplyConfiguration represents a declarative configuration of the ContainerRestartRuleOnExitCodes type for use
// with apply.
type ContainerRestartRuleOnExitCodesApplyConfiguration struct {
	Operator *corev1.ContainerRestartRuleOnExitCodesOperator `json:"operator,omitempty"`
	Values   []int32                                         `json:"values,omitempty"`
}

// ContainerRestartRuleOnExitCodesApplyConfiguration constructs a declarative configuration of the ContainerRestartRuleOnExitCodes type for use with
// apply.
func ContainerRestartRuleOnExitCodes() *ContainerRestartRuleOnExitCodesApplyConfiguration {
	return &ContainerRestartRuleOnExitCodesApplyConfiguration{}
}

// WithOperator sets the Operator field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Operator field is set to the value of the last call.
func (b *ContainerRestartRuleOnExitCodesApplyConfiguration) WithOperator(value corev1.ContainerRestartRuleOnExitCodesOperator) *ContainerRestartRuleOnExitCodesApplyConfiguration {
	b.Operator = &value
	return b
}

// WithValues adds the given value to the Values field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Values field.
func (b *ContainerRestartRuleOnExitCodesApplyConfiguration) WithValues(values ...int32) *ContainerRestartRuleOnExitCodesApplyConfiguration {
	for i := range values {
		b.Values = append(b.Values, values[i])
	}
	return b
}
