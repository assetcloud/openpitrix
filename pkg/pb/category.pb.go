// Code generated by protoc-gen-go. DO NOT EDIT.
// source: category.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Category struct {
	CategoryId           *wrappers.StringValue `protobuf:"bytes,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Locale               *wrappers.StringValue `protobuf:"bytes,3,opt,name=locale" json:"locale,omitempty"`
	Owner                *wrappers.StringValue `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
	CreateTime           *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_01b76ca5827ebaa0, []int{0}
}
func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (dst *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(dst, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetCategoryId() *wrappers.StringValue {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

func (m *Category) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Category) GetLocale() *wrappers.StringValue {
	if m != nil {
		return m.Locale
	}
	return nil
}

func (m *Category) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Category) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Category) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type DescribeCategoriesRequest struct {
	SearchWord           *wrappers.StringValue `protobuf:"bytes,1,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Limit                uint32                `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Offset               uint32                `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	CategoryId           []string              `protobuf:"bytes,4,rep,name=category_id,json=categoryId" json:"category_id,omitempty"`
	Name                 []string              `protobuf:"bytes,5,rep,name=name" json:"name,omitempty"`
	Owner                []string              `protobuf:"bytes,6,rep,name=owner" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DescribeCategoriesRequest) Reset()         { *m = DescribeCategoriesRequest{} }
func (m *DescribeCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeCategoriesRequest) ProtoMessage()    {}
func (*DescribeCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_01b76ca5827ebaa0, []int{1}
}
func (m *DescribeCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeCategoriesRequest.Unmarshal(m, b)
}
func (m *DescribeCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeCategoriesRequest.Marshal(b, m, deterministic)
}
func (dst *DescribeCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeCategoriesRequest.Merge(dst, src)
}
func (m *DescribeCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeCategoriesRequest.Size(m)
}
func (m *DescribeCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeCategoriesRequest proto.InternalMessageInfo

func (m *DescribeCategoriesRequest) GetSearchWord() *wrappers.StringValue {
	if m != nil {
		return m.SearchWord
	}
	return nil
}

func (m *DescribeCategoriesRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeCategoriesRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeCategoriesRequest) GetCategoryId() []string {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

func (m *DescribeCategoriesRequest) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *DescribeCategoriesRequest) GetOwner() []string {
	if m != nil {
		return m.Owner
	}
	return nil
}

type DescribeCategoriesResponse struct {
	TotalCount           uint32      `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	CategorySet          []*Category `protobuf:"bytes,2,rep,name=category_set,json=categorySet" json:"category_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DescribeCategoriesResponse) Reset()         { *m = DescribeCategoriesResponse{} }
func (m *DescribeCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeCategoriesResponse) ProtoMessage()    {}
func (*DescribeCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_01b76ca5827ebaa0, []int{2}
}
func (m *DescribeCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeCategoriesResponse.Unmarshal(m, b)
}
func (m *DescribeCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeCategoriesResponse.Marshal(b, m, deterministic)
}
func (dst *DescribeCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeCategoriesResponse.Merge(dst, src)
}
func (m *DescribeCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeCategoriesResponse.Size(m)
}
func (m *DescribeCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeCategoriesResponse proto.InternalMessageInfo

func (m *DescribeCategoriesResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeCategoriesResponse) GetCategorySet() []*Category {
	if m != nil {
		return m.CategorySet
	}
	return nil
}

type CreateCategoryRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Locale               *wrappers.StringValue `protobuf:"bytes,3,opt,name=locale" json:"locale,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateCategoryRequest) Reset()         { *m = CreateCategoryRequest{} }
func (m *CreateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryRequest) ProtoMessage()    {}
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_01b76ca5827ebaa0, []int{3}
}
func (m *CreateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryRequest.Unmarshal(m, b)
}
func (m *CreateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryRequest.Merge(dst, src)
}
func (m *CreateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryRequest.Size(m)
}
func (m *CreateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryRequest proto.InternalMessageInfo

func (m *CreateCategoryRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateCategoryRequest) GetLocale() *wrappers.StringValue {
	if m != nil {
		return m.Locale
	}
	return nil
}

type CreateCategoryResponse struct {
	CategoryId           *wrappers.StringValue `protobuf:"bytes,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateCategoryResponse) Reset()         { *m = CreateCategoryResponse{} }
func (m *CreateCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryResponse) ProtoMessage()    {}
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_01b76ca5827ebaa0, []int{4}
}
func (m *CreateCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryResponse.Unmarshal(m, b)
}
func (m *CreateCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryResponse.Marshal(b, m, deterministic)
}
func (dst *CreateCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryResponse.Merge(dst, src)
}
func (m *CreateCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryResponse.Size(m)
}
func (m *CreateCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryResponse proto.InternalMessageInfo

func (m *CreateCategoryResponse) GetCategoryId() *wrappers.StringValue {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

type ModifyCategoryRequest struct {
	CategoryId           *wrappers.StringValue `protobuf:"bytes,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Locale               *wrappers.StringValue `protobuf:"bytes,3,opt,name=locale" json:"locale,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ModifyCategoryRequest) Reset()         { *m = ModifyCategoryRequest{} }
func (m *ModifyCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyCategoryRequest) ProtoMessage()    {}
func (*ModifyCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_01b76ca5827ebaa0, []int{5}
}
func (m *ModifyCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyCategoryRequest.Unmarshal(m, b)
}
func (m *ModifyCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyCategoryRequest.Marshal(b, m, deterministic)
}
func (dst *ModifyCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyCategoryRequest.Merge(dst, src)
}
func (m *ModifyCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyCategoryRequest.Size(m)
}
func (m *ModifyCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyCategoryRequest proto.InternalMessageInfo

func (m *ModifyCategoryRequest) GetCategoryId() *wrappers.StringValue {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

func (m *ModifyCategoryRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ModifyCategoryRequest) GetLocale() *wrappers.StringValue {
	if m != nil {
		return m.Locale
	}
	return nil
}

type ModifyCategoryResponse struct {
	CategoryId           *wrappers.StringValue `protobuf:"bytes,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ModifyCategoryResponse) Reset()         { *m = ModifyCategoryResponse{} }
func (m *ModifyCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyCategoryResponse) ProtoMessage()    {}
func (*ModifyCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_01b76ca5827ebaa0, []int{6}
}
func (m *ModifyCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyCategoryResponse.Unmarshal(m, b)
}
func (m *ModifyCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyCategoryResponse.Marshal(b, m, deterministic)
}
func (dst *ModifyCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyCategoryResponse.Merge(dst, src)
}
func (m *ModifyCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyCategoryResponse.Size(m)
}
func (m *ModifyCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyCategoryResponse proto.InternalMessageInfo

func (m *ModifyCategoryResponse) GetCategoryId() *wrappers.StringValue {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

type DeleteCategoriesRequest struct {
	CategoryId           []string `protobuf:"bytes,1,rep,name=category_id,json=categoryId" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCategoriesRequest) Reset()         { *m = DeleteCategoriesRequest{} }
func (m *DeleteCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCategoriesRequest) ProtoMessage()    {}
func (*DeleteCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_01b76ca5827ebaa0, []int{7}
}
func (m *DeleteCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCategoriesRequest.Unmarshal(m, b)
}
func (m *DeleteCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCategoriesRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCategoriesRequest.Merge(dst, src)
}
func (m *DeleteCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCategoriesRequest.Size(m)
}
func (m *DeleteCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCategoriesRequest proto.InternalMessageInfo

func (m *DeleteCategoriesRequest) GetCategoryId() []string {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

type DeleteCategoriesResponse struct {
	CategoryId           []string `protobuf:"bytes,1,rep,name=category_id,json=categoryId" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCategoriesResponse) Reset()         { *m = DeleteCategoriesResponse{} }
func (m *DeleteCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCategoriesResponse) ProtoMessage()    {}
func (*DeleteCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_01b76ca5827ebaa0, []int{8}
}
func (m *DeleteCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCategoriesResponse.Unmarshal(m, b)
}
func (m *DeleteCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCategoriesResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCategoriesResponse.Merge(dst, src)
}
func (m *DeleteCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCategoriesResponse.Size(m)
}
func (m *DeleteCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCategoriesResponse proto.InternalMessageInfo

func (m *DeleteCategoriesResponse) GetCategoryId() []string {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

func init() {
	proto.RegisterType((*Category)(nil), "openpitrix.Category")
	proto.RegisterType((*DescribeCategoriesRequest)(nil), "openpitrix.DescribeCategoriesRequest")
	proto.RegisterType((*DescribeCategoriesResponse)(nil), "openpitrix.DescribeCategoriesResponse")
	proto.RegisterType((*CreateCategoryRequest)(nil), "openpitrix.CreateCategoryRequest")
	proto.RegisterType((*CreateCategoryResponse)(nil), "openpitrix.CreateCategoryResponse")
	proto.RegisterType((*ModifyCategoryRequest)(nil), "openpitrix.ModifyCategoryRequest")
	proto.RegisterType((*ModifyCategoryResponse)(nil), "openpitrix.ModifyCategoryResponse")
	proto.RegisterType((*DeleteCategoriesRequest)(nil), "openpitrix.DeleteCategoriesRequest")
	proto.RegisterType((*DeleteCategoriesResponse)(nil), "openpitrix.DeleteCategoriesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CategoryManagerClient is the client API for CategoryManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CategoryManagerClient interface {
	DescribeCategories(ctx context.Context, in *DescribeCategoriesRequest, opts ...grpc.CallOption) (*DescribeCategoriesResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	ModifyCategory(ctx context.Context, in *ModifyCategoryRequest, opts ...grpc.CallOption) (*ModifyCategoryResponse, error)
	DeleteCategories(ctx context.Context, in *DeleteCategoriesRequest, opts ...grpc.CallOption) (*DeleteCategoriesResponse, error)
}

type categoryManagerClient struct {
	cc *grpc.ClientConn
}

func NewCategoryManagerClient(cc *grpc.ClientConn) CategoryManagerClient {
	return &categoryManagerClient{cc}
}

func (c *categoryManagerClient) DescribeCategories(ctx context.Context, in *DescribeCategoriesRequest, opts ...grpc.CallOption) (*DescribeCategoriesResponse, error) {
	out := new(DescribeCategoriesResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.CategoryManager/DescribeCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryManagerClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.CategoryManager/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryManagerClient) ModifyCategory(ctx context.Context, in *ModifyCategoryRequest, opts ...grpc.CallOption) (*ModifyCategoryResponse, error) {
	out := new(ModifyCategoryResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.CategoryManager/ModifyCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryManagerClient) DeleteCategories(ctx context.Context, in *DeleteCategoriesRequest, opts ...grpc.CallOption) (*DeleteCategoriesResponse, error) {
	out := new(DeleteCategoriesResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.CategoryManager/DeleteCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryManagerServer is the server API for CategoryManager service.
type CategoryManagerServer interface {
	DescribeCategories(context.Context, *DescribeCategoriesRequest) (*DescribeCategoriesResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	ModifyCategory(context.Context, *ModifyCategoryRequest) (*ModifyCategoryResponse, error)
	DeleteCategories(context.Context, *DeleteCategoriesRequest) (*DeleteCategoriesResponse, error)
}

func RegisterCategoryManagerServer(s *grpc.Server, srv CategoryManagerServer) {
	s.RegisterService(&_CategoryManager_serviceDesc, srv)
}

func _CategoryManager_DescribeCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryManagerServer).DescribeCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.CategoryManager/DescribeCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryManagerServer).DescribeCategories(ctx, req.(*DescribeCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryManager_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryManagerServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.CategoryManager/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryManagerServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryManager_ModifyCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryManagerServer).ModifyCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.CategoryManager/ModifyCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryManagerServer).ModifyCategory(ctx, req.(*ModifyCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryManager_DeleteCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryManagerServer).DeleteCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.CategoryManager/DeleteCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryManagerServer).DeleteCategories(ctx, req.(*DeleteCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CategoryManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.CategoryManager",
	HandlerType: (*CategoryManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeCategories",
			Handler:    _CategoryManager_DescribeCategories_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _CategoryManager_CreateCategory_Handler,
		},
		{
			MethodName: "ModifyCategory",
			Handler:    _CategoryManager_ModifyCategory_Handler,
		},
		{
			MethodName: "DeleteCategories",
			Handler:    _CategoryManager_DeleteCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category.proto",
}

func init() { proto.RegisterFile("category.proto", fileDescriptor_category_01b76ca5827ebaa0) }

var fileDescriptor_category_01b76ca5827ebaa0 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xe5, 0x34, 0x89, 0xe0, 0x84, 0xde, 0xa6, 0x17, 0x42, 0x54, 0x91, 0xd4, 0x5c, 0x54,
	0x55, 0x34, 0x86, 0x80, 0x84, 0xd4, 0x8a, 0x45, 0x69, 0x37, 0x2c, 0xba, 0x71, 0x11, 0x95, 0xd8,
	0x44, 0x13, 0xfb, 0xc4, 0x1d, 0xc9, 0xf1, 0x98, 0xf1, 0xa4, 0xa1, 0x2b, 0x24, 0x16, 0x2c, 0x59,
	0x94, 0x07, 0xe0, 0x4d, 0x78, 0x89, 0x4a, 0x6c, 0xd9, 0xf0, 0x20, 0xc8, 0xe3, 0x71, 0x5a, 0x3b,
	0xa6, 0x8d, 0x04, 0x42, 0x62, 0x55, 0xcd, 0x9c, 0xff, 0xaf, 0xbf, 0x39, 0xe7, 0x3f, 0x81, 0x39,
	0x87, 0x4a, 0xf4, 0xb8, 0x38, 0x6d, 0x87, 0x82, 0x4b, 0x4e, 0x80, 0x87, 0x18, 0x84, 0x4c, 0x0a,
	0xf6, 0xbe, 0x71, 0xd7, 0xe3, 0xdc, 0xf3, 0xd1, 0x52, 0x95, 0xde, 0xb0, 0x6f, 0x8d, 0x04, 0x0d,
	0x43, 0x14, 0x51, 0xa2, 0x6d, 0x34, 0xf3, 0x75, 0xc9, 0x06, 0x18, 0x49, 0x3a, 0x08, 0xb5, 0x60,
	0x4d, 0x0b, 0x68, 0xc8, 0x2c, 0x1a, 0x04, 0x5c, 0x52, 0xc9, 0x78, 0x90, 0xda, 0x1f, 0xa9, 0x3f,
	0xce, 0x96, 0x87, 0xc1, 0x56, 0x34, 0xa2, 0x9e, 0x87, 0xc2, 0xe2, 0xa1, 0x52, 0x4c, 0xaa, 0xcd,
	0x1f, 0x25, 0xb8, 0xb1, 0xa7, 0x59, 0xc9, 0x0b, 0xa8, 0xa5, 0xdc, 0x5d, 0xe6, 0xd6, 0x8d, 0x96,
	0xb1, 0x51, 0xeb, 0xac, 0xb5, 0x93, 0xcf, 0xb5, 0x53, 0x9e, 0xf6, 0xa1, 0x14, 0x2c, 0xf0, 0xde,
	0x50, 0x7f, 0x88, 0x36, 0xa4, 0x86, 0x57, 0x2e, 0x79, 0x0c, 0xe5, 0x80, 0x0e, 0xb0, 0x5e, 0x9a,
	0xc2, 0xa7, 0x94, 0xe4, 0x19, 0x54, 0x7d, 0xee, 0x50, 0x1f, 0xeb, 0x33, 0x53, 0x78, 0xb4, 0x96,
	0x74, 0xa0, 0xc2, 0x47, 0x01, 0x8a, 0x7a, 0x79, 0x0a, 0x53, 0x22, 0x25, 0x3b, 0x50, 0x73, 0x04,
	0x52, 0x89, 0xdd, 0xb8, 0x9b, 0xf5, 0x8a, 0x72, 0x36, 0x26, 0x9c, 0xaf, 0xd3, 0x56, 0xdb, 0x90,
	0xc8, 0xe3, 0x8b, 0xd8, 0x3c, 0x0c, 0xdd, 0xb1, 0xb9, 0x7a, 0xbd, 0x39, 0x91, 0xc7, 0x17, 0xe6,
	0xb9, 0x01, 0x77, 0xf6, 0x31, 0x72, 0x04, 0xeb, 0xa1, 0xee, 0x34, 0xc3, 0xc8, 0xc6, 0x77, 0x43,
	0x8c, 0x64, 0xdc, 0xf2, 0x08, 0xa9, 0x70, 0x8e, 0xbb, 0x23, 0x2e, 0xa6, 0x6c, 0x79, 0x62, 0x38,
	0xe2, 0xc2, 0x25, 0xcb, 0x50, 0xf1, 0xd9, 0x80, 0x49, 0xd5, 0xf3, 0x59, 0x3b, 0x39, 0x90, 0x55,
	0xa8, 0xf2, 0x7e, 0x3f, 0x42, 0xa9, 0xda, 0x3a, 0x6b, 0xeb, 0x13, 0x69, 0x66, 0xe7, 0x5b, 0x6e,
	0xcd, 0x6c, 0xdc, 0xcc, 0x4c, 0x90, 0xe8, 0x09, 0x56, 0x54, 0x25, 0x99, 0xd1, 0x72, 0xda, 0xed,
	0xaa, 0xba, 0x4c, 0x0e, 0xe6, 0x09, 0x34, 0x8a, 0x1e, 0x15, 0x85, 0x3c, 0x88, 0x30, 0xfe, 0x90,
	0xe4, 0x92, 0xfa, 0x5d, 0x87, 0x0f, 0x03, 0xa9, 0x5e, 0x35, 0x6b, 0x83, 0xba, 0xda, 0x8b, 0x6f,
	0xc8, 0x73, 0xb8, 0x35, 0x26, 0x89, 0x39, 0x4b, 0xad, 0x99, 0x8d, 0x5a, 0x67, 0xb9, 0x7d, 0xb1,
	0x26, 0xed, 0x34, 0x95, 0xf6, 0x98, 0xf9, 0x10, 0xa5, 0xf9, 0x01, 0x56, 0xf6, 0xd4, 0x60, 0xc6,
	0x65, 0xdd, 0xc8, 0x7f, 0x14, 0x3e, 0xf3, 0x08, 0x56, 0xf3, 0x00, 0xfa, 0xd1, 0x7f, 0xb6, 0x3d,
	0xe6, 0x37, 0x03, 0x56, 0x0e, 0xb8, 0xcb, 0xfa, 0xa7, 0xf9, 0xa7, 0xfd, 0x1f, 0x6b, 0x19, 0x77,
	0x26, 0xcf, 0xff, 0x77, 0x3a, 0xb3, 0x0d, 0xb7, 0xf7, 0xd1, 0x47, 0x59, 0xb0, 0x3e, 0xcd, 0xfc,
	0x7f, 0xce, 0x25, 0xda, 0xdc, 0x81, 0xfa, 0xa4, 0xf7, 0x22, 0xa5, 0x57, 0x9a, 0x3b, 0xdf, 0xcb,
	0x30, 0x9f, 0x3e, 0xe6, 0x80, 0x06, 0xd4, 0x43, 0x41, 0xbe, 0x1a, 0x40, 0x26, 0x93, 0x4f, 0x1e,
	0x5c, 0x8e, 0xee, 0x6f, 0xd7, 0xbd, 0xf1, 0xf0, 0x3a, 0x59, 0x82, 0x66, 0x6e, 0x9f, 0xed, 0xae,
	0x93, 0xa6, 0xab, 0x05, 0x2d, 0x67, 0xac, 0x68, 0x8d, 0x98, 0x3c, 0x6e, 0xf5, 0x99, 0x2f, 0x51,
	0x7c, 0x3c, 0xff, 0xf9, 0xa5, 0xb4, 0x40, 0xe6, 0xac, 0x93, 0x27, 0xd6, 0x85, 0x82, 0x7c, 0x32,
	0x60, 0x2e, 0x1b, 0x51, 0xb2, 0x9e, 0x59, 0xac, 0xa2, 0xfd, 0x69, 0x98, 0x57, 0x49, 0x34, 0xd5,
	0xd6, 0xd9, 0xee, 0x22, 0x99, 0x4f, 0x7e, 0x18, 0x53, 0xa6, 0x53, 0x45, 0xb1, 0x64, 0xe6, 0x28,
	0xb6, 0x8d, 0x4d, 0x05, 0x92, 0x4d, 0x44, 0x16, 0xa4, 0x30, 0xed, 0x59, 0x90, 0xe2, 0x40, 0x69,
	0x90, 0x81, 0x2a, 0xe6, 0x40, 0x3a, 0x05, 0x20, 0x9f, 0x0d, 0x58, 0xc8, 0xa7, 0x80, 0xdc, 0xcb,
	0x8e, 0xa2, 0x30, 0x5f, 0x8d, 0xfb, 0x57, 0x8b, 0x34, 0x8e, 0x75, 0xb6, 0xbb, 0x44, 0x16, 0x5d,
	0x55, 0xbe, 0x34, 0xab, 0x04, 0x68, 0x73, 0x12, 0xe8, 0x65, 0xf9, 0x6d, 0x29, 0xec, 0xf5, 0xaa,
	0x2a, 0xf9, 0x4f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x74, 0x1c, 0x81, 0xa1, 0x2d, 0x08, 0x00,
	0x00,
}