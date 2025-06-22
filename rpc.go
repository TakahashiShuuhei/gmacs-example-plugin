package main

import (
	"context"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// RPCPlugin はHashiCorp go-pluginのRPCプラグイン実装
type RPCPlugin struct {
	Impl Plugin
}

func (p *RPCPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (p *RPCPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

// RPCServer はプラグイン側のRPCサーバー実装
type RPCServer struct {
	Impl Plugin
}

// RPCClient はホスト側のRPCクライアント実装
type RPCClient struct {
	client *rpc.Client
}

// RPCクライアント側のメソッド実装
func (c *RPCClient) Name() string {
	var resp string
	err := c.client.Call("Plugin.Name", new(interface{}), &resp)
	if err != nil {
		return ""
	}
	return resp
}

func (c *RPCClient) Version() string {
	var resp string
	err := c.client.Call("Plugin.Version", new(interface{}), &resp)
	if err != nil {
		return ""
	}
	return resp
}

func (c *RPCClient) Description() string {
	var resp string
	err := c.client.Call("Plugin.Description", new(interface{}), &resp)
	if err != nil {
		return ""
	}
	return resp
}

func (c *RPCClient) Initialize(ctx context.Context, host HostInterface) error {
	var resp interface{}
	// Note: RPC doesn't easily support context and complex interfaces
	// In real implementation, this would require additional RPC methods
	err := c.client.Call("Plugin.Initialize", new(interface{}), &resp)
	return err
}

func (c *RPCClient) Cleanup() error {
	var resp interface{}
	err := c.client.Call("Plugin.Cleanup", new(interface{}), &resp)
	return err
}

func (c *RPCClient) GetCommands() []CommandSpec {
	var resp []CommandSpec
	err := c.client.Call("Plugin.GetCommands", new(interface{}), &resp)
	if err != nil {
		return nil
	}
	return resp
}

func (c *RPCClient) GetMajorModes() []MajorModeSpec {
	var resp []MajorModeSpec
	err := c.client.Call("Plugin.GetMajorModes", new(interface{}), &resp)
	if err != nil {
		return nil
	}
	return resp
}

func (c *RPCClient) GetMinorModes() []MinorModeSpec {
	var resp []MinorModeSpec
	err := c.client.Call("Plugin.GetMinorModes", new(interface{}), &resp)
	if err != nil {
		return nil
	}
	return resp
}

func (c *RPCClient) GetKeyBindings() []KeyBindingSpec {
	var resp []KeyBindingSpec
	err := c.client.Call("Plugin.GetKeyBindings", new(interface{}), &resp)
	if err != nil {
		return nil
	}
	return resp
}

// RPCサーバー側のメソッド実装
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

func (s *RPCServer) Initialize(args interface{}, resp *interface{}) error {
	// Note: Simplified initialization for RPC
	// In real implementation, this would handle HostInterface properly
	return s.Impl.Initialize(context.Background(), nil)
}

func (s *RPCServer) Cleanup(args interface{}, resp *interface{}) error {
	return s.Impl.Cleanup()
}

func (s *RPCServer) GetCommands(args interface{}, resp *[]CommandSpec) error {
	*resp = s.Impl.GetCommands()
	return nil
}

func (s *RPCServer) GetMajorModes(args interface{}, resp *[]MajorModeSpec) error {
	*resp = s.Impl.GetMajorModes()
	return nil
}

func (s *RPCServer) GetMinorModes(args interface{}, resp *[]MinorModeSpec) error {
	*resp = s.Impl.GetMinorModes()
	return nil
}

func (s *RPCServer) GetKeyBindings(args interface{}, resp *[]KeyBindingSpec) error {
	*resp = s.Impl.GetKeyBindings()
	return nil
}

// GRPCPlugin はHashiCorp go-pluginのGRPCプラグイン実装（将来使用）
type GRPCPlugin struct {
	Impl Plugin
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	// TODO: gRPCサーバー実装（protobuf生成後）
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	// TODO: gRPCクライアント実装（protobuf生成後）
	return nil, nil
}

// プラグインマップ定義
var PluginMap = map[string]plugin.Plugin{
	"gmacs-plugin": &RPCPlugin{},
}

// Handshake はプラグインとホスト間のハンドシェイク設定
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "GMACS_PLUGIN",
	MagicCookieValue: "gmacs-plugin-magic-cookie",
}