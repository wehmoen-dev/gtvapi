# gtvapi Examples

This directory contains example programs demonstrating how to use the gtvapi library.

## Running the Examples

To run any example, use the `go run` command:

```bash
go run examples/basic.go
go run examples/custom_config.go
```

## Examples

### basic.go

Demonstrates basic usage of the library:
- Getting video information
- Retrieving all tags
- Discovering recent videos
- Checking Twitch live status
- Fetching video comments
- Getting video playlist URLs

### custom_config.go

Shows advanced configuration options:
- Enabling debug logging
- Setting custom timeouts
- Configuring retry behavior
- Using custom HTTP clients
- Setting custom user agents

## Note

These examples require access to the Gronkh.TV API. If you run them, they will make real API requests.
