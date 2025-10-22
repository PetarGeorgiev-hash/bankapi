package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

func (server *Server) extractMetada(ctx context.Context) *Metadata {
	meta := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if userAgent := md.Get(grpcGatewayUserAgentHeader); len(userAgent) > 0 {
			meta.UserAgent = userAgent[0]
		}
		if userAgentHeader := md.Get(userAgentHeader); len(userAgentHeader) > 0 {
			meta.UserAgent = userAgentHeader[0]
		}
		if clientIp := md.Get(xForwardedForHeader); len(clientIp) > 0 {
			meta.ClientIP = clientIp[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		meta.ClientIP = p.Addr.String()
	}
	return meta
}
