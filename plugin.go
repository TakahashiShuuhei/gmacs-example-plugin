package main

import (
	"context"
	"fmt"
)

// プラグインインターフェースの定義（gmacs/plugin/interface.goから移植）

// Plugin は全プラグインが実装すべき基本インターフェース
type Plugin interface {
	// プラグイン情報
	Name() string
	Version() string
	Description() string
	
	// ライフサイクル
	Initialize(ctx context.Context, host HostInterface) error
	Cleanup() error
	
	// 機能提供
	GetCommands() []CommandSpec
	GetMajorModes() []MajorModeSpec
	GetMinorModes() []MinorModeSpec
	GetKeyBindings() []KeyBindingSpec
}

// HostInterface はホスト（gmacs）がプラグインに提供するAPI
type HostInterface interface {
	// エディタ操作
	GetCurrentBuffer() BufferInterface
	GetCurrentWindow() WindowInterface
	SetStatus(message string)
	ShowMessage(message string)
	
	// コマンド実行
	ExecuteCommand(name string, args ...interface{}) error
	
	// モード管理
	SetMajorMode(bufferName, modeName string) error
	ToggleMinorMode(bufferName, modeName string) error
	
	// イベント・フック
	AddHook(event string, handler func(...interface{}) error)
	TriggerHook(event string, args ...interface{})
	
	// バッファ操作
	CreateBuffer(name string) BufferInterface
	FindBuffer(name string) BufferInterface
	SwitchToBuffer(name string) error
	
	// ファイル操作
	OpenFile(path string) error
	SaveBuffer(bufferName string) error
	
	// 設定
	GetOption(name string) (interface{}, error)
	SetOption(name string, value interface{}) error
}

// BufferInterface はプラグインからアクセス可能なバッファAPI
type BufferInterface interface {
	Name() string
	Content() string
	SetContent(content string)
	InsertAt(pos int, text string)
	DeleteRange(start, end int)
	CursorPosition() int
	SetCursorPosition(pos int)
	MarkDirty()
	IsDirty() bool
	Filename() string
}

// WindowInterface はプラグインからアクセス可能なウィンドウAPI
type WindowInterface interface {
	Buffer() BufferInterface
	SetBuffer(buffer BufferInterface)
	Width() int
	Height() int
	ScrollOffset() int
	SetScrollOffset(offset int)
}

// CommandSpec はプラグインが提供するコマンド仕様
type CommandSpec struct {
	Name        string
	Description string
	Interactive bool
	Handler     string // プラグイン内のハンドラー名
}

// MajorModeSpec はメジャーモード仕様
type MajorModeSpec struct {
	Name         string
	Extensions   []string // 対象ファイル拡張子
	Description  string
	KeyBindings  []KeyBindingSpec
}

// MinorModeSpec はマイナーモード仕様
type MinorModeSpec struct {
	Name        string
	Description string
	Global      bool // グローバルモードかバッファローカルか
	KeyBindings []KeyBindingSpec
}

// KeyBindingSpec はキーバインディング仕様
type KeyBindingSpec struct {
	Sequence string // "C-c C-g", "M-x" など
	Command  string
	Mode     string // 対象モード（空の場合はグローバル）
}

// ExamplePlugin はサンプルプラグインの実装
type ExamplePlugin struct {
	host HostInterface
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
func (p *ExamplePlugin) Initialize(ctx context.Context, host HostInterface) error {
	p.host = host
	p.config = make(map[string]interface{})
	
	// デフォルト設定
	p.config["greeting_message"] = "Hello from example plugin!"
	p.config["auto_greet"] = true
	p.config["prefix"] = "[EXAMPLE]"
	
	// プラグイン初期化完了メッセージ
	if p.config["auto_greet"].(bool) {
		message := fmt.Sprintf("%s %s", p.config["prefix"], p.config["greeting_message"])
		host.ShowMessage(message)
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
func (p *ExamplePlugin) GetCommands() []CommandSpec {
	return []CommandSpec{
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
func (p *ExamplePlugin) GetMajorModes() []MajorModeSpec {
	return []MajorModeSpec{
		{
			Name:        "example-mode",
			Extensions:  []string{".example", ".ex"},
			Description: "Example file mode with basic syntax highlighting",
			KeyBindings: []KeyBindingSpec{
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
func (p *ExamplePlugin) GetMinorModes() []MinorModeSpec {
	return []MinorModeSpec{
		{
			Name:        "example-minor-mode",
			Description: "Example minor mode that adds helpful features",
			Global:      false, // バッファローカル
			KeyBindings: []KeyBindingSpec{
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
func (p *ExamplePlugin) GetKeyBindings() []KeyBindingSpec {
	return []KeyBindingSpec{
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

// プラグイン設定の共通化された実装
var pluginInstance = &ExamplePlugin{}