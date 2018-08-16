package proto

import "context"

type TranslateEndpointFunc func(ctx context.Context, req *TranslateRequest) (*TranslateResponse, error)

type TranslatorEndpoints struct {
    TranslateEndpoint TranslateEndpointFunc
}

type TranslatorEndpointsWrapper struct {
    translate TranslateEndpointFunc
}

func NewTranslatorEndpointsWrapper(e TranslatorEndpoints) *TranslatorEndpointsWrapper {
    return &TranslatorEndpointsWrapper {
        translate: e.TranslateEndpoint,
    }
}

func (w TranslatorEndpointsWrapper) Translate(ctx context.Context, req *TranslateRequest) (*TranslateResponse, error) {
    return w.translate(ctx, req)
}
