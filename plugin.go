package main

import (
	"context"
	"fmt"
	"time"

	pluginsdk "github.com/TakahashiShuuhei/gmacs-plugin-sdk"
)

// ExamplePlugin はサンプルプラグインの実装
type ExamplePlugin struct {
	host   pluginsdk.HostInterface
	config map[string]interface{}
}

// Name implements Plugin interface
func (p *ExamplePlugin) Name() string {
	return "example-plugin"
}

// Version implements Plugin interface
func (p *ExamplePlugin) Version() string {
	return "1.0.0"
}

// Description implements Plugin interface
func (p *ExamplePlugin) Description() string {
	return "A simple example plugin for gmacs demonstrating basic functionality"
}

// Initialize implements Plugin interface
func (p *ExamplePlugin) Initialize(ctx context.Context, host pluginsdk.HostInterface) error {
	p.host = host
	p.config = make(map[string]interface{})

	// デフォルト設定
	p.config["greeting_message"] = "Hello from example plugin!"
	p.config["auto_greet"] = true
	p.config["prefix"] = "[EXAMPLE]"

	// プラグイン初期化完了メッセージ
	if p.config["auto_greet"].(bool) {
		message := fmt.Sprintf("%s %s", p.config["prefix"], p.config["greeting_message"])
		if host != nil {
			host.ShowMessage(message)
		}
	}

	return nil
}

// Cleanup implements Plugin interface
func (p *ExamplePlugin) Cleanup() error {
	if p.host != nil {
		p.host.ShowMessage("[EXAMPLE] Plugin shutting down. Goodbye!")
	}
	return nil
}

// GetCommands implements Plugin interface
func (p *ExamplePlugin) GetCommands() []pluginsdk.CommandSpec {
	return []pluginsdk.CommandSpec{
		{
			Name:        "example-greet",
			Description: "Display a greeting message from the example plugin",
			Interactive: true,
			Handler:     "HandleGreet",
		},
		{
			Name:        "example-info",
			Description: "Show example plugin information",
			Interactive: true,
			Handler:     "HandleInfo",
		},
		{
			Name:        "example-insert-timestamp",
			Description: "Insert current timestamp at cursor position",
			Interactive: true,
			Handler:     "HandleInsertTimestamp",
		},
	}
}

// GetMajorModes implements Plugin interface
func (p *ExamplePlugin) GetMajorModes() []pluginsdk.MajorModeSpec {
	return []pluginsdk.MajorModeSpec{
		{
			Name:        "example-mode",
			Extensions:  []string{".example", ".ex"},
			Description: "Example file mode with basic syntax highlighting",
			KeyBindings: []pluginsdk.KeyBindingSpec{
				{
					Sequence: "C-c C-e",
					Command:  "example-greet",
					Mode:     "example-mode",
				},
			},
		},
	}
}

// GetMinorModes implements Plugin interface
func (p *ExamplePlugin) GetMinorModes() []pluginsdk.MinorModeSpec {
	return []pluginsdk.MinorModeSpec{
		{
			Name:        "example-minor-mode",
			Description: "Example minor mode that adds helpful features",
			Global:      false, // バッファローカル
			KeyBindings: []pluginsdk.KeyBindingSpec{
				{
					Sequence: "C-c e",
					Command:  "example-insert-timestamp",
					Mode:     "example-minor-mode",
				},
			},
		},
	}
}

// GetKeyBindings implements Plugin interface
func (p *ExamplePlugin) GetKeyBindings() []pluginsdk.KeyBindingSpec {
	return []pluginsdk.KeyBindingSpec{
		{
			Sequence: "C-c C-x e",
			Command:  "example-greet",
			Mode:     "", // グローバル
		},
		{
			Sequence: "C-c C-x i",
			Command:  "example-info",
			Mode:     "", // グローバル
		},
	}
}

// プラグインコマンドハンドラー（デモ用）
func (p *ExamplePlugin) HandleGreet() error {
	if p.host != nil {
		message := fmt.Sprintf("%s %s", p.config["prefix"], p.config["greeting_message"])
		p.host.ShowMessage(message)
	}
	return nil
}

func (p *ExamplePlugin) HandleInfo() error {
	if p.host != nil {
		info := fmt.Sprintf("[EXAMPLE] %s v%s - %s", p.Name(), p.Version(), p.Description())
		p.host.ShowMessage(info)
	}
	return nil
}

func (p *ExamplePlugin) HandleInsertTimestamp() error {
	if p.host != nil {
		buffer := p.host.GetCurrentBuffer()
		if buffer != nil {
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			pos := buffer.CursorPosition()
			buffer.InsertAt(pos, timestamp)
		}
	}
	return nil
}

// プラグインインスタンス
var pluginInstance = &ExamplePlugin{}