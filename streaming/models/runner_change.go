// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunnerChange runner change
//
// swagger:model RunnerChange
type RunnerChange struct {

	// Available To Back - PriceVol tuple delta of price changes (0 vol is remove)
	Atb [][]float64 `json:"atb"`

	// Available To Lay - PriceVol tuple delta of price changes (0 vol is remove)
	Atl [][]float64 `json:"atl"`

	// Best Available To Back - LevelPriceVol triple delta of price changes, keyed by level (0 vol is remove)
	Batb [][]float64 `json:"batb"`

	// Best Available To Lay - LevelPriceVol triple delta of price changes, keyed by level (0 vol is remove)
	Batl [][]float64 `json:"batl"`

	// Best Display Available To Back (includes virtual prices)- LevelPriceVol triple delta of price changes, keyed by level (0 vol is remove)
	Bdatb [][]float64 `json:"bdatb"`

	// Best Display Available To Lay (includes virtual prices)- LevelPriceVol triple delta of price changes, keyed by level (0 vol is remove)
	Bdatl [][]float64 `json:"bdatl"`

	// Handicap - the handicap of the runner (selection) (null if not applicable)
	Hc float64 `json:"hc,omitempty"`

	// Selection Id - the id of the runner (selection)
	ID int64 `json:"id,omitempty"`

	// Last Traded Price - The last traded price (or null if un-changed)
	Ltp float64 `json:"ltp,omitempty"`

	// Starting Price Back - PriceVol tuple delta of price changes (0 vol is remove)
	Spb [][]float64 `json:"spb"`

	// Starting Price Far - The far starting price (or null if un-changed)
	Spf float64 `json:"spf,omitempty"`

	// Starting Price Lay - PriceVol tuple delta of price changes (0 vol is remove)
	Spl [][]float64 `json:"spl"`

	// Starting Price Near - The far starting price (or null if un-changed)
	Spn float64 `json:"spn,omitempty"`

	// Traded - PriceVol tuple delta of price changes (0 vol is remove)
	Trd [][]float64 `json:"trd"`

	// The total amount matched. This value is truncated at 2dp.
	Tv float64 `json:"tv,omitempty"`
}

// Validate validates this runner change
func (m *RunnerChange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this runner change based on context it is used
func (m *RunnerChange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunnerChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunnerChange) UnmarshalBinary(b []byte) error {
	var res RunnerChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
