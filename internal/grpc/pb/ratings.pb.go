// Command used in this folder: protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\ratings.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: ratings.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRatingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartDate string `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *GetRatingsRequest) Reset() {
	*x = GetRatingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingsRequest) ProtoMessage() {}

func (x *GetRatingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingsRequest.ProtoReflect.Descriptor instead.
func (*GetRatingsRequest) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{0}
}

func (x *GetRatingsRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRatingsRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetRatingsRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type RatingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Rating      float64 `protobuf:"fixed64,3,opt,name=rating,proto3" json:"rating,omitempty"`
	DateOrWeek  string  `protobuf:"bytes,4,opt,name=dateOrWeek,proto3" json:"dateOrWeek,omitempty"`
	Count       int32   `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	TotalRating int32   `protobuf:"varint,6,opt,name=total_rating,json=totalRating,proto3" json:"total_rating,omitempty"`
}

func (x *RatingResult) Reset() {
	*x = RatingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingResult) ProtoMessage() {}

func (x *RatingResult) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingResult.ProtoReflect.Descriptor instead.
func (*RatingResult) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{1}
}

func (x *RatingResult) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RatingResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RatingResult) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *RatingResult) GetDateOrWeek() string {
	if x != nil {
		return x.DateOrWeek
	}
	return ""
}

func (x *RatingResult) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RatingResult) GetTotalRating() int32 {
	if x != nil {
		return x.TotalRating
	}
	return 0
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Rating      int32  `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	DateOrWeek  string `protobuf:"bytes,4,opt,name=dateOrWeek,proto3" json:"dateOrWeek,omitempty"`
	Count       int32  `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	TotalRating int32  `protobuf:"varint,6,opt,name=total_rating,json=totalRating,proto3" json:"total_rating,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{2}
}

func (x *Rating) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rating) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rating) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Rating) GetDateOrWeek() string {
	if x != nil {
		return x.DateOrWeek
	}
	return ""
}

func (x *Rating) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Rating) GetTotalRating() int32 {
	if x != nil {
		return x.TotalRating
	}
	return 0
}

type GetRatingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating []*Rating `protobuf:"bytes,1,rep,name=rating,proto3" json:"rating,omitempty"`
}

func (x *GetRatingsResponse) Reset() {
	*x = GetRatingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingsResponse) ProtoMessage() {}

func (x *GetRatingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingsResponse.ProtoReflect.Descriptor instead.
func (*GetRatingsResponse) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{3}
}

func (x *GetRatingsResponse) GetRating() []*Rating {
	if x != nil {
		return x.Rating
	}
	return nil
}

type GetOverallRatingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OverallRating int32 `protobuf:"varint,1,opt,name=overall_rating,json=overallRating,proto3" json:"overall_rating,omitempty"`
}

func (x *GetOverallRatingsResponse) Reset() {
	*x = GetOverallRatingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOverallRatingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverallRatingsResponse) ProtoMessage() {}

func (x *GetOverallRatingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverallRatingsResponse.ProtoReflect.Descriptor instead.
func (*GetOverallRatingsResponse) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{4}
}

func (x *GetOverallRatingsResponse) GetOverallRating() int32 {
	if x != nil {
		return x.OverallRating
	}
	return 0
}

type GetOverallRatingsChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PercentageChange int32 `protobuf:"varint,1,opt,name=percentage_change,json=percentageChange,proto3" json:"percentage_change,omitempty"`
}

func (x *GetOverallRatingsChangeResponse) Reset() {
	*x = GetOverallRatingsChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOverallRatingsChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverallRatingsChangeResponse) ProtoMessage() {}

func (x *GetOverallRatingsChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverallRatingsChangeResponse.ProtoReflect.Descriptor instead.
func (*GetOverallRatingsChangeResponse) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{5}
}

func (x *GetOverallRatingsChangeResponse) GetPercentageChange() int32 {
	if x != nil {
		return x.PercentageChange
	}
	return 0
}

type GetRatingsAggregatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingGroup []*RatingGroup `protobuf:"bytes,1,rep,name=rating_group,json=ratingGroup,proto3" json:"rating_group,omitempty"`
}

func (x *GetRatingsAggregatedResponse) Reset() {
	*x = GetRatingsAggregatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingsAggregatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingsAggregatedResponse) ProtoMessage() {}

func (x *GetRatingsAggregatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingsAggregatedResponse.ProtoReflect.Descriptor instead.
func (*GetRatingsAggregatedResponse) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{6}
}

func (x *GetRatingsAggregatedResponse) GetRatingGroup() []*RatingGroup {
	if x != nil {
		return x.RatingGroup
	}
	return nil
}

type CategoryRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RatingCategoryDetail []*RatingCategoryDetail `protobuf:"bytes,3,rep,name=rating_category_detail,json=ratingCategoryDetail,proto3" json:"rating_category_detail,omitempty"`
}

func (x *CategoryRating) Reset() {
	*x = CategoryRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryRating) ProtoMessage() {}

func (x *CategoryRating) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryRating.ProtoReflect.Descriptor instead.
func (*CategoryRating) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{7}
}

func (x *CategoryRating) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CategoryRating) GetRatingCategoryDetail() []*RatingCategoryDetail {
	if x != nil {
		return x.RatingCategoryDetail
	}
	return nil
}

type RatingCategoryDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AverageRating int32  `protobuf:"varint,3,opt,name=average_rating,json=averageRating,proto3" json:"average_rating,omitempty"`
}

func (x *RatingCategoryDetail) Reset() {
	*x = RatingCategoryDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingCategoryDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingCategoryDetail) ProtoMessage() {}

func (x *RatingCategoryDetail) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingCategoryDetail.ProtoReflect.Descriptor instead.
func (*RatingCategoryDetail) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{8}
}

func (x *RatingCategoryDetail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RatingCategoryDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RatingCategoryDetail) GetAverageRating() int32 {
	if x != nil {
		return x.AverageRating
	}
	return 0
}

type GetCategoryRatingsAggregatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryRating []*CategoryRating `protobuf:"bytes,1,rep,name=category_rating,json=categoryRating,proto3" json:"category_rating,omitempty"`
}

func (x *GetCategoryRatingsAggregatedResponse) Reset() {
	*x = GetCategoryRatingsAggregatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryRatingsAggregatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRatingsAggregatedResponse) ProtoMessage() {}

func (x *GetCategoryRatingsAggregatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRatingsAggregatedResponse.ProtoReflect.Descriptor instead.
func (*GetCategoryRatingsAggregatedResponse) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{9}
}

func (x *GetCategoryRatingsAggregatedResponse) GetCategoryRating() []*CategoryRating {
	if x != nil {
		return x.CategoryRating
	}
	return nil
}

type RatingGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Count   int32           `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Ratings []*RatingDetail `protobuf:"bytes,4,rep,name=ratings,proto3" json:"ratings,omitempty"`
}

func (x *RatingGroup) Reset() {
	*x = RatingGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingGroup) ProtoMessage() {}

func (x *RatingGroup) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingGroup.ProtoReflect.Descriptor instead.
func (*RatingGroup) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{10}
}

func (x *RatingGroup) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RatingGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RatingGroup) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RatingGroup) GetRatings() []*RatingDetail {
	if x != nil {
		return x.Ratings
	}
	return nil
}

type RatingDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateOrWeek string `protobuf:"bytes,1,opt,name=date_or_week,json=dateOrWeek,proto3" json:"date_or_week,omitempty"`
	Rating     int32  `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *RatingDetail) Reset() {
	*x = RatingDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratings_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingDetail) ProtoMessage() {}

func (x *RatingDetail) ProtoReflect() protoreflect.Message {
	mi := &file_ratings_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingDetail.ProtoReflect.Descriptor instead.
func (*RatingDetail) Descriptor() ([]byte, []int) {
	return file_ratings_proto_rawDescGZIP(), []int{11}
}

func (x *RatingDetail) GetDateOrWeek() string {
	if x != nil {
		return x.DateOrWeek
	}
	return ""
}

func (x *RatingDetail) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

var File_ratings_proto protoreflect.FileDescriptor

var file_ratings_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x5d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x57, 0x65,
	0x65, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x57, 0x65, 0x65, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x9d, 0x01,
	0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x57, 0x65,
	0x65, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x57, 0x65, 0x65, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x3d, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x42, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x4e, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x22, 0x57, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0b, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x75, 0x0a, 0x0e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x53, 0x0a, 0x16, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x14, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x22, 0x61, 0x0a, 0x14, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x22, 0x68, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x78, 0x0a,
	0x0b, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x48, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x72, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x32, 0xe2, 0x04, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x42, 0x79, 0x49, 0x64, 0x57, 0x69, 0x74, 0x68, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a,
	0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x20, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a,
	0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5f, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72,
	0x61, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6b, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x65, 0x74,
	0x77, 0x65, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x75, 0x0a, 0x28, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x2e, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x64, 0x75, 0x75, 0x6d, 0x61, 0x72, 0x2f, 0x6b,
	0x6c, 0x61, 0x75, 0x73, 0x61, 0x70, 0x70, 0x2d, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x74, 0x61,
	0x73, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ratings_proto_rawDescOnce sync.Once
	file_ratings_proto_rawDescData = file_ratings_proto_rawDesc
)

func file_ratings_proto_rawDescGZIP() []byte {
	file_ratings_proto_rawDescOnce.Do(func() {
		file_ratings_proto_rawDescData = protoimpl.X.CompressGZIP(file_ratings_proto_rawDescData)
	})
	return file_ratings_proto_rawDescData
}

var file_ratings_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_ratings_proto_goTypes = []interface{}{
	(*GetRatingsRequest)(nil),                    // 0: ratings.GetRatingsRequest
	(*RatingResult)(nil),                         // 1: ratings.RatingResult
	(*Rating)(nil),                               // 2: ratings.Rating
	(*GetRatingsResponse)(nil),                   // 3: ratings.GetRatingsResponse
	(*GetOverallRatingsResponse)(nil),            // 4: ratings.GetOverallRatingsResponse
	(*GetOverallRatingsChangeResponse)(nil),      // 5: ratings.GetOverallRatingsChangeResponse
	(*GetRatingsAggregatedResponse)(nil),         // 6: ratings.GetRatingsAggregatedResponse
	(*CategoryRating)(nil),                       // 7: ratings.CategoryRating
	(*RatingCategoryDetail)(nil),                 // 8: ratings.RatingCategoryDetail
	(*GetCategoryRatingsAggregatedResponse)(nil), // 9: ratings.GetCategoryRatingsAggregatedResponse
	(*RatingGroup)(nil),                          // 10: ratings.RatingGroup
	(*RatingDetail)(nil),                         // 11: ratings.RatingDetail
}
var file_ratings_proto_depIdxs = []int32{
	2,  // 0: ratings.GetRatingsResponse.rating:type_name -> ratings.Rating
	10, // 1: ratings.GetRatingsAggregatedResponse.rating_group:type_name -> ratings.RatingGroup
	8,  // 2: ratings.CategoryRating.rating_category_detail:type_name -> ratings.RatingCategoryDetail
	7,  // 3: ratings.GetCategoryRatingsAggregatedResponse.category_rating:type_name -> ratings.CategoryRating
	11, // 4: ratings.RatingGroup.ratings:type_name -> ratings.RatingDetail
	0,  // 5: ratings.RatingService.GetRatingByIdWithWeight:input_type -> ratings.GetRatingsRequest
	0,  // 6: ratings.RatingService.GetRatingsBetweenDates:input_type -> ratings.GetRatingsRequest
	0,  // 7: ratings.RatingService.GetRatingsBetweenDatesAggregated:input_type -> ratings.GetRatingsRequest
	0,  // 8: ratings.RatingService.GetOverallRatingsBetweenDates:input_type -> ratings.GetRatingsRequest
	0,  // 9: ratings.RatingService.GetOverallRatingsChangeBetweenDates:input_type -> ratings.GetRatingsRequest
	0,  // 10: ratings.RatingService.GetCategoryRatingsBetweenDatesAggregated:input_type -> ratings.GetRatingsRequest
	3,  // 11: ratings.RatingService.GetRatingByIdWithWeight:output_type -> ratings.GetRatingsResponse
	3,  // 12: ratings.RatingService.GetRatingsBetweenDates:output_type -> ratings.GetRatingsResponse
	6,  // 13: ratings.RatingService.GetRatingsBetweenDatesAggregated:output_type -> ratings.GetRatingsAggregatedResponse
	4,  // 14: ratings.RatingService.GetOverallRatingsBetweenDates:output_type -> ratings.GetOverallRatingsResponse
	5,  // 15: ratings.RatingService.GetOverallRatingsChangeBetweenDates:output_type -> ratings.GetOverallRatingsChangeResponse
	9,  // 16: ratings.RatingService.GetCategoryRatingsBetweenDatesAggregated:output_type -> ratings.GetCategoryRatingsAggregatedResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_ratings_proto_init() }
func file_ratings_proto_init() {
	if File_ratings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ratings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatingsRequest); i {
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
		file_ratings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingResult); i {
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
		file_ratings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
		file_ratings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatingsResponse); i {
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
		file_ratings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOverallRatingsResponse); i {
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
		file_ratings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOverallRatingsChangeResponse); i {
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
		file_ratings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatingsAggregatedResponse); i {
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
		file_ratings_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryRating); i {
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
		file_ratings_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingCategoryDetail); i {
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
		file_ratings_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryRatingsAggregatedResponse); i {
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
		file_ratings_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingGroup); i {
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
		file_ratings_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingDetail); i {
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
			RawDescriptor: file_ratings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ratings_proto_goTypes,
		DependencyIndexes: file_ratings_proto_depIdxs,
		MessageInfos:      file_ratings_proto_msgTypes,
	}.Build()
	File_ratings_proto = out.File
	file_ratings_proto_rawDesc = nil
	file_ratings_proto_goTypes = nil
	file_ratings_proto_depIdxs = nil
}
