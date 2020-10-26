package rpc_client

import "context"

type GenericService struct {
	Invoke func(ctx context.Context, methodName string, args map[string]interface{}) (interface{}, error) `rpc:"$invoke"`
}
