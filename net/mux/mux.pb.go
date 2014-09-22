// Code generated by protoc-gen-gogo.
// source: mux.proto
// DO NOT EDIT!

/*
Package mux is a generated protocol buffer package.

It is generated from these files:
	mux.proto

It has these top-level messages:
	PBProtocolMessage
*/
package mux

import proto "github.com/jbenet/go-ipfs/Godeps/_workspace/src/code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ProtocolID int32

const (
	ProtocolID_Test     ProtocolID = 0
	ProtocolID_Identify ProtocolID = 1
	ProtocolID_Routing  ProtocolID = 2
	ProtocolID_Exchange ProtocolID = 3
)

var ProtocolID_name = map[int32]string{
	0: "Test",
	1: "Identify",
	2: "Routing",
	3: "Exchange",
}
var ProtocolID_value = map[string]int32{
	"Test":     0,
	"Identify": 1,
	"Routing":  2,
	"Exchange": 3,
}

func (x ProtocolID) Enum() *ProtocolID {
	p := new(ProtocolID)
	*p = x
	return p
}
func (x ProtocolID) String() string {
	return proto.EnumName(ProtocolID_name, int32(x))
}
func (x *ProtocolID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProtocolID_value, data, "ProtocolID")
	if err != nil {
		return err
	}
	*x = ProtocolID(value)
	return nil
}

type PBProtocolMessage struct {
	ProtocolID       *ProtocolID `protobuf:"varint,1,req,enum=mux.ProtocolID" json:"ProtocolID,omitempty"`
	Data             []byte      `protobuf:"bytes,2,req" json:"Data,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PBProtocolMessage) Reset()         { *m = PBProtocolMessage{} }
func (m *PBProtocolMessage) String() string { return proto.CompactTextString(m) }
func (*PBProtocolMessage) ProtoMessage()    {}

func (m *PBProtocolMessage) GetProtocolID() ProtocolID {
	if m != nil && m.ProtocolID != nil {
		return *m.ProtocolID
	}
	return ProtocolID_Test
}

func (m *PBProtocolMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("mux.ProtocolID", ProtocolID_name, ProtocolID_value)
}
