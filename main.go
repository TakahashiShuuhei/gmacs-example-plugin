package main

import (
	"net/rpc"
	"github.com/hashicorp/go-plugin"
	pluginsdk "github.com/TakahashiShuuhei/gmacs-plugin-sdk"
)

// RPCServer wraps our plugin implementation to provide RPC methods
type RPCServer struct {
	Impl *ExamplePlugin
}

// ExecuteCommand implements the RPC ExecuteCommand method
func (s *RPCServer) ExecuteCommand(args map[string]interface{}, resp *error) error {
	name, _ := args["name"].(string)
	argsStrings, _ := args["args"].([]string)
	
	// Convert string slice back to []interface{}
	pluginArgs := make([]interface{}, len(argsStrings))
	for i, arg := range argsStrings {
		pluginArgs[i] = arg
	}
	
	*resp = s.Impl.ExecuteCommand(name, pluginArgs...)
	return nil
}

// GetCompletions implements the RPC GetCompletions method
func (s *RPCServer) GetCompletions(args map[string]interface{}, resp *[]string) error {
	command, _ := args["command"].(string)
	prefix, _ := args["prefix"].(string)
	*resp = s.Impl.GetCompletions(command, prefix)
	return nil
}

// Forward other standard plugin methods
func (s *RPCServer) Name(args interface{}, resp *string) error {
	*resp = s.Impl.Name()
	return nil
}

func (s *RPCServer) Version(args interface{}, resp *string) error {
	*resp = s.Impl.Version()
	return nil
}

func (s *RPCServer) Description(args interface{}, resp *string) error {
	*resp = s.Impl.Description()
	return nil
}

func (s *RPCServer) Initialize(args map[string]interface{}, resp *error) error {
	*resp = s.Impl.Initialize(nil, nil) // Simplified for now
	return nil
}

func (s *RPCServer) Cleanup(args interface{}, resp *error) error {
	*resp = s.Impl.Cleanup()
	return nil
}

func (s *RPCServer) GetCommands(args interface{}, resp *[]pluginsdk.CommandSpec) error {
	*resp = s.Impl.GetCommands()
	return nil
}

func (s *RPCServer) GetMajorModes(args interface{}, resp *[]pluginsdk.MajorModeSpec) error {
	*resp = s.Impl.GetMajorModes()
	return nil
}

func (s *RPCServer) GetMinorModes(args interface{}, resp *[]pluginsdk.MinorModeSpec) error {
	*resp = s.Impl.GetMinorModes()
	return nil
}

func (s *RPCServer) GetKeyBindings(args interface{}, resp *[]pluginsdk.KeyBindingSpec) error {
	*resp = s.Impl.GetKeyBindings()
	return nil
}

// CustomRPCPlugin implements go-plugin interface
type CustomRPCPlugin struct{}

func (p *CustomRPCPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: pluginInstance}, nil
}

func (p *CustomRPCPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	// Not used for plugin server
	return nil, nil
}

func main() {
	// Use direct go-plugin RPC implementation
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "GMACS_PLUGIN",
			MagicCookieValue: "gmacs-plugin-magic-cookie",
		},
		Plugins: map[string]plugin.Plugin{
			"gmacs-plugin": &CustomRPCPlugin{},
		},
	})
}