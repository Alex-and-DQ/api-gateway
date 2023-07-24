// Code generated by Kitex v0.6.1. DO NOT EDIT.

package divisionmanagement

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	management "github.com/huangwei021230/api-gateway/microservices/division-service/kitex_gen/division/management"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DivideNumbers(ctx context.Context, req *management.DivisionRequest, callOptions ...callopt.Option) (r *management.DivisionResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kDivisionManagementClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kDivisionManagementClient struct {
	*kClient
}

func (p *kDivisionManagementClient) DivideNumbers(ctx context.Context, req *management.DivisionRequest, callOptions ...callopt.Option) (r *management.DivisionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DivideNumbers(ctx, req)
}
