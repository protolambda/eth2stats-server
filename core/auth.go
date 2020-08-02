package core

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	"github.com/ConsenSys/eth2stats-server/store"
)

func (c *Core) authenticate(ctx context.Context) *store.Client {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Error("could not get metadata")
	}

	if token, exists := md["token"]; exists && token[0] != "" {
		return c.store.GetClient(token[0])
	}

	return nil
}

func (c *Core) extractIP(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)

	if realIP, exists := md["x-real-ip"]; exists && realIP[0] != "" {
		return realIP[0]
	}

	p, ok := peer.FromContext(ctx)
	if !ok {
		log.WithField("md", md).Error("could not get IP")
	}
	addrStr := p.Addr.String()
	if strings.Contains(addrStr, ":") {
		addrStr = addrStr[:strings.Index(addrStr, ":")]
	}
	return addrStr
}
