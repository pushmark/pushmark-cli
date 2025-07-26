# pushmark

A simple CLI tool for sending push notifications.

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap pushmark/pushmark-cli
brew install pushmark
```

### Pre-built binaries

Download the latest release from [GitHub Releases](https://github.com/pushmark/pushmark-cli/releases) for your platform.

### Build from source

```bash
git clone <repository-url>
cd cli-go-pusht
go build -o pushmark
```

## Usage

```bash
pushmark [global options] <channelHash> <message>
```

## Examples

```bash
# Send a simple notification (info - default)
pushmark abc123 "Backup completed successfully"

# Send different notification types
pushmark -t log abc123 "User login recorded"
pushmark -t warning abc123 "Server CPU usage high"
pushmark -t success abc123 "Deployment completed"
pushmark -t error abc123 "Database connection failed"

# Show help
pushmark --help
pushmark -h

# Show version
pushmark --version
pushmark -v
```

## Features

- ✅ Professional command-line interface powered by [urfave/cli](https://github.com/urfave/cli)
- ✅ Colored output (green for success, red for errors)
- ✅ Multiple notification types (`info`, `log`, `warning`, `success`, `error`)
- ✅ Auto-generated help and version information
- ✅ Proper error handling for HTTP and JSON errors
- ✅ 30-second timeout for HTTP requests
- ✅ Short and long flag support (`-t` / `--type`)
- ✅ Comprehensive argument validation
- ✅ Single binary with minimal dependencies

## API

The tool sends a POST request to `https://api.pushmark.app/<channelHash>` with the following JSON payload:

```json
{
  "message": "<message>",
  "type": "<info|log|warning|success|error>"
}
```

## Global Options

- `--type`, `-t`: Notification type (`info`, `log`, `warning`, `success`, `error`) - defaults to `info`
- `--help`, `-h`: Show help information
- `--version`, `-v`: Show version information

## Requirements

- Go 1.21 or later (for building from source)

## Dependencies

- [urfave/cli/v2](https://github.com/urfave/cli) - Professional CLI framework for Go
