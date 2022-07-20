package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/ledgermgr/general"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type TemplateServer struct {
	template.UnimplementedGeneralServer
}

func Register(server grpc.ServiceRegistrar) {
	template.RegisterGeneralServer(server, &TemplateServer{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return template.RegisterGeneralHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
