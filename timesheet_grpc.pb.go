// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package timesheet

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

// TimesheetServiceClient is the client API for TimesheetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimesheetServiceClient interface {
	// Create creates a new sample
	Create(ctx context.Context, in *CreateTimesheetRequest, opts ...grpc.CallOption) (*Timesheet, error)
	// Update updates an existent consultant
	Update(ctx context.Context, in *UpdateTimesheetRequest, opts ...grpc.CallOption) (*Timesheet, error)
	// Get retrieves a sample by id
	Get(ctx context.Context, in *TimesheetIdRequest, opts ...grpc.CallOption) (*Timesheet, error)
	// Search retrieves a sample by owner
	Search(ctx context.Context, in *SearchTimesheetRequest, opts ...grpc.CallOption) (*SearchTimesheetResponse, error)
	// Delete sample
	Delete(ctx context.Context, in *TimesheetIdRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Create creates a new sample
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*Event, error)
	// Update updates an existent consultant
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Event, error)
	// Get retrieves a sample by id
	GetEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*Event, error)
	// Search searches samples
	SearchEvents(ctx context.Context, in *TimesheetIdRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Delete sample
	DeleteEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type timesheetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimesheetServiceClient(cc grpc.ClientConnInterface) TimesheetServiceClient {
	return &timesheetServiceClient{cc}
}

func (c *timesheetServiceClient) Create(ctx context.Context, in *CreateTimesheetRequest, opts ...grpc.CallOption) (*Timesheet, error) {
	out := new(Timesheet)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) Update(ctx context.Context, in *UpdateTimesheetRequest, opts ...grpc.CallOption) (*Timesheet, error) {
	out := new(Timesheet)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) Get(ctx context.Context, in *TimesheetIdRequest, opts ...grpc.CallOption) (*Timesheet, error) {
	out := new(Timesheet)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) Search(ctx context.Context, in *SearchTimesheetRequest, opts ...grpc.CallOption) (*SearchTimesheetResponse, error) {
	out := new(SearchTimesheetResponse)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) Delete(ctx context.Context, in *TimesheetIdRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) GetEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) SearchEvents(ctx context.Context, in *TimesheetIdRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/SearchEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesheetServiceClient) DeleteEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/timesheet.TimesheetService/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimesheetServiceServer is the server API for TimesheetService service.
// All implementations must embed UnimplementedTimesheetServiceServer
// for forward compatibility
type TimesheetServiceServer interface {
	// Create creates a new sample
	Create(context.Context, *CreateTimesheetRequest) (*Timesheet, error)
	// Update updates an existent consultant
	Update(context.Context, *UpdateTimesheetRequest) (*Timesheet, error)
	// Get retrieves a sample by id
	Get(context.Context, *TimesheetIdRequest) (*Timesheet, error)
	// Search retrieves a sample by owner
	Search(context.Context, *SearchTimesheetRequest) (*SearchTimesheetResponse, error)
	// Delete sample
	Delete(context.Context, *TimesheetIdRequest) (*EmptyResponse, error)
	// Create creates a new sample
	CreateEvent(context.Context, *CreateEventRequest) (*Event, error)
	// Update updates an existent consultant
	UpdateEvent(context.Context, *UpdateEventRequest) (*Event, error)
	// Get retrieves a sample by id
	GetEvent(context.Context, *EventIdRequest) (*Event, error)
	// Search searches samples
	SearchEvents(context.Context, *TimesheetIdRequest) (*SearchResponse, error)
	// Delete sample
	DeleteEvent(context.Context, *EventIdRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedTimesheetServiceServer()
}

// UnimplementedTimesheetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimesheetServiceServer struct {
}

func (UnimplementedTimesheetServiceServer) Create(context.Context, *CreateTimesheetRequest) (*Timesheet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTimesheetServiceServer) Update(context.Context, *UpdateTimesheetRequest) (*Timesheet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTimesheetServiceServer) Get(context.Context, *TimesheetIdRequest) (*Timesheet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTimesheetServiceServer) Search(context.Context, *SearchTimesheetRequest) (*SearchTimesheetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedTimesheetServiceServer) Delete(context.Context, *TimesheetIdRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTimesheetServiceServer) CreateEvent(context.Context, *CreateEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedTimesheetServiceServer) UpdateEvent(context.Context, *UpdateEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedTimesheetServiceServer) GetEvent(context.Context, *EventIdRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedTimesheetServiceServer) SearchEvents(context.Context, *TimesheetIdRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchEvents not implemented")
}
func (UnimplementedTimesheetServiceServer) DeleteEvent(context.Context, *EventIdRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedTimesheetServiceServer) mustEmbedUnimplementedTimesheetServiceServer() {}

// UnsafeTimesheetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimesheetServiceServer will
// result in compilation errors.
type UnsafeTimesheetServiceServer interface {
	mustEmbedUnimplementedTimesheetServiceServer()
}

func RegisterTimesheetServiceServer(s grpc.ServiceRegistrar, srv TimesheetServiceServer) {
	s.RegisterService(&TimesheetService_ServiceDesc, srv)
}

func _TimesheetService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTimesheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).Create(ctx, req.(*CreateTimesheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTimesheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).Update(ctx, req.(*UpdateTimesheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimesheetIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).Get(ctx, req.(*TimesheetIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTimesheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).Search(ctx, req.(*SearchTimesheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimesheetIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).Delete(ctx, req.(*TimesheetIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).GetEvent(ctx, req.(*EventIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_SearchEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimesheetIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).SearchEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/SearchEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).SearchEvents(ctx, req.(*TimesheetIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimesheetService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimesheetServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timesheet.TimesheetService/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimesheetServiceServer).DeleteEvent(ctx, req.(*EventIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimesheetService_ServiceDesc is the grpc.ServiceDesc for TimesheetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimesheetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timesheet.TimesheetService",
	HandlerType: (*TimesheetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TimesheetService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TimesheetService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TimesheetService_Get_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _TimesheetService_Search_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TimesheetService_Delete_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _TimesheetService_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _TimesheetService_UpdateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _TimesheetService_GetEvent_Handler,
		},
		{
			MethodName: "SearchEvents",
			Handler:    _TimesheetService_SearchEvents_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _TimesheetService_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timesheet/timesheet.proto",
}
