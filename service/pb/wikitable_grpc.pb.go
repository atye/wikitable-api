// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WikiTableJSONAPIClient is the client API for WikiTableJSONAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WikiTableJSONAPIClient interface {
	GetTables(ctx context.Context, in *TablesRequest, opts ...grpc.CallOption) (*TablesResponse, error)
}

type wikiTableJSONAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewWikiTableJSONAPIClient(cc grpc.ClientConnInterface) WikiTableJSONAPIClient {
	return &wikiTableJSONAPIClient{cc}
}

func (c *wikiTableJSONAPIClient) GetTables(ctx context.Context, in *TablesRequest, opts ...grpc.CallOption) (*TablesResponse, error) {
	out := new(TablesResponse)
	err := c.cc.Invoke(ctx, "/wikipedia.WikiTableJSONAPI/GetTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WikiTableJSONAPIServer is the server API for WikiTableJSONAPI service.
// All implementations must embed UnimplementedWikiTableJSONAPIServer
// for forward compatibility
type WikiTableJSONAPIServer interface {
	GetTables(context.Context, *TablesRequest) (*TablesResponse, error)
	mustEmbedUnimplementedWikiTableJSONAPIServer()
}

// UnimplementedWikiTableJSONAPIServer must be embedded to have forward compatible implementations.
type UnimplementedWikiTableJSONAPIServer struct {
}

func (UnimplementedWikiTableJSONAPIServer) GetTables(context.Context, *TablesRequest) (*TablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTables not implemented")
}
func (UnimplementedWikiTableJSONAPIServer) mustEmbedUnimplementedWikiTableJSONAPIServer() {}

// UnsafeWikiTableJSONAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WikiTableJSONAPIServer will
// result in compilation errors.
type UnsafeWikiTableJSONAPIServer interface {
	mustEmbedUnimplementedWikiTableJSONAPIServer()
}

func RegisterWikiTableJSONAPIServer(s *grpc.Server, srv WikiTableJSONAPIServer) {
	s.RegisterService(&_WikiTableJSONAPI_serviceDesc, srv)
}

func _WikiTableJSONAPI_GetTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WikiTableJSONAPIServer).GetTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wikipedia.WikiTableJSONAPI/GetTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WikiTableJSONAPIServer).GetTables(ctx, req.(*TablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WikiTableJSONAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wikipedia.WikiTableJSONAPI",
	HandlerType: (*WikiTableJSONAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTables",
			Handler:    _WikiTableJSONAPI_GetTables_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/pb/wikitable.proto",
}
