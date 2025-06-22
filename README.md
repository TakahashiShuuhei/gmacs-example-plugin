# gmacs Example Plugin

A simple example plugin for gmacs demonstrating basic functionality and serving as a template for plugin development.

## Features

### Commands
- `example-greet` - Display a greeting message from the example plugin
- `example-info` - Show example plugin information  
- `example-insert-timestamp` - Insert current timestamp at cursor position

### Major Mode
- `example-mode` - Example file mode for `.example` and `.ex` files
  - Key binding: `C-c C-e` to trigger greeting

### Minor Mode
- `example-minor-mode` - Example minor mode with helpful features
  - Key binding: `C-c e` to insert timestamp

### Global Key Bindings
- `C-c C-x e` - example-greet command
- `C-c C-x i` - example-info command

## Configuration

Default configuration values:
```json
{
  "greeting_message": "Hello from example plugin!",
  "auto_greet": true,
  "prefix": "[EXAMPLE]"
}
```

## Installation

### From Source
```bash
git clone https://github.com/TakahashiShuuhei/gmacs-example-plugin.git
cd gmacs-example-plugin
go build -o example-plugin
```

### Using gmacs plugin manager (when implemented)
```bash
gmacs plugin install github.com/TakahashiShuuhei/gmacs-example-plugin
```

## Development

This plugin serves as a reference implementation demonstrating:
- Basic plugin structure and interfaces
- RPC communication with gmacs host
- Command registration and handling
- Major/minor mode implementation
- Key binding configuration
- Plugin configuration management

## License

This example plugin is released under the same license as gmacs.