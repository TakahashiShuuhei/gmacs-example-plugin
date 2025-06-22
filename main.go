package main

import (
	pluginsdk "github.com/TakahashiShuuhei/gmacs-plugin-sdk"
)

func main() {
	// ServePluginヘルパーを使用してプラグインサーバーを起動
	pluginsdk.ServePlugin(pluginInstance)
}