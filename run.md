# Run ByteClaw

**Goal**: make running steps clear, simple, and efficient.

## Prerequisites
- Go 1.25+ installed and on PATH
- Basic shell environment (macOS/Linux/WSL)
- API keys for your chosen model/provider if using online models

## Config Setup
- Default config location: `~/.byteclaw/config.json`
- Override with `BYTECLAW_CONFIG=/path/to/config.json`
- Start from example: copy [config/config.example.json](file:///Users/uk/Development/picoclaw/config/config.example.json) to `~/.picoclaw/config.json`, then fill:
  - model_list: your models and keys (model-centric)
  - agents.defaults: model_name, max_tokens, etc.
  - channels: enable and add tokens only for platforms you need

## Build
- Fast build for current platform:

```bash
make build
```

- Binary appears at `build/byteclaw-<platform>-<arch>` and symlink `build/byteclaw`

## Run: Agent (CLI)
- Single message (non-interactive):

```bash
./build/byteclaw agent -m "Hello" -s cli:default
```

- Interactive chat:

```bash
./build/byteclaw agent
```

- Switch model temporarily:

```bash
./build/byteclaw agent --model openai/gpt-5.2
```

- Debug logs:

```bash
./build/byteclaw agent -d
```

## Run: Gateway (Channels + HTTP)
- Starts channels, HTTP health, media store, cron, heartbeat, and the agent loop:

```bash
./build/byteclaw gateway
```

- Debug logs:

```bash
./build/byteclaw gateway -d
```

- Health endpoints:
  - http://HOST:PORT/health
  - http://HOST:PORT/ready
  - HOST/PORT configured in `gateway` section of config

## Docker (optional)
- Minimal image (Alpine):

```bash
make docker-build
docker compose -f docker/docker-compose.yml --profile gateway up
```

- Full-featured image (Node 24, extended MCP tools):

```bash
make docker-build-full
docker compose -f docker/docker-compose.full.yml --profile gateway up
```

## Common Tasks
- Install binary locally:

```bash
make install
```

- Lint and tests:

```bash
make lint
make test
```

- Uninstall:

```bash
make uninstall      # remove binary
make uninstall-all  # remove ~/.picoclaw workspace and data
```

## Troubleshooting
- `go: command not found`
  - Install Go 1.25+, ensure `go version` works and `go` is on PATH
- Gateway starts but no channels
  - Enable a channel in config and add required tokens; check rate limits
- Model errors or token limits
  - Verify `model_list` entries and API keys; reduce `max_tokens`; tune summarization in agents.defaults
- High CPU/memory
  - Run without `-d` to avoid heavy debug logging
  - Limit parallel tool execution via `agents.defaults.max_concurrent_tools`
  - Disable unneeded tools (web/mcp/voice) in `tools` config

## How The Run Works
- Agent loop receives messages, builds context, calls the model, and executes requested tools
- Gateway wires channels to the message bus, exposes HTTP health, and manages media lifecycle
- Concurrency is bounded for tools; logging is efficient unless DEBUG is enabled
- Sessions and summaries reduce memory while preserving relevant context
