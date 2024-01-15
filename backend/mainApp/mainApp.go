package mainApp

import (
	"context"
	_ "embed"
	"fmt"
	"my_test_pro/backend/configApp"
)

// MainApp struct
type MainApp struct {
	ctx context.Context
}

// NewMainApp creates a new App application struct
func NewMainApp() *MainApp {
	return &MainApp{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *MainApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *MainApp) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// 获取name
func (a *MainApp) GetAppName() string {
	return configApp.AppName
}

// 通知
func (a *MainApp) GetNotify(str string) string {
	return str
}
