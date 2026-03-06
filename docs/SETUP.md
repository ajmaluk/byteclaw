# ByteClaw Setup

## Prerequisites
- Go 1.25+
- Git and GitHub account

## Install
```bash
git clone https://github.com/ajmaluk/byteclaw.git
cd byteclaw
make deps
make build
```

## Configure
Create `~/.byteclaw/config.json` or use env vars.

### Providers
Add API keys to `model_list` entries. Example:
```json
{
  "model_list": [
    {
      "model_name": "gpt-5.2",
      "model": "openai/gpt-5.2",
      "api_base": "https://api.openai.com/v1",
      "api_key": "sk-..."
    }
  ]
}
```

### Google OAuth
```bash
export GOOGLE_OAUTH_CLIENT_ID=...
export GOOGLE_OAUTH_CLIENT_SECRET=...
export GOOGLE_OAUTH_PORT=51121
```

### Optional Text Proxy
```bash
export BYTECLAW_TEXT_PROXY_URL=http://localhost:3002
```

## Run
```bash
./byteclaw onboard
./byteclaw agent
```

## Release
- Tag `vX.Y.Z` to trigger GoReleaser.

## Troubleshooting
- Check CI badges and Actions logs.
- Ensure no secrets are committed; use env vars.
