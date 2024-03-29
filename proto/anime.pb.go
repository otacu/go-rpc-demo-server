// Code generated by protoc-gen-go. DO NOT EDIT.
// source: anime.proto

package com_egoist_go_demo_anime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type QueryAnimeInfoRequest struct {
	AnimeId              int32    `protobuf:"varint,1,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryAnimeInfoRequest) Reset()         { *m = QueryAnimeInfoRequest{} }
func (m *QueryAnimeInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAnimeInfoRequest) ProtoMessage()    {}
func (*QueryAnimeInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cdf86e3cafc4bf7, []int{0}
}

func (m *QueryAnimeInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryAnimeInfoRequest.Unmarshal(m, b)
}
func (m *QueryAnimeInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryAnimeInfoRequest.Marshal(b, m, deterministic)
}
func (m *QueryAnimeInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAnimeInfoRequest.Merge(m, src)
}
func (m *QueryAnimeInfoRequest) XXX_Size() int {
	return xxx_messageInfo_QueryAnimeInfoRequest.Size(m)
}
func (m *QueryAnimeInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAnimeInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAnimeInfoRequest proto.InternalMessageInfo

func (m *QueryAnimeInfoRequest) GetAnimeId() int32 {
	if m != nil {
		return m.AnimeId
	}
	return 0
}

type AnimeThemeSong struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Singer               string   `protobuf:"bytes,2,opt,name=singer,proto3" json:"singer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnimeThemeSong) Reset()         { *m = AnimeThemeSong{} }
func (m *AnimeThemeSong) String() string { return proto.CompactTextString(m) }
func (*AnimeThemeSong) ProtoMessage()    {}
func (*AnimeThemeSong) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cdf86e3cafc4bf7, []int{1}
}

func (m *AnimeThemeSong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnimeThemeSong.Unmarshal(m, b)
}
func (m *AnimeThemeSong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnimeThemeSong.Marshal(b, m, deterministic)
}
func (m *AnimeThemeSong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnimeThemeSong.Merge(m, src)
}
func (m *AnimeThemeSong) XXX_Size() int {
	return xxx_messageInfo_AnimeThemeSong.Size(m)
}
func (m *AnimeThemeSong) XXX_DiscardUnknown() {
	xxx_messageInfo_AnimeThemeSong.DiscardUnknown(m)
}

var xxx_messageInfo_AnimeThemeSong proto.InternalMessageInfo

func (m *AnimeThemeSong) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnimeThemeSong) GetSinger() string {
	if m != nil {
		return m.Singer
	}
	return ""
}

type QueryAnimeInfoResponse struct {
	AnimeId              int32             `protobuf:"varint,1,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	EnName               string            `protobuf:"bytes,2,opt,name=en_name,json=enName,proto3" json:"en_name,omitempty"`
	ThemeSongArray       []*AnimeThemeSong `protobuf:"bytes,3,rep,name=themeSongArray,proto3" json:"themeSongArray,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *QueryAnimeInfoResponse) Reset()         { *m = QueryAnimeInfoResponse{} }
func (m *QueryAnimeInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAnimeInfoResponse) ProtoMessage()    {}
func (*QueryAnimeInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cdf86e3cafc4bf7, []int{2}
}

func (m *QueryAnimeInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryAnimeInfoResponse.Unmarshal(m, b)
}
func (m *QueryAnimeInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryAnimeInfoResponse.Marshal(b, m, deterministic)
}
func (m *QueryAnimeInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAnimeInfoResponse.Merge(m, src)
}
func (m *QueryAnimeInfoResponse) XXX_Size() int {
	return xxx_messageInfo_QueryAnimeInfoResponse.Size(m)
}
func (m *QueryAnimeInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAnimeInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAnimeInfoResponse proto.InternalMessageInfo

func (m *QueryAnimeInfoResponse) GetAnimeId() int32 {
	if m != nil {
		return m.AnimeId
	}
	return 0
}

func (m *QueryAnimeInfoResponse) GetEnName() string {
	if m != nil {
		return m.EnName
	}
	return ""
}

func (m *QueryAnimeInfoResponse) GetThemeSongArray() []*AnimeThemeSong {
	if m != nil {
		return m.ThemeSongArray
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAnimeInfoRequest)(nil), "com.egoist.go.demo.anime.QueryAnimeInfoRequest")
	proto.RegisterType((*AnimeThemeSong)(nil), "com.egoist.go.demo.anime.AnimeThemeSong")
	proto.RegisterType((*QueryAnimeInfoResponse)(nil), "com.egoist.go.demo.anime.QueryAnimeInfoResponse")
}

func init() { proto.RegisterFile("anime.proto", fileDescriptor_5cdf86e3cafc4bf7) }

var fileDescriptor_5cdf86e3cafc4bf7 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x09, 0x85, 0x16, 0xae, 0x28, 0xc3, 0x49, 0x94, 0xc0, 0x54, 0x65, 0xca, 0x64, 0x50,
	0x58, 0x59, 0x3a, 0x76, 0x41, 0x90, 0xb2, 0x57, 0xa1, 0x39, 0x82, 0x07, 0xdf, 0x15, 0xdb, 0x41,
	0xea, 0x03, 0xf0, 0x1a, 0x3c, 0x2b, 0xea, 0x51, 0x86, 0x56, 0x14, 0xb1, 0xf9, 0xb7, 0xff, 0xbb,
	0xef, 0xf7, 0x1d, 0x0c, 0x6b, 0xb6, 0x8e, 0xcc, 0xd2, 0x4b, 0x14, 0xcc, 0x16, 0xe2, 0x0c, 0xb5,
	0x62, 0x43, 0x34, 0xad, 0x98, 0x86, 0x9c, 0x18, 0x7d, 0xcf, 0x4b, 0x38, 0x7f, 0xec, 0xc8, 0xaf,
	0x26, 0x6b, 0x35, 0xe5, 0x17, 0xa9, 0xe8, 0xad, 0xa3, 0x10, 0xf1, 0x12, 0x4e, 0xd4, 0x31, 0xb7,
	0x4d, 0x96, 0x8c, 0x93, 0xe2, 0xb8, 0x1a, 0xa8, 0x9e, 0x36, 0xf9, 0x1d, 0xa4, 0x6a, 0x7f, 0x7a,
	0x25, 0x47, 0x33, 0xe1, 0x16, 0x11, 0x8e, 0xb8, 0x76, 0xa4, 0xc6, 0xd3, 0x4a, 0xcf, 0x38, 0x82,
	0x7e, 0xb0, 0xdc, 0x92, 0xcf, 0x0e, 0xf5, 0x76, 0xa3, 0xf2, 0xcf, 0x04, 0x46, 0xbb, 0xc8, 0xb0,
	0x14, 0x0e, 0xf4, 0x07, 0x13, 0x2f, 0x60, 0x40, 0x3c, 0x57, 0xc8, 0xa6, 0x1d, 0xf1, 0xfd, 0x1a,
	0xf3, 0x00, 0x69, 0xfc, 0xc9, 0x31, 0xf1, 0xbe, 0x5e, 0x65, 0xbd, 0x71, 0xaf, 0x18, 0x96, 0x85,
	0xd9, 0xf7, 0x67, 0xb3, 0x1d, 0xbe, 0xda, 0xa9, 0x2f, 0x3f, 0x12, 0x38, 0x53, 0xcb, 0x8c, 0xfc,
	0xbb, 0x5d, 0x10, 0x76, 0x90, 0x6e, 0x07, 0xc6, 0xeb, 0xfd, 0xcd, 0x7f, 0x9d, 0xe6, 0xd5, 0xcd,
	0xff, 0x0b, 0xbe, 0x67, 0x91, 0x1f, 0x3c, 0xf7, 0x75, 0x77, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x23, 0x76, 0x3a, 0x21, 0xca, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AnimeService service

type AnimeServiceClient interface {
	QueryAnimeInfo(ctx context.Context, in *QueryAnimeInfoRequest, opts ...client.CallOption) (*QueryAnimeInfoResponse, error)
}

type animeServiceClient struct {
	c           client.Client
	serviceName string
}

func NewAnimeServiceClient(serviceName string, c client.Client) AnimeServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.egoist.go.demo.anime"
	}
	return &animeServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *animeServiceClient) QueryAnimeInfo(ctx context.Context, in *QueryAnimeInfoRequest, opts ...client.CallOption) (*QueryAnimeInfoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AnimeService.QueryAnimeInfo", in)
	out := new(QueryAnimeInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AnimeService service

type AnimeServiceHandler interface {
	QueryAnimeInfo(context.Context, *QueryAnimeInfoRequest, *QueryAnimeInfoResponse) error
}

func RegisterAnimeServiceHandler(s server.Server, hdlr AnimeServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AnimeService{hdlr}, opts...))
}

type AnimeService struct {
	AnimeServiceHandler
}

func (h *AnimeService) QueryAnimeInfo(ctx context.Context, in *QueryAnimeInfoRequest, out *QueryAnimeInfoResponse) error {
	return h.AnimeServiceHandler.QueryAnimeInfo(ctx, in, out)
}
