package main

import (
	"github.com/hashicorp/go-plugin"
)

func main() {
	// プラグインのRPCマップにプラグインインスタンスを設定
	PluginMap["gmacs-plugin"] = &RPCPlugin{Impl: pluginInstance}

	// プラグインサーバーを起動
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins:         PluginMap,
	})
}