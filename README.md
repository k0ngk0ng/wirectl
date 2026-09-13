# wirectl

Lightweight command host and shared Go CLI library. macOS, Linux, and Windows.

Each repository releases a `wirectl-<command>` executable. Install it on `PATH`
alongside `wirectl`; `wirectl <command> ...` forwards arguments and standard
streams to that executable and preserves its exit status. No plugin registry,
dynamic linking, or shared runtime is required.

The `wire-download` repository supplies `wirectl-download`:

```sh
wirectl download init
wirectl download daemon start
wirectl download 'magnet:?xt=urn:btih:...'
wirectl download watch
```

Build the host with `go build -trimpath -ldflags='-s -w' ./cmd/wirectl`.
The `cli` package provides shared help, routing, plugin dispatch, and exit codes.
Plugins can use it or implement the executable contract in any language.
