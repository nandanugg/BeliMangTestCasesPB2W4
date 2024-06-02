// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: contract.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ZoneOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Area                     float32 `protobuf:"fixed32,1,opt,name=area,proto3" json:"area,omitempty"`
	Gap                      float32 `protobuf:"fixed32,2,opt,name=gap,proto3" json:"gap,omitempty"`
	NumberOfZones            int32   `protobuf:"varint,3,opt,name=numberOfZones,proto3" json:"numberOfZones,omitempty"`
	NumberOfMerchantsPerZone int32   `protobuf:"varint,4,opt,name=numberOfMerchantsPerZone,proto3" json:"numberOfMerchantsPerZone,omitempty"`
}

func (x *ZoneOpts) Reset() {
	*x = ZoneOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneOpts) ProtoMessage() {}

func (x *ZoneOpts) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneOpts.ProtoReflect.Descriptor instead.
func (*ZoneOpts) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneOpts) GetArea() float32 {
	if x != nil {
		return x.Area
	}
	return 0
}

func (x *ZoneOpts) GetGap() float32 {
	if x != nil {
		return x.Gap
	}
	return 0
}

func (x *ZoneOpts) GetNumberOfZones() int32 {
	if x != nil {
		return x.NumberOfZones
	}
	return 0
}

func (x *ZoneOpts) GetNumberOfMerchantsPerZone() int32 {
	if x != nil {
		return x.NumberOfMerchantsPerZone
	}
	return 0
}

type AssignMerchant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PregeneratedId string `protobuf:"bytes,1,opt,name=pregeneratedId,proto3" json:"pregeneratedId,omitempty"`
	MerchantId     string `protobuf:"bytes,2,opt,name=merchantId,proto3" json:"merchantId,omitempty"`
}

func (x *AssignMerchant) Reset() {
	*x = AssignMerchant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignMerchant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignMerchant) ProtoMessage() {}

func (x *AssignMerchant) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignMerchant.ProtoReflect.Descriptor instead.
func (*AssignMerchant) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

func (x *AssignMerchant) GetPregeneratedId() string {
	if x != nil {
		return x.PregeneratedId
	}
	return ""
}

func (x *AssignMerchant) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type LocationPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat  float32 `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Long float32 `protobuf:"fixed32,2,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *LocationPoint) Reset() {
	*x = LocationPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationPoint) ProtoMessage() {}

func (x *LocationPoint) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationPoint.ProtoReflect.Descriptor instead.
func (*LocationPoint) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

func (x *LocationPoint) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *LocationPoint) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

type Merchant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId     string         `protobuf:"bytes,1,opt,name=merchantId,proto3" json:"merchantId,omitempty"`
	PregeneratedId string         `protobuf:"bytes,2,opt,name=pregeneratedId,proto3" json:"pregeneratedId,omitempty"`
	Location       *LocationPoint `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Merchant) Reset() {
	*x = Merchant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Merchant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Merchant) ProtoMessage() {}

func (x *Merchant) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Merchant.ProtoReflect.Descriptor instead.
func (*Merchant) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{3}
}

func (x *Merchant) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *Merchant) GetPregeneratedId() string {
	if x != nil {
		return x.PregeneratedId
	}
	return ""
}

func (x *Merchant) GetLocation() *LocationPoint {
	if x != nil {
		return x.Location
	}
	return nil
}

type MerchantNearestRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartingPoint *LocationPoint      `protobuf:"bytes,1,opt,name=startingPoint,proto3" json:"startingPoint,omitempty"`
	TotalDuration int64               `protobuf:"varint,2,opt,name=totalDuration,proto3" json:"totalDuration,omitempty"`
	Merchants     map[int64]*Merchant `protobuf:"bytes,3,rep,name=merchants,proto3" json:"merchants,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MerchantNearestRecord) Reset() {
	*x = MerchantNearestRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantNearestRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantNearestRecord) ProtoMessage() {}

func (x *MerchantNearestRecord) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantNearestRecord.ProtoReflect.Descriptor instead.
func (*MerchantNearestRecord) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{4}
}

func (x *MerchantNearestRecord) GetStartingPoint() *LocationPoint {
	if x != nil {
		return x.StartingPoint
	}
	return nil
}

func (x *MerchantNearestRecord) GetTotalDuration() int64 {
	if x != nil {
		return x.TotalDuration
	}
	return 0
}

func (x *MerchantNearestRecord) GetMerchants() map[int64]*Merchant {
	if x != nil {
		return x.Merchants
	}
	return nil
}

type MerchantNearestRecordZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*MerchantNearestRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *MerchantNearestRecordZone) Reset() {
	*x = MerchantNearestRecordZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantNearestRecordZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantNearestRecordZone) ProtoMessage() {}

func (x *MerchantNearestRecordZone) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantNearestRecordZone.ProtoReflect.Descriptor instead.
func (*MerchantNearestRecordZone) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{5}
}

func (x *MerchantNearestRecordZone) GetRecords() []*MerchantNearestRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

type AllMerchantNearestRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zones []*MerchantNearestRecordZone `protobuf:"bytes,1,rep,name=zones,proto3" json:"zones,omitempty"`
}

func (x *AllMerchantNearestRecord) Reset() {
	*x = AllMerchantNearestRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllMerchantNearestRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllMerchantNearestRecord) ProtoMessage() {}

func (x *AllMerchantNearestRecord) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllMerchantNearestRecord.ProtoReflect.Descriptor instead.
func (*AllMerchantNearestRecord) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{6}
}

func (x *AllMerchantNearestRecord) GetZones() []*MerchantNearestRecordZone {
	if x != nil {
		return x.Zones
	}
	return nil
}

type PregeneratedMerchant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Merchant []*Merchant `protobuf:"bytes,1,rep,name=merchant,proto3" json:"merchant,omitempty"`
}

func (x *PregeneratedMerchant) Reset() {
	*x = PregeneratedMerchant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PregeneratedMerchant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PregeneratedMerchant) ProtoMessage() {}

func (x *PregeneratedMerchant) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PregeneratedMerchant.ProtoReflect.Descriptor instead.
func (*PregeneratedMerchant) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{7}
}

func (x *PregeneratedMerchant) GetMerchant() []*Merchant {
	if x != nil {
		return x.Merchant
	}
	return nil
}

type GeneratedRoutes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GeneratedRoutes map[int64]*Merchant `protobuf:"bytes,1,rep,name=generatedRoutes,proto3" json:"generatedRoutes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StartingPoint   *LocationPoint      `protobuf:"bytes,2,opt,name=startingPoint,proto3" json:"startingPoint,omitempty"`
	TotalDistance   float32             `protobuf:"fixed32,3,opt,name=totalDistance,proto3" json:"totalDistance,omitempty"`
	TotalDuration   int64               `protobuf:"varint,4,opt,name=totalDuration,proto3" json:"totalDuration,omitempty"`
	StartingIndex   int64               `protobuf:"varint,5,opt,name=startingIndex,proto3" json:"startingIndex,omitempty"`
}

func (x *GeneratedRoutes) Reset() {
	*x = GeneratedRoutes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratedRoutes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratedRoutes) ProtoMessage() {}

func (x *GeneratedRoutes) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratedRoutes.ProtoReflect.Descriptor instead.
func (*GeneratedRoutes) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{8}
}

func (x *GeneratedRoutes) GetGeneratedRoutes() map[int64]*Merchant {
	if x != nil {
		return x.GeneratedRoutes
	}
	return nil
}

func (x *GeneratedRoutes) GetStartingPoint() *LocationPoint {
	if x != nil {
		return x.StartingPoint
	}
	return nil
}

func (x *GeneratedRoutes) GetTotalDistance() float32 {
	if x != nil {
		return x.TotalDistance
	}
	return 0
}

func (x *GeneratedRoutes) GetTotalDuration() int64 {
	if x != nil {
		return x.TotalDuration
	}
	return 0
}

func (x *GeneratedRoutes) GetStartingIndex() int64 {
	if x != nil {
		return x.StartingIndex
	}
	return 0
}

type RouteZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*GeneratedRoutes `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *RouteZone) Reset() {
	*x = RouteZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteZone) ProtoMessage() {}

func (x *RouteZone) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteZone.ProtoReflect.Descriptor instead.
func (*RouteZone) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{9}
}

func (x *RouteZone) GetRoutes() []*GeneratedRoutes {
	if x != nil {
		return x.Routes
	}
	return nil
}

type AllGeneratedRoutes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zone []*RouteZone `protobuf:"bytes,1,rep,name=zone,proto3" json:"zone,omitempty"`
}

func (x *AllGeneratedRoutes) Reset() {
	*x = AllGeneratedRoutes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllGeneratedRoutes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllGeneratedRoutes) ProtoMessage() {}

func (x *AllGeneratedRoutes) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllGeneratedRoutes.ProtoReflect.Descriptor instead.
func (*AllGeneratedRoutes) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{10}
}

func (x *AllGeneratedRoutes) GetZone() []*RouteZone {
	if x != nil {
		return x.Zone
	}
	return nil
}

var File_contract_proto protoreflect.FileDescriptor

var file_contract_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x67, 0x61, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x50,
	0x65, 0x72, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x50,
	0x65, 0x72, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x58, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x35, 0x0a, 0x0d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8a, 0x02, 0x0a, 0x15,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x37, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x09, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x09, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x1a, 0x4a, 0x0a, 0x0e,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x50, 0x0a, 0x19, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x4f, 0x0a, 0x18, 0x41, 0x6c,
	0x6c, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x33, 0x0a, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x14, 0x50,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x22, 0xe2, 0x02,
	0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x52, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x1a, 0x50, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x38, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x12,
	0x41, 0x6c, 0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52,
	0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x32, 0xfd, 0x04, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x1a, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x50, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x56, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e, 0x65,
	0x61, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x46, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x6c, 0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x1c, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x1b, 0x49, 0x6e, 0x69, 0x74, 0x50, 0x65, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x54, 0x53, 0x50, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x54, 0x0a, 0x22, 0x49, 0x6e, 0x69, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_proto_rawDescOnce sync.Once
	file_contract_proto_rawDescData = file_contract_proto_rawDesc
)

func file_contract_proto_rawDescGZIP() []byte {
	file_contract_proto_rawDescOnce.Do(func() {
		file_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_proto_rawDescData)
	})
	return file_contract_proto_rawDescData
}

var file_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_contract_proto_goTypes = []interface{}{
	(*ZoneOpts)(nil),                  // 0: pb.ZoneOpts
	(*AssignMerchant)(nil),            // 1: pb.AssignMerchant
	(*LocationPoint)(nil),             // 2: pb.LocationPoint
	(*Merchant)(nil),                  // 3: pb.Merchant
	(*MerchantNearestRecord)(nil),     // 4: pb.MerchantNearestRecord
	(*MerchantNearestRecordZone)(nil), // 5: pb.MerchantNearestRecordZone
	(*AllMerchantNearestRecord)(nil),  // 6: pb.AllMerchantNearestRecord
	(*PregeneratedMerchant)(nil),      // 7: pb.PregeneratedMerchant
	(*GeneratedRoutes)(nil),           // 8: pb.GeneratedRoutes
	(*RouteZone)(nil),                 // 9: pb.RouteZone
	(*AllGeneratedRoutes)(nil),        // 10: pb.AllGeneratedRoutes
	nil,                               // 11: pb.MerchantNearestRecord.MerchantsEntry
	nil,                               // 12: pb.GeneratedRoutes.GeneratedRoutesEntry
	(*emptypb.Empty)(nil),             // 13: google.protobuf.Empty
}
var file_contract_proto_depIdxs = []int32{
	2,  // 0: pb.Merchant.location:type_name -> pb.LocationPoint
	2,  // 1: pb.MerchantNearestRecord.startingPoint:type_name -> pb.LocationPoint
	11, // 2: pb.MerchantNearestRecord.merchants:type_name -> pb.MerchantNearestRecord.MerchantsEntry
	4,  // 3: pb.MerchantNearestRecordZone.records:type_name -> pb.MerchantNearestRecord
	5,  // 4: pb.AllMerchantNearestRecord.zones:type_name -> pb.MerchantNearestRecordZone
	3,  // 5: pb.PregeneratedMerchant.merchant:type_name -> pb.Merchant
	12, // 6: pb.GeneratedRoutes.generatedRoutes:type_name -> pb.GeneratedRoutes.GeneratedRoutesEntry
	2,  // 7: pb.GeneratedRoutes.startingPoint:type_name -> pb.LocationPoint
	8,  // 8: pb.RouteZone.routes:type_name -> pb.GeneratedRoutes
	9,  // 9: pb.AllGeneratedRoutes.zone:type_name -> pb.RouteZone
	3,  // 10: pb.MerchantNearestRecord.MerchantsEntry.value:type_name -> pb.Merchant
	3,  // 11: pb.GeneratedRoutes.GeneratedRoutesEntry.value:type_name -> pb.Merchant
	1,  // 12: pb.MerchantService.AssignPregeneratedMerchant:input_type -> pb.AssignMerchant
	13, // 13: pb.MerchantService.GetAllMerchantNearestLocations:input_type -> google.protobuf.Empty
	13, // 14: pb.MerchantService.GetAllMerchantRoutes:input_type -> google.protobuf.Empty
	13, // 15: pb.MerchantService.GetAllPregeneratedMerchants:input_type -> google.protobuf.Empty
	13, // 16: pb.MerchantService.InitMerchantNearestLocations:input_type -> google.protobuf.Empty
	13, // 17: pb.MerchantService.InitPegeneratedTSPMerchants:input_type -> google.protobuf.Empty
	13, // 18: pb.MerchantService.InitZonesWithPregeneratedMerchants:input_type -> google.protobuf.Empty
	13, // 19: pb.MerchantService.ResetAll:input_type -> google.protobuf.Empty
	13, // 20: pb.MerchantService.AssignPregeneratedMerchant:output_type -> google.protobuf.Empty
	6,  // 21: pb.MerchantService.GetAllMerchantNearestLocations:output_type -> pb.AllMerchantNearestRecord
	10, // 22: pb.MerchantService.GetAllMerchantRoutes:output_type -> pb.AllGeneratedRoutes
	7,  // 23: pb.MerchantService.GetAllPregeneratedMerchants:output_type -> pb.PregeneratedMerchant
	13, // 24: pb.MerchantService.InitMerchantNearestLocations:output_type -> google.protobuf.Empty
	13, // 25: pb.MerchantService.InitPegeneratedTSPMerchants:output_type -> google.protobuf.Empty
	13, // 26: pb.MerchantService.InitZonesWithPregeneratedMerchants:output_type -> google.protobuf.Empty
	13, // 27: pb.MerchantService.ResetAll:output_type -> google.protobuf.Empty
	20, // [20:28] is the sub-list for method output_type
	12, // [12:20] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_contract_proto_init() }
func file_contract_proto_init() {
	if File_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneOpts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignMerchant); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Merchant); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantNearestRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantNearestRecordZone); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllMerchantNearestRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PregeneratedMerchant); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratedRoutes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteZone); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllGeneratedRoutes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_proto_goTypes,
		DependencyIndexes: file_contract_proto_depIdxs,
		MessageInfos:      file_contract_proto_msgTypes,
	}.Build()
	File_contract_proto = out.File
	file_contract_proto_rawDesc = nil
	file_contract_proto_goTypes = nil
	file_contract_proto_depIdxs = nil
}
