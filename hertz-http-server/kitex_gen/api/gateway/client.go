// Code generated by Kitex v0.6.1. DO NOT EDIT.

package gateway

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	api "github.com/huangwei021230/api-gateway/hertz-http-server/kitex_gen/api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddNumbers(ctx context.Context, req *api.AdditionRequest, callOptions ...callopt.Option) (r *api.AdditionResponse, err error)
	MultiplyNumbers(ctx context.Context, req *api.MultiplicationRequest, callOptions ...callopt.Option) (r *api.MultiplicationResponse, err error)
	DivideNumbers(ctx context.Context, req *api.DivisionRequest, callOptions ...callopt.Option) (r *api.DivisionResponse, err error)
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
	return &kGatewayClient{
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

type kGatewayClient struct {
	*kClient
}

func (p *kGatewayClient) AddNumbers(ctx context.Context, req *api.AdditionRequest, callOptions ...callopt.Option) (r *api.AdditionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddNumbers(ctx, req)
}

func (p *kGatewayClient) MultiplyNumbers(ctx context.Context, req *api.MultiplicationRequest, callOptions ...callopt.Option) (r *api.MultiplicationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MultiplyNumbers(ctx, req)
}

func (p *kGatewayClient) DivideNumbers(ctx context.Context, req *api.DivisionRequest, callOptions ...callopt.Option) (r *api.DivisionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DivideNumbers(ctx, req)
}
