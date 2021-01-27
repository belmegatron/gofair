// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderFilter order filter
//
// swagger:model OrderFilter
type OrderFilter struct {

	// Internal use only & should not be set on your filter (your subscription is already locked to your account). If set subscription will fail.
	AccountIds []int64 `json:"accountIds"`

	// Restricts to specified customerStrategyRefs; this will filter orders and StrategyMatchChanges accordingly (Note: overall postition is not filtered)
	CustomerStrategyRefs []string `json:"customerStrategyRefs"`

	// Returns overall / net position (See: OrderRunnerChange.mb / OrderRunnerChange.ml). Default=true
	IncludeOverallPosition bool `json:"includeOverallPosition,omitempty"`

	// Returns strategy positions (See: OrderRunnerChange.smc=Map<customerStrategyRef, StrategyMatchChange>) - these are sent in delta format as per overall position. Default=false
	PartitionMatchedByStrategyRef bool `json:"partitionMatchedByStrategyRef,omitempty"`
}

// Validate validates this order filter
func (m *OrderFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this order filter based on context it is used
func (m *OrderFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderFilter) UnmarshalBinary(b []byte) error {
	var res OrderFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}