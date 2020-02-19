// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/core/v1alpha1/selector.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ClusterSelector
func (this *ClusterSelector) MarshalJSON() ([]byte, error) {
	str, err := SelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ClusterSelector
func (this *ClusterSelector) UnmarshalJSON(b []byte) error {
	return SelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Selector
func (this *Selector) MarshalJSON() ([]byte, error) {
	str, err := SelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Selector
func (this *Selector) UnmarshalJSON(b []byte) error {
	return SelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	SelectorMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	SelectorUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)