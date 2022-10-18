// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: identity/agent.proto

package identity

import (
	context "context"
	common "github.com/hyperledger-labs/weaver-dlt-interoperability/common/protos-go/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IINAgentClient is the client API for IINAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IINAgentClient interface {
	// user or agent triggers a sync of external/foreign network unit's state
	SyncExternalState(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error)
	// Requesting network unit's state from a foreign IIN agent.
	RequestIdentityConfiguration(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error)
	// Handling network unit's state sent by a foreign IIN agent.
	SendIdentityConfiguration(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error)
	// user or agent triggers a flow to collect signatures attesting an
	// external/foreign network unit's state and recording to ledger
	FlowAndRecordAttestations(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error)
	// Requesting attestation from a local IIN agent.
	RequestAttestation(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error)
	// Handling attestation sent by a local IIN agent.
	SendAttestation(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error)
}

type iINAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewIINAgentClient(cc grpc.ClientConnInterface) IINAgentClient {
	return &iINAgentClient{cc}
}

func (c *iINAgentClient) SyncExternalState(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, "/identity.agent.IINAgent/SyncExternalState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iINAgentClient) RequestIdentityConfiguration(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, "/identity.agent.IINAgent/RequestIdentityConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iINAgentClient) SendIdentityConfiguration(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, "/identity.agent.IINAgent/SendIdentityConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iINAgentClient) FlowAndRecordAttestations(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, "/identity.agent.IINAgent/FlowAndRecordAttestations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iINAgentClient) RequestAttestation(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, "/identity.agent.IINAgent/RequestAttestation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iINAgentClient) SendAttestation(ctx context.Context, in *NetworkUnitIdentity, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, "/identity.agent.IINAgent/SendAttestation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IINAgentServer is the server API for IINAgent service.
// All implementations must embed UnimplementedIINAgentServer
// for forward compatibility
type IINAgentServer interface {
	// user or agent triggers a sync of external/foreign network unit's state
	SyncExternalState(context.Context, *NetworkUnitIdentity) (*common.Ack, error)
	// Requesting network unit's state from a foreign IIN agent.
	RequestIdentityConfiguration(context.Context, *NetworkUnitIdentity) (*common.Ack, error)
	// Handling network unit's state sent by a foreign IIN agent.
	SendIdentityConfiguration(context.Context, *NetworkUnitIdentity) (*common.Ack, error)
	// user or agent triggers a flow to collect signatures attesting an
	// external/foreign network unit's state and recording to ledger
	FlowAndRecordAttestations(context.Context, *NetworkUnitIdentity) (*common.Ack, error)
	// Requesting attestation from a local IIN agent.
	RequestAttestation(context.Context, *NetworkUnitIdentity) (*common.Ack, error)
	// Handling attestation sent by a local IIN agent.
	SendAttestation(context.Context, *NetworkUnitIdentity) (*common.Ack, error)
	mustEmbedUnimplementedIINAgentServer()
}

// UnimplementedIINAgentServer must be embedded to have forward compatible implementations.
type UnimplementedIINAgentServer struct {
}

func (UnimplementedIINAgentServer) SyncExternalState(context.Context, *NetworkUnitIdentity) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncExternalState not implemented")
}
func (UnimplementedIINAgentServer) RequestIdentityConfiguration(context.Context, *NetworkUnitIdentity) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestIdentityConfiguration not implemented")
}
func (UnimplementedIINAgentServer) SendIdentityConfiguration(context.Context, *NetworkUnitIdentity) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendIdentityConfiguration not implemented")
}
func (UnimplementedIINAgentServer) FlowAndRecordAttestations(context.Context, *NetworkUnitIdentity) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlowAndRecordAttestations not implemented")
}
func (UnimplementedIINAgentServer) RequestAttestation(context.Context, *NetworkUnitIdentity) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAttestation not implemented")
}
func (UnimplementedIINAgentServer) SendAttestation(context.Context, *NetworkUnitIdentity) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAttestation not implemented")
}
func (UnimplementedIINAgentServer) mustEmbedUnimplementedIINAgentServer() {}

// UnsafeIINAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IINAgentServer will
// result in compilation errors.
type UnsafeIINAgentServer interface {
	mustEmbedUnimplementedIINAgentServer()
}

func RegisterIINAgentServer(s grpc.ServiceRegistrar, srv IINAgentServer) {
	s.RegisterService(&IINAgent_ServiceDesc, srv)
}

func _IINAgent_SyncExternalState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkUnitIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IINAgentServer).SyncExternalState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.agent.IINAgent/SyncExternalState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IINAgentServer).SyncExternalState(ctx, req.(*NetworkUnitIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _IINAgent_RequestIdentityConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkUnitIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IINAgentServer).RequestIdentityConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.agent.IINAgent/RequestIdentityConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IINAgentServer).RequestIdentityConfiguration(ctx, req.(*NetworkUnitIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _IINAgent_SendIdentityConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkUnitIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IINAgentServer).SendIdentityConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.agent.IINAgent/SendIdentityConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IINAgentServer).SendIdentityConfiguration(ctx, req.(*NetworkUnitIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _IINAgent_FlowAndRecordAttestations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkUnitIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IINAgentServer).FlowAndRecordAttestations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.agent.IINAgent/FlowAndRecordAttestations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IINAgentServer).FlowAndRecordAttestations(ctx, req.(*NetworkUnitIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _IINAgent_RequestAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkUnitIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IINAgentServer).RequestAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.agent.IINAgent/RequestAttestation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IINAgentServer).RequestAttestation(ctx, req.(*NetworkUnitIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _IINAgent_SendAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkUnitIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IINAgentServer).SendAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.agent.IINAgent/SendAttestation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IINAgentServer).SendAttestation(ctx, req.(*NetworkUnitIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

// IINAgent_ServiceDesc is the grpc.ServiceDesc for IINAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IINAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity.agent.IINAgent",
	HandlerType: (*IINAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncExternalState",
			Handler:    _IINAgent_SyncExternalState_Handler,
		},
		{
			MethodName: "RequestIdentityConfiguration",
			Handler:    _IINAgent_RequestIdentityConfiguration_Handler,
		},
		{
			MethodName: "SendIdentityConfiguration",
			Handler:    _IINAgent_SendIdentityConfiguration_Handler,
		},
		{
			MethodName: "FlowAndRecordAttestations",
			Handler:    _IINAgent_FlowAndRecordAttestations_Handler,
		},
		{
			MethodName: "RequestAttestation",
			Handler:    _IINAgent_RequestAttestation_Handler,
		},
		{
			MethodName: "SendAttestation",
			Handler:    _IINAgent_SendAttestation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/agent.proto",
}
