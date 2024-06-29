package delivery

import "context"

type protocolServiceInterface interface {
	Execute(ctx context.Context, request string) (response string, err error)
}
