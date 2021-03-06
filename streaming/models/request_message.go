// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// RequestMessage request message
//
// swagger:discriminator RequestMessage op
type RequestMessage interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Client generated unique id to link request with response (like json rpc)
	ID() int32
	SetID(int32)

	// The operation type
	Op() string
	SetOp(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type requestMessage struct {
	idField int32

	opField string
}

// ID gets the id of this polymorphic type
func (m *requestMessage) ID() int32 {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *requestMessage) SetID(val int32) {
	m.idField = val
}

// Op gets the op of this polymorphic type
func (m *requestMessage) Op() string {
	return "RequestMessage"
}

// SetOp sets the op of this polymorphic type
func (m *requestMessage) SetOp(val string) {
}

// UnmarshalRequestMessageSlice unmarshals polymorphic slices of RequestMessage
func UnmarshalRequestMessageSlice(reader io.Reader, consumer runtime.Consumer) ([]RequestMessage, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []RequestMessage
	for _, element := range elements {
		obj, err := unmarshalRequestMessage(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalRequestMessage unmarshals polymorphic RequestMessage
func UnmarshalRequestMessage(reader io.Reader, consumer runtime.Consumer) (RequestMessage, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalRequestMessage(data, consumer)
}

func unmarshalRequestMessage(data []byte, consumer runtime.Consumer) (RequestMessage, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the op property.
	var getType struct {
		Op string `json:"op"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("op", "body", getType.Op); err != nil {
		return nil, err
	}

	// The value of op is used to determine which type to create and unmarshal the data into
	switch getType.Op {
	case "AuthenticationMessage":
		var result AuthenticationMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "HeartbeatMessage":
		var result HeartbeatMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "MarketSubscriptionMessage":
		var result MarketSubscriptionMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "OrderSubscriptionMessage":
		var result OrderSubscriptionMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "RequestMessage":
		var result requestMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid op value: %q", getType.Op)
}

// Validate validates this request message
func (m *requestMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this request message based on context it is used
func (m *requestMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
