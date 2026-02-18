# gRPC Interceptor Sample

This sample demonstrates a custom RoadRunner gRPC interceptor plugin.

The interceptor implementation is in `plugin.go` and exposes:

- `Name() string`
- `UnaryServerInterceptor() grpc.UnaryServerInterceptor`

## Register the plugin in RR

When you build your custom RoadRunner binary, register this plugin in the RR container list (for example, in `container/plugins.go`):

```go
import (
	grpcPlugin "github.com/roadrunner-server/grpc/v5"
	grpcInterceptor "github.com/roadrunner-server/samples/grpc_interceptor"
)

func Plugins() []any {
	return []any{
		// ...
		&grpcInterceptor.Plugin{},
		&grpcPlugin.Plugin{},
		// ...
	}
}
```

## Configure interceptor chain

Use the plugin name returned by `Name()` under `grpc.interceptors`:

```yaml
grpc:
  interceptors:
    - sample-grpc-interceptor
```

The full runnable config example is in `.rr.yaml`.

## Interceptor order

RoadRunner applies configured interceptors in reverse order from `grpc.interceptors`.

Example:

```yaml
grpc:
  interceptors: ["first", "second", "third"]
```

Execution order will be: `third -> second -> first -> handler`.
