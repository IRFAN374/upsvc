package transport

import (
	"context"
	"fmt"
	"time"

	endpoint "github.com/go-kit/kit/endpoint"
	metrics "github.com/go-kit/kit/metrics"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	opentracinggo "github.com/opentracing/opentracing-go"
)

type EndpointsSet struct {
	RegisterEndpoint       endpoint.Endpoint
	LoginEndpoint          endpoint.Endpoint
	GetUserByIdEndpoint    endpoint.Endpoint
	GetUserByNameEndpoint  endpoint.Endpoint
	GetUserByEmailEndpoint endpoint.Endpoint
	UpdateUserEndpoint     endpoint.Endpoint
	DeleteUserEndpoint     endpoint.Endpoint
	IsExistEndpoint        endpoint.Endpoint
}

func InstrumentingEndpoint(endpoints EndpointsSet, tracer opentracinggo.Tracer) EndpointsSet {
	return EndpointsSet{
		RegisterEndpoint:       opentracing.TraceServer(tracer, "Register")(endpoints.RegisterEndpoint),
		LoginEndpoint:          opentracing.TraceServer(tracer, "Login")(endpoints.LoginEndpoint),
		GetUserByIdEndpoint:    opentracing.TraceServer(tracer, "GetUserById")(endpoints.GetUserByIdEndpoint),
		GetUserByNameEndpoint:  opentracing.TraceServer(tracer, "GetUserByName")(endpoints.GetUserByNameEndpoint),
		GetUserByEmailEndpoint: opentracing.TraceServer(tracer, "GetUserByEmail")(endpoints.GetUserByEmailEndpoint),
		UpdateUserEndpoint:     opentracing.TraceServer(tracer, "UpdateUser")(endpoints.UpdateUserEndpoint),
		DeleteUserEndpoint:     opentracing.TraceServer(tracer, "DeleteUser")(endpoints.DeleteUserEndpoint),
		IsExistEndpoint:        opentracing.TraceServer(tracer, "IsExist")(endpoints.IsExistEndpoint),
	}
}

func LatencyMiddleware(dur metrics.Histogram, methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		dur := dur.With("method", methodName)
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				dur.With("success", fmt.Sprint(err == nil)).Observe(float64(time.Since(begin).Seconds()))
			}(time.Now())

			return next(ctx, request)

		}

	}
}

func RequestFrequencyMiddleware(freq metrics.Gauge, methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		freq := freq.With("method", methodName)
		return func(ctx context.Context, request interface{}) (resposne interface{}, err error) {
			freq.Add(1)
			resposne, err = next(ctx, request)
			freq.Add(-1)
			return resposne, err

		}

	}

}
