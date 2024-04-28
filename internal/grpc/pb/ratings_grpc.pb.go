// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: ratings.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	GetRatingByIdWithWeight(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsResponse, error)
	GetRatingsBetweenDates(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsResponse, error)
	GetRatingsBetweenDatesAggregated(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsAggregatedResponse, error)
	GetOverallRatingsBetweenDates(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetOverallRatingsResponse, error)
	GetOverallRatingsChangeBetweenDates(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetOverallRatingsChangeResponse, error)
	GetCategoryRatingsBetweenDatesAggregated(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetCategoryRatingsAggregatedResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) GetRatingByIdWithWeight(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsResponse, error) {
	out := new(GetRatingsResponse)
	err := c.cc.Invoke(ctx, "/ratings.RatingService/GetRatingByIdWithWeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetRatingsBetweenDates(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsResponse, error) {
	out := new(GetRatingsResponse)
	err := c.cc.Invoke(ctx, "/ratings.RatingService/GetRatingsBetweenDates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetRatingsBetweenDatesAggregated(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsAggregatedResponse, error) {
	out := new(GetRatingsAggregatedResponse)
	err := c.cc.Invoke(ctx, "/ratings.RatingService/GetRatingsBetweenDatesAggregated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetOverallRatingsBetweenDates(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetOverallRatingsResponse, error) {
	out := new(GetOverallRatingsResponse)
	err := c.cc.Invoke(ctx, "/ratings.RatingService/GetOverallRatingsBetweenDates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetOverallRatingsChangeBetweenDates(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetOverallRatingsChangeResponse, error) {
	out := new(GetOverallRatingsChangeResponse)
	err := c.cc.Invoke(ctx, "/ratings.RatingService/GetOverallRatingsChangeBetweenDates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetCategoryRatingsBetweenDatesAggregated(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetCategoryRatingsAggregatedResponse, error) {
	out := new(GetCategoryRatingsAggregatedResponse)
	err := c.cc.Invoke(ctx, "/ratings.RatingService/GetCategoryRatingsBetweenDatesAggregated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	GetRatingByIdWithWeight(context.Context, *GetRatingsRequest) (*GetRatingsResponse, error)
	GetRatingsBetweenDates(context.Context, *GetRatingsRequest) (*GetRatingsResponse, error)
	GetRatingsBetweenDatesAggregated(context.Context, *GetRatingsRequest) (*GetRatingsAggregatedResponse, error)
	GetOverallRatingsBetweenDates(context.Context, *GetRatingsRequest) (*GetOverallRatingsResponse, error)
	GetOverallRatingsChangeBetweenDates(context.Context, *GetRatingsRequest) (*GetOverallRatingsChangeResponse, error)
	GetCategoryRatingsBetweenDatesAggregated(context.Context, *GetRatingsRequest) (*GetCategoryRatingsAggregatedResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) GetRatingByIdWithWeight(context.Context, *GetRatingsRequest) (*GetRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatingByIdWithWeight not implemented")
}
func (UnimplementedRatingServiceServer) GetRatingsBetweenDates(context.Context, *GetRatingsRequest) (*GetRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatingsBetweenDates not implemented")
}
func (UnimplementedRatingServiceServer) GetRatingsBetweenDatesAggregated(context.Context, *GetRatingsRequest) (*GetRatingsAggregatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatingsBetweenDatesAggregated not implemented")
}
func (UnimplementedRatingServiceServer) GetOverallRatingsBetweenDates(context.Context, *GetRatingsRequest) (*GetOverallRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverallRatingsBetweenDates not implemented")
}
func (UnimplementedRatingServiceServer) GetOverallRatingsChangeBetweenDates(context.Context, *GetRatingsRequest) (*GetOverallRatingsChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverallRatingsChangeBetweenDates not implemented")
}
func (UnimplementedRatingServiceServer) GetCategoryRatingsBetweenDatesAggregated(context.Context, *GetRatingsRequest) (*GetCategoryRatingsAggregatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryRatingsBetweenDatesAggregated not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_GetRatingByIdWithWeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRatingByIdWithWeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratings.RatingService/GetRatingByIdWithWeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRatingByIdWithWeight(ctx, req.(*GetRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetRatingsBetweenDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRatingsBetweenDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratings.RatingService/GetRatingsBetweenDates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRatingsBetweenDates(ctx, req.(*GetRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetRatingsBetweenDatesAggregated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRatingsBetweenDatesAggregated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratings.RatingService/GetRatingsBetweenDatesAggregated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRatingsBetweenDatesAggregated(ctx, req.(*GetRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetOverallRatingsBetweenDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetOverallRatingsBetweenDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratings.RatingService/GetOverallRatingsBetweenDates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetOverallRatingsBetweenDates(ctx, req.(*GetRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetOverallRatingsChangeBetweenDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetOverallRatingsChangeBetweenDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratings.RatingService/GetOverallRatingsChangeBetweenDates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetOverallRatingsChangeBetweenDates(ctx, req.(*GetRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetCategoryRatingsBetweenDatesAggregated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetCategoryRatingsBetweenDatesAggregated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratings.RatingService/GetCategoryRatingsBetweenDatesAggregated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetCategoryRatingsBetweenDatesAggregated(ctx, req.(*GetRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ratings.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRatingByIdWithWeight",
			Handler:    _RatingService_GetRatingByIdWithWeight_Handler,
		},
		{
			MethodName: "GetRatingsBetweenDates",
			Handler:    _RatingService_GetRatingsBetweenDates_Handler,
		},
		{
			MethodName: "GetRatingsBetweenDatesAggregated",
			Handler:    _RatingService_GetRatingsBetweenDatesAggregated_Handler,
		},
		{
			MethodName: "GetOverallRatingsBetweenDates",
			Handler:    _RatingService_GetOverallRatingsBetweenDates_Handler,
		},
		{
			MethodName: "GetOverallRatingsChangeBetweenDates",
			Handler:    _RatingService_GetOverallRatingsChangeBetweenDates_Handler,
		},
		{
			MethodName: "GetCategoryRatingsBetweenDatesAggregated",
			Handler:    _RatingService_GetCategoryRatingsBetweenDatesAggregated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ratings.proto",
}