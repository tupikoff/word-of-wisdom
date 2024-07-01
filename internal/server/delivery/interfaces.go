package delivery

import (
	"context"

	"github.com/tupikoff/word-of-wisdom/pkg/tcp"
)

type protocolServiceInterface interface {
	Execute(ctx context.Context, connection *tcp.Connection) error
}
