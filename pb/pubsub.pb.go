// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubsub.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Empty message is used for requests
// and responses that dont need any particular
// data.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// ListPeersRequest is used to return a list of
// peers that are subscribed to the given topic(s)
type ListPeersRequest struct {
	// the topics for which we should
	// list peers for
	Topics               []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPeersRequest) Reset()         { *m = ListPeersRequest{} }
func (m *ListPeersRequest) String() string { return proto.CompactTextString(m) }
func (*ListPeersRequest) ProtoMessage()    {}
func (*ListPeersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{1}
}

func (m *ListPeersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPeersRequest.Unmarshal(m, b)
}
func (m *ListPeersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPeersRequest.Marshal(b, m, deterministic)
}
func (m *ListPeersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPeersRequest.Merge(m, src)
}
func (m *ListPeersRequest) XXX_Size() int {
	return xxx_messageInfo_ListPeersRequest.Size(m)
}
func (m *ListPeersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPeersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPeersRequest proto.InternalMessageInfo

func (m *ListPeersRequest) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

// ListPeersResponse is a response to a ListPeersRequest
type ListPeersResponse struct {
	Peers                []*ListPeersResponse_Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListPeersResponse) Reset()         { *m = ListPeersResponse{} }
func (m *ListPeersResponse) String() string { return proto.CompactTextString(m) }
func (*ListPeersResponse) ProtoMessage()    {}
func (*ListPeersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{2}
}

func (m *ListPeersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPeersResponse.Unmarshal(m, b)
}
func (m *ListPeersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPeersResponse.Marshal(b, m, deterministic)
}
func (m *ListPeersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPeersResponse.Merge(m, src)
}
func (m *ListPeersResponse) XXX_Size() int {
	return xxx_messageInfo_ListPeersResponse.Size(m)
}
func (m *ListPeersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPeersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPeersResponse proto.InternalMessageInfo

func (m *ListPeersResponse) GetPeers() []*ListPeersResponse_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type ListPeersResponse_Peer struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	PeerID               string   `protobuf:"bytes,2,opt,name=peerID,proto3" json:"peerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPeersResponse_Peer) Reset()         { *m = ListPeersResponse_Peer{} }
func (m *ListPeersResponse_Peer) String() string { return proto.CompactTextString(m) }
func (*ListPeersResponse_Peer) ProtoMessage()    {}
func (*ListPeersResponse_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{2, 0}
}

func (m *ListPeersResponse_Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPeersResponse_Peer.Unmarshal(m, b)
}
func (m *ListPeersResponse_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPeersResponse_Peer.Marshal(b, m, deterministic)
}
func (m *ListPeersResponse_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPeersResponse_Peer.Merge(m, src)
}
func (m *ListPeersResponse_Peer) XXX_Size() int {
	return xxx_messageInfo_ListPeersResponse_Peer.Size(m)
}
func (m *ListPeersResponse_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPeersResponse_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_ListPeersResponse_Peer proto.InternalMessageInfo

func (m *ListPeersResponse_Peer) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ListPeersResponse_Peer) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

// PublishRequest is a message used to publish data to a topic
type PublishRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{3}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// SubscribeRequest is used to initiate a subscription
// to a given pubsub topic and stream received messages
type SubscribeRequest struct {
	// the topic we should subscribe to
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// indicates whether the server should
	// perform service discover for peers on the same topic
	Discover             bool     `protobuf:"varint,2,opt,name=discover,proto3" json:"discover,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{4}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SubscribeRequest) GetDiscover() bool {
	if m != nil {
		return m.Discover
	}
	return false
}

// Topics is a response that returns
// the names of all known topics
type TopicsResponse struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicsResponse) Reset()         { *m = TopicsResponse{} }
func (m *TopicsResponse) String() string { return proto.CompactTextString(m) }
func (*TopicsResponse) ProtoMessage()    {}
func (*TopicsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{5}
}

func (m *TopicsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicsResponse.Unmarshal(m, b)
}
func (m *TopicsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicsResponse.Marshal(b, m, deterministic)
}
func (m *TopicsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicsResponse.Merge(m, src)
}
func (m *TopicsResponse) XXX_Size() int {
	return xxx_messageInfo_TopicsResponse.Size(m)
}
func (m *TopicsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TopicsResponse proto.InternalMessageInfo

func (m *TopicsResponse) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// PubSubMessageResposne is a received pubsub message
// sent as a response to a subscription rpc call
type PubSubMessageResponse struct {
	From                 []byte   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Seqno                []byte   `protobuf:"bytes,3,opt,name=seqno,proto3" json:"seqno,omitempty"`
	TopicIDs             []string `protobuf:"bytes,4,rep,name=topicIDs,proto3" json:"topicIDs,omitempty"`
	Signature            []byte   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Key                  []byte   `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubSubMessageResponse) Reset()         { *m = PubSubMessageResponse{} }
func (m *PubSubMessageResponse) String() string { return proto.CompactTextString(m) }
func (*PubSubMessageResponse) ProtoMessage()    {}
func (*PubSubMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{6}
}

func (m *PubSubMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubMessageResponse.Unmarshal(m, b)
}
func (m *PubSubMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubMessageResponse.Marshal(b, m, deterministic)
}
func (m *PubSubMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubMessageResponse.Merge(m, src)
}
func (m *PubSubMessageResponse) XXX_Size() int {
	return xxx_messageInfo_PubSubMessageResponse.Size(m)
}
func (m *PubSubMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubMessageResponse proto.InternalMessageInfo

func (m *PubSubMessageResponse) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PubSubMessageResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PubSubMessageResponse) GetSeqno() []byte {
	if m != nil {
		return m.Seqno
	}
	return nil
}

func (m *PubSubMessageResponse) GetTopicIDs() []string {
	if m != nil {
		return m.TopicIDs
	}
	return nil
}

func (m *PubSubMessageResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PubSubMessageResponse) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*ListPeersRequest)(nil), "pb.ListPeersRequest")
	proto.RegisterType((*ListPeersResponse)(nil), "pb.ListPeersResponse")
	proto.RegisterType((*ListPeersResponse_Peer)(nil), "pb.ListPeersResponse.Peer")
	proto.RegisterType((*PublishRequest)(nil), "pb.PublishRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "pb.SubscribeRequest")
	proto.RegisterType((*TopicsResponse)(nil), "pb.TopicsResponse")
	proto.RegisterType((*PubSubMessageResponse)(nil), "pb.PubSubMessageResponse")
}

func init() { proto.RegisterFile("pubsub.proto", fileDescriptor_91df006b05e20cf7) }

var fileDescriptor_91df006b05e20cf7 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0x1b, 0x3b, 0xad, 0x87, 0x50, 0x85, 0x51, 0x8a, 0x8c, 0xc5, 0xa1, 0xda, 0x03, 0x8a,
	0x50, 0x65, 0x55, 0x85, 0x53, 0x4f, 0x1c, 0x82, 0x50, 0x25, 0x90, 0x22, 0x87, 0x17, 0xf0, 0xba,
	0x43, 0xb1, 0x20, 0xf6, 0xd6, 0xbb, 0x1b, 0x29, 0xe2, 0x69, 0x78, 0x3c, 0xde, 0x02, 0xed, 0xd8,
	0x71, 0x42, 0xb0, 0x7a, 0x9b, 0xef, 0xd3, 0x37, 0x33, 0xdf, 0xfc, 0xc0, 0x44, 0x59, 0xa9, 0xad,
	0x4c, 0x55, 0x53, 0x9b, 0x1a, 0x4f, 0x94, 0x14, 0xa7, 0x10, 0x7e, 0x5c, 0x2b, 0xb3, 0x15, 0x6f,
	0x61, 0xfa, 0xb9, 0xd4, 0x66, 0x49, 0xd4, 0xe8, 0x8c, 0x1e, 0x2d, 0x69, 0x83, 0x2f, 0x61, 0x6c,
	0x6a, 0x55, 0x16, 0x3a, 0xf6, 0x2f, 0x47, 0xf3, 0x28, 0xeb, 0x90, 0xf8, 0x05, 0x2f, 0x0e, 0xb4,
	0x5a, 0xd5, 0x95, 0x26, 0xbc, 0x86, 0x50, 0x39, 0x82, 0xb5, 0xcf, 0x6e, 0x92, 0x54, 0xc9, 0xf4,
	0x3f, 0x55, 0xea, 0x50, 0xd6, 0x0a, 0x93, 0xf7, 0x10, 0x38, 0x88, 0x33, 0x08, 0xb9, 0x70, 0xec,
	0x5f, 0xfa, 0xf3, 0x28, 0x6b, 0x81, 0x6b, 0xee, 0x64, 0x77, 0x8b, 0xf8, 0x84, 0xe9, 0x0e, 0x89,
	0x5b, 0x38, 0x5f, 0x5a, 0xf9, 0xb3, 0xd4, 0xdf, 0x77, 0x36, 0x87, 0xf3, 0x11, 0x82, 0xfb, 0xdc,
	0xe4, 0x9c, 0x3d, 0xc9, 0x38, 0x16, 0x0b, 0x98, 0xae, 0xac, 0xd4, 0x45, 0x53, 0x4a, 0x7a, 0x3a,
	0x3b, 0x81, 0xb3, 0xfb, 0x52, 0x17, 0xf5, 0x86, 0x1a, 0xae, 0x70, 0x96, 0xf5, 0x58, 0xbc, 0x81,
	0xf3, 0xaf, 0xbc, 0x88, 0x7e, 0xf6, 0x19, 0x84, 0x55, 0xbe, 0xa6, 0xdd, 0x9e, 0x5a, 0x20, 0x7e,
	0xfb, 0x70, 0xb1, 0xb4, 0x72, 0x65, 0xe5, 0x17, 0xd2, 0x3a, 0x7f, 0xa0, 0x5e, 0x8f, 0x10, 0x7c,
	0x6b, 0xea, 0x35, 0xb7, 0x9c, 0x64, 0x1c, 0x0f, 0xf9, 0x75, 0x75, 0x35, 0x3d, 0x56, 0x75, 0x3c,
	0x62, 0xb2, 0x05, 0xce, 0x1b, 0x9b, 0xbc, 0x5b, 0xe8, 0x38, 0xe0, 0x86, 0x3d, 0xc6, 0xd7, 0x10,
	0xe9, 0xf2, 0xa1, 0xca, 0x8d, 0x6d, 0x28, 0x0e, 0x39, 0x6b, 0x4f, 0xe0, 0x14, 0x46, 0x3f, 0x68,
	0x1b, 0x8f, 0x99, 0x77, 0xe1, 0xcd, 0x1f, 0x1f, 0x9e, 0xb7, 0x1e, 0x57, 0xd4, 0x6c, 0xca, 0x82,
	0xf0, 0x0a, 0xa2, 0x4f, 0x64, 0xda, 0x01, 0x31, 0x72, 0x57, 0xe4, 0x07, 0x49, 0xd0, 0x85, 0xff,
	0xce, 0x2d, 0x3c, 0xbc, 0x85, 0xa8, 0x3f, 0x32, 0xce, 0x8e, 0x6e, 0xce, 0x0b, 0x4e, 0x2e, 0x06,
	0x3f, 0x41, 0x78, 0xf8, 0x01, 0xa2, 0xfe, 0x1a, 0x6d, 0xee, 0xf1, 0x71, 0x92, 0x57, 0x8e, 0x1d,
	0xdc, 0xa1, 0xf0, 0xae, 0x7d, 0xbc, 0x82, 0xd3, 0xee, 0x17, 0x10, 0x3b, 0xe5, 0xc1, 0x63, 0x24,
	0x7b, 0xf7, 0xc2, 0x9b, 0xfb, 0x72, 0xcc, 0x6f, 0xff, 0xee, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x89, 0xa6, 0x22, 0x2b, 0x06, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubSubServiceClient is the client API for PubSubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubSubServiceClient interface {
	// GetTopics is used to return a list of all
	// known topics the pubsub instance is subscribed to.
	//
	// This is a unary rpc
	GetTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicsResponse, error)
	// ListPeers is used to return a list of peers subscribed
	// to a given topic or topics.
	//
	// This is a unary rpc
	ListPeers(ctx context.Context, in *ListPeersRequest, opts ...grpc.CallOption) (*ListPeersResponse, error)
	// Subscribe is used to subscribe to a topic and receive messages
	// Server will stream the messages received on the topic specified
	// during the initial subscription call, and send each message
	// back to the client as it is received.
	//
	// This is a server streaming rpc
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error)
	// Publish is used to send a stream of messages to a pubsub topic.
	//
	// This is a client streaming rpc
	Publish(ctx context.Context, opts ...grpc.CallOption) (PubSubService_PublishClient, error)
}

type pubSubServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubSubServiceClient(cc *grpc.ClientConn) PubSubServiceClient {
	return &pubSubServiceClient{cc}
}

func (c *pubSubServiceClient) GetTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicsResponse, error) {
	out := new(TopicsResponse)
	err := c.cc.Invoke(ctx, "/pb.PubSubService/GetTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubServiceClient) ListPeers(ctx context.Context, in *ListPeersRequest, opts ...grpc.CallOption) (*ListPeersResponse, error) {
	out := new(ListPeersResponse)
	err := c.cc.Invoke(ctx, "/pb.PubSubService/ListPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubSubService_serviceDesc.Streams[0], "/pb.PubSubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubSubService_SubscribeClient interface {
	Recv() (*PubSubMessageResponse, error)
	grpc.ClientStream
}

type pubSubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubSubServiceSubscribeClient) Recv() (*PubSubMessageResponse, error) {
	m := new(PubSubMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubSubServiceClient) Publish(ctx context.Context, opts ...grpc.CallOption) (PubSubService_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubSubService_serviceDesc.Streams[1], "/pb.PubSubService/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubServicePublishClient{stream}
	return x, nil
}

type PubSubService_PublishClient interface {
	Send(*PublishRequest) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type pubSubServicePublishClient struct {
	grpc.ClientStream
}

func (x *pubSubServicePublishClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pubSubServicePublishClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubSubServiceServer is the server API for PubSubService service.
type PubSubServiceServer interface {
	// GetTopics is used to return a list of all
	// known topics the pubsub instance is subscribed to.
	//
	// This is a unary rpc
	GetTopics(context.Context, *Empty) (*TopicsResponse, error)
	// ListPeers is used to return a list of peers subscribed
	// to a given topic or topics.
	//
	// This is a unary rpc
	ListPeers(context.Context, *ListPeersRequest) (*ListPeersResponse, error)
	// Subscribe is used to subscribe to a topic and receive messages
	// Server will stream the messages received on the topic specified
	// during the initial subscription call, and send each message
	// back to the client as it is received.
	//
	// This is a server streaming rpc
	Subscribe(*SubscribeRequest, PubSubService_SubscribeServer) error
	// Publish is used to send a stream of messages to a pubsub topic.
	//
	// This is a client streaming rpc
	Publish(PubSubService_PublishServer) error
}

func RegisterPubSubServiceServer(s *grpc.Server, srv PubSubServiceServer) {
	s.RegisterService(&_PubSubService_serviceDesc, srv)
}

func _PubSubService_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServiceServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PubSubService/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServiceServer).GetTopics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSubService_ListPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServiceServer).ListPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PubSubService/ListPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServiceServer).ListPeers(ctx, req.(*ListPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubSubServiceServer).Subscribe(m, &pubSubServiceSubscribeServer{stream})
}

type PubSubService_SubscribeServer interface {
	Send(*PubSubMessageResponse) error
	grpc.ServerStream
}

type pubSubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubSubServiceSubscribeServer) Send(m *PubSubMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PubSubService_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PubSubServiceServer).Publish(&pubSubServicePublishServer{stream})
}

type PubSubService_PublishServer interface {
	SendAndClose(*Empty) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type pubSubServicePublishServer struct {
	grpc.ServerStream
}

func (x *pubSubServicePublishServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pubSubServicePublishServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PubSubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PubSubService",
	HandlerType: (*PubSubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopics",
			Handler:    _PubSubService_GetTopics_Handler,
		},
		{
			MethodName: "ListPeers",
			Handler:    _PubSubService_ListPeers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubSubService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Publish",
			Handler:       _PubSubService_Publish_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pubsub.proto",
}
