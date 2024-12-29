# Server Monitoring Plugin

This is a plugin designed for monitoring server resources, providing real-time status updates, alerts, and configurable settings through a simple web interface. The plugin includes API endpoints for monitoring resources, configuring settings, and managing notifications. The backend is implemented in Go (Golang), and it integrates with a database for persistent storage.

## Features

- **Server Monitoring**: Track server resource usage like CPU, memory, disk space, and more.
- **Email Notifications**: Get alerted via email when server resources cross critical thresholds.
- **Configurable Settings**: Customize monitoring intervals, thresholds, and notification settings.
- **Web Interface**: A simple HTML-based UI for viewing server status and configuring the plugin.
- **Database Support**: Persistent storage for settings and monitoring data.

## Project Structure

The project is organized into the following directories:

```plaintext
server-monitoring-plugin/
├── main.go                  # Entry point for the plugin
├── handlers/                
│   ├── monitoring.go        # Handler for resource monitoring and use-package gopsutil
│   └── config.go            # Handler for configuration settings
├── db/                      
│   └── db.go                # Database setup and connection
├── utils/                   
│   ├── email.go             # Email notification logic
│   └── monitoring.go        # System monitoring utility
├── templates/               
│   └── index.html           # Frontend UI template
├── go.mod                   # Module dependencies
