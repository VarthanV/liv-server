# liv-server
---
A lightweight development server built with Go and Gin that serves files with live reload functionality via WebSocket connections.

## Features

- **File Server**: Serves static files and directories from the current working directory
- **Directory Listing**: Automatically generates HTML listings for directories
- **Live Reload**: Automatically reloads HTML pages when files change using WebSocket connections
- **File Watching**: Monitors the entire project directory for changes using `fsnotify`
- **HTML Enhancement**: Automatically injects live reload scripts into HTML files
- **Dual Port Architecture**: Separates file serving (port 8060) from WebSocket connections (port 8070)

## Architecture

The project follows a clean architecture with separate concerns:

- **Controllers**: Handle HTTP routing and WebSocket connections
- **File Service**: Manages file operations, directory listing, and file watching
- **Templates**: HTML templates for directory listings

## Project Structure

```
.
├── .gitignore
├── Makefile
├── README.md
├── controllers/
│   ├── controller.go      # Main controller with routing setup
│   ├── objects.go         # Error definitions
│   ├── serve_file.go      # File serving logic
│   └── socket.go          # WebSocket handling
├── fileservice/
│   ├── fileservice.go     # File operations and watching
│   ├── main.go           # Service constructor
│   └── objects.go        # Data structures
├── go.mod
├── go.sum
├── main.go               # Application entry point
└── templates/
    └── list_files.html   # Directory listing template
```

## Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd liv-server
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

## Usage

### Using Make (Recommended)

```bash
# Build and run
make run

# Or build only
make build

# Clean build artifacts
make clean

# Run tests
make test
```

### Manual Build

```bash
# Build the project
go build -o liv-server

# Run the server
./liv-server
```

## Server Endpoints

- **File Server**: `http://127.0.0.1:8060/*` - Serves files and directories
- **WebSocket**: `ws://127.0.0.1:8070/ws` - Live reload connection

### Examples

- View directory listing: `http://127.0.0.1:8060/`
- Serve a file: `http://127.0.0.1:8060/path/to/file.txt`
- Serve HTML with live reload: `http://127.0.0.1:8060/index.html`

## How Live Reload Works

1. When serving HTML files, the server automatically injects a WebSocket client script
2. The client connects to the WebSocket server on port 8070
3. The file service watches all directories recursively for changes
4. When a file is modified, a "reload" message is sent to all connected clients
5. Clients automatically refresh the page upon receiving the reload signal

## Configuration

The server configuration is defined in `controllers/controller.go`:

- **File Server Host**: `127.0.0.1`
- **File Server Port**: `8060`
- **WebSocket Port**: `8070`

## Dependencies

- **[Gin](https://github.com/gin-gonic/gin)**: HTTP web framework
- **[Gorilla WebSocket](https://github.com/gorilla/websocket)**: WebSocket implementation
- **[fsnotify](https://github.com/fsnotify/fsnotify)**: File system notifications
- **[Logrus](https://github.com/sirupsen/logrus)**: Structured logging

## Development

The server is designed for development workflows where you need:

- Quick file serving from any directory
- Automatic browser refresh when files change
- Directory browsing capabilities
- Minimal setup and configuration

## File Serving Behavior

- **Directories**: Shows an HTML listing of contents (top-level only)
- **HTML Files**: Served with live reload script injection
- **Other Files**: Served as static content
- **Missing Files**: Returns appropriate error responses

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
