// Code generated by protoc-gen-go.
// source: orderer/configuration.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/hyperledger/fabric/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=maxMessageCount" json:"maxMessageCount,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout string `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *BatchTimeout) Reset()                    { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string            { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()               {}
func (*BatchTimeout) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// When submitting a new chain configuration transaction to create a new chain,
// the first configuration item must be of type Orderer with Key CreationPolicy
// and contents of a Marshaled CreationPolicy. The policy should be set to the
// policy which was supplied by the ordering service for the client's chain
// creation. The digest should be the hash of the concatenation of the remaining
// ConfigurationItem bytes. The signatures of the configuration item should
// satisfy the policy for chain creation.
type CreationPolicy struct {
	// The name of the policy which should be used to validate the creation of
	// this chain
	Policy string `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
	// The hash of the concatenation of remaining configuration item bytes
	Digest []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *CreationPolicy) Reset()                    { *m = CreationPolicy{} }
func (m *CreationPolicy) String() string            { return proto.CompactTextString(m) }
func (*CreationPolicy) ProtoMessage()               {}
func (*CreationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type ChainCreators struct {
	// A list of policies, any of which may be specified as the chain creation
	// policy in a chain creation request
	Policies []string `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *ChainCreators) Reset()                    { *m = ChainCreators{} }
func (m *ChainCreators) String() string            { return proto.CompactTextString(m) }
func (*ChainCreators) ProtoMessage()               {}
func (*ChainCreators) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *KafkaBrokers) Reset()                    { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string            { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()               {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*CreationPolicy)(nil), "orderer.CreationPolicy")
	proto.RegisterType((*ChainCreators)(nil), "orderer.ChainCreators")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4b, 0xf3, 0x40,
	0x10, 0xc5, 0xc9, 0xf7, 0x49, 0x6b, 0x97, 0x56, 0x61, 0x05, 0x09, 0xf5, 0x52, 0xe2, 0x25, 0xa0,
	0x74, 0x0f, 0xe2, 0x5d, 0x92, 0xa3, 0x08, 0x12, 0x7b, 0xf2, 0xb6, 0xd9, 0x4c, 0x92, 0xa5, 0xcd,
	0x4e, 0x98, 0xdd, 0x80, 0xf1, 0xaf, 0x97, 0x6c, 0xb6, 0x1e, 0x3c, 0xcd, 0xfb, 0xcd, 0xbc, 0x19,
	0x1e, 0xc3, 0xee, 0x90, 0x2a, 0x20, 0x20, 0xa1, 0xd0, 0xd4, 0xba, 0x19, 0x48, 0x3a, 0x8d, 0x66,
	0xdf, 0x13, 0x3a, 0xe4, 0xcb, 0x30, 0xdc, 0xde, 0x28, 0xec, 0x3a, 0x34, 0x62, 0x2e, 0xf3, 0x34,
	0xb9, 0x67, 0x9b, 0x1c, 0x8d, 0x05, 0x63, 0x07, 0x7b, 0x18, 0x7b, 0xe0, 0x9c, 0x5d, 0xb8, 0xb1,
	0x87, 0x38, 0xda, 0x45, 0xe9, 0xaa, 0xf0, 0x3a, 0x79, 0x66, 0xab, 0x4c, 0x3a, 0xd5, 0x7e, 0xe8,
	0x6f, 0xe0, 0x29, 0xbb, 0xee, 0xe4, 0xd7, 0x1b, 0x58, 0x2b, 0x1b, 0xc8, 0x71, 0x30, 0xce, 0x7b,
	0x37, 0xc5, 0xdf, 0x76, 0x92, 0xb2, 0xb5, 0x5f, 0x3b, 0xe8, 0x0e, 0x70, 0x70, 0x3c, 0x66, 0x4b,
	0x37, 0xcb, 0x70, 0xfd, 0x8c, 0xc9, 0x0b, 0xbb, 0xca, 0x09, 0x7c, 0xea, 0x77, 0x3c, 0x69, 0x35,
	0xf2, 0x5b, 0xb6, 0xe8, 0xbd, 0x0a, 0xd6, 0x40, 0x53, 0xbf, 0xd2, 0x0d, 0x58, 0x17, 0xff, 0xdb,
	0x45, 0xe9, 0xba, 0x08, 0x94, 0x3c, 0xb0, 0x4d, 0xde, 0x4a, 0x6d, 0xfc, 0x19, 0x24, 0xcb, 0xb7,
	0xec, 0xd2, 0xaf, 0x68, 0xb0, 0x71, 0xb4, 0xfb, 0x9f, 0xae, 0x8a, 0x5f, 0x9e, 0x82, 0xbd, 0xca,
	0xfa, 0x28, 0x33, 0xc2, 0x23, 0x90, 0x9d, 0x82, 0x95, 0xb3, 0x0c, 0xd6, 0x33, 0x66, 0xfb, 0xcf,
	0xc7, 0x46, 0xbb, 0x76, 0x28, 0xf7, 0x0a, 0x3b, 0xd1, 0x8e, 0x3d, 0xd0, 0x09, 0xaa, 0x06, 0x48,
	0xd4, 0xb2, 0x24, 0xad, 0x84, 0xff, 0xa2, 0x15, 0xe1, 0xc7, 0xe5, 0xc2, 0xf3, 0xd3, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x71, 0x0c, 0x1d, 0x71, 0x92, 0x01, 0x00, 0x00,
}
