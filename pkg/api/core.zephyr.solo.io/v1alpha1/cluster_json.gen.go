// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/core/v1alpha1/cluster.proto

package v1alpha1

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

// MarshalJSON is a custom marshaler for KubernetesClusterSpec
func (this *KubernetesClusterSpec) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubernetesClusterSpec
func (this *KubernetesClusterSpec) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ClusterMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ClusterUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
