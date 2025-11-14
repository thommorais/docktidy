package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thommorais/docktidy/internal/domain"
	"github.com/thommorais/docktidy/pkg/text"
)

func TestNew(t *testing.T) {
	app := New()
	if app == nil {
		t.Fatal("New() returned nil")
	}

	if app.program != nil {
		t.Error("New() should not initialize program until Run() is called")
	}

	if app.config.dockerStatus.Message == "" {
		t.Error("New() config has empty docker status message")
	}
}

func TestInitialModel(t *testing.T) {
	m := initialModel(defaultAppConfig())

	if m.width != 0 {
		t.Errorf("initialModel() width = %d, want 0", m.width)
	}

	if m.height != 0 {
		t.Errorf("initialModel() height = %d, want 0", m.height)
	}

	if m.text == nil {
		t.Fatal("initialModel() text is nil")
	}

	if m.text.GetLocale() != text.LocaleEN {
		t.Errorf("initialModel() text locale = %v, want %v", m.text.GetLocale(), text.LocaleEN)
	}

	if m.dockerStatus.Message == "" {
		t.Fatal("initialModel() docker status message is empty")
	}
}

func TestModelInit(t *testing.T) {
	m := initialModel(defaultAppConfig())
	cmd := m.Init()

	if cmd != nil {
		t.Errorf("Init() returned non-nil cmd, want nil")
	}
}

func TestModelUpdate_WindowSizeMsg(t *testing.T) {
	m := initialModel(defaultAppConfig())

	msg := tea.WindowSizeMsg{
		Width:  100,
		Height: 50,
	}

	updatedModel, cmd := m.Update(msg)

	if cmd != nil {
		t.Errorf("Update(WindowSizeMsg) returned non-nil cmd")
	}

	updated := updatedModel.(model)
	if updated.width != 100 {
		t.Errorf("Update(WindowSizeMsg) width = %d, want 100", updated.width)
	}

	if updated.height != 50 {
		t.Errorf("Update(WindowSizeMsg) height = %d, want 50", updated.height)
	}
}

func TestModelUpdate_QuitKeys(t *testing.T) {
	tests := []struct {
		name   string
		keyMsg tea.KeyMsg
	}{
		{
			name: "q key",
			keyMsg: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune{'q'},
			},
		},
		{
			name: "esc key",
			keyMsg: tea.KeyMsg{
				Type: tea.KeyEsc,
			},
		},
		{
			name: "ctrl+c",
			keyMsg: tea.KeyMsg{
				Type: tea.KeyCtrlC,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := initialModel(defaultAppConfig())

			_, cmd := m.Update(tt.keyMsg)

			if cmd == nil {
				t.Errorf("Update(%s) returned nil cmd, want tea.Quit", tt.name)
			}
		})
	}
}

func TestModelView(t *testing.T) {
	m := initialModel(defaultAppConfig())
	view := m.View()

	if view == "" {
		t.Error("View() returned empty string")
	}

	expectedContent := []string{
		"docktidy - Spark joy in your Docker environment",
		"Interactive resource selection with risk levels",
		"Usage history tracking to protect active resources",
		"Dry-run mode to preview changes before applying",
		"Detailed cleanup history and recovery commands",
		"Press Enter to continue",
		"Press 'q', 'esc', or ctrl+c to quit",
	}

	for _, content := range expectedContent {
		if !strings.Contains(view, content) {
			t.Errorf("View() missing expected content: %q", content)
		}
	}

	if !strings.Contains(view, m.dockerStatus.Message) {
		t.Errorf("View() missing docker status message %q", m.dockerStatus.Message)
	}
}

func TestModelView_UsesTextPackage(t *testing.T) {
	m := initialModel(defaultAppConfig())

	title := m.text.Get(text.KeyAppTagline)
	if title == "" {
		t.Error("text.Get(KeyAppTagline) returned empty string")
	}

	view := m.View()
	if !strings.Contains(view, title) {
		t.Errorf("View() does not contain tagline from text package: %q", title)
	}
}

func TestModelView_AfterResize(t *testing.T) {
	m := initialModel(defaultAppConfig())

	msg := tea.WindowSizeMsg{
		Width:  120,
		Height: 60,
	}

	updatedModel, _ := m.Update(msg)
	updated := updatedModel.(model)

	view := updated.View()
	if view == "" {
		t.Error("View() after resize returned empty string")
	}

	if !strings.Contains(view, "docktidy") {
		t.Error("View() after resize missing app name")
	}
}

func TestWithDockerStatusOption(t *testing.T) {
	status := StatusMessage{
		Message: "Docker: Connected",
		Level:   StatusLevelHealthy,
	}

	app := New(WithDockerStatus(status))

	m := initialModel(app.config)
	if m.dockerStatus.Message != status.Message {
		t.Fatalf("dockerStatus.Message = %q, want %q", m.dockerStatus.Message, status.Message)
	}

	if m.dockerStatus.Level != status.Level {
		t.Fatalf("dockerStatus.Level = %v, want %v", m.dockerStatus.Level, status.Level)
	}
}

func TestWithDiskUsageOption(t *testing.T) {
	usage := domain.DiskUsage{
		Rows: []domain.DiskUsageRow{
			{Type: "Images", Total: 2},
		},
	}

	app := New(WithDiskUsage(usage))
	m := initialModel(app.config)

	if len(m.usage.Rows) != 1 {
		t.Fatalf("usage rows = %d, want 1", len(m.usage.Rows))
	}
}

func TestModelNavigationToDashboard(t *testing.T) {
	usage := domain.DiskUsage{
		Rows: []domain.DiskUsageRow{
			{
				Type:             "Images",
				Total:            2,
				Active:           1,
				SizeBytes:        2048,
				ReclaimableBytes: 512,
			},
		},
	}

	app := New(WithDiskUsage(usage))
	m := initialModel(app.config)

	enterMsg := tea.KeyMsg{Type: tea.KeyEnter}
	dashModel, _ := m.Update(enterMsg)
	dash := dashModel.(model)

	if dash.screen != screenDashboard {
		t.Fatalf("screen = %v, want dashboard", dash.screen)
	}

	view := dash.View()
	if !strings.Contains(view, "Docker Disk Usage") {
		t.Fatalf("dashboard view missing title")
	}
	if !strings.Contains(view, "Images") {
		t.Fatalf("dashboard view missing usage row")
	}
	if !strings.Contains(view, "Press 'b' to return to the welcome menu") {
		t.Fatalf("dashboard footer missing back instruction")
	}

	backMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'b'}}
	welcomeModel, _ := dash.Update(backMsg)
	welcome := welcomeModel.(model)
	if welcome.screen != screenWelcome {
		t.Fatalf("screen after back = %v, want welcome", welcome.screen)
	}
}

func TestRenderUsageTable(t *testing.T) {
	rows := []domain.DiskUsageRow{
		{Type: "Images", Total: 2, Active: 1, SizeBytes: 1024, ReclaimableBytes: 256},
		{Type: "Containers", Total: 3, Active: 2, SizeBytes: 2048, ReclaimableBytes: 1024},
	}

	table := renderUsageTable(rows)
	if !strings.Contains(table, "TYPE") {
		t.Fatalf("table missing header TYPE")
	}
	if !strings.Contains(table, "Images") || !strings.Contains(table, "Containers") {
		t.Fatalf("table missing row labels")
	}
	if !strings.Contains(table, "1.0 kB") {
		t.Fatalf("table missing formatted size")
	}
}

func TestColorConstants(t *testing.T) {
	colors := map[string]string{
		"colorPrimary":   colorPrimary,
		"colorSecondary": colorSecondary,
		"colorAccent":    colorAccent,
		"colorText":      colorText,
	}

	for name, color := range colors {
		if color == "" {
			t.Errorf("%s is empty", name)
		}

		if !strings.HasPrefix(color, "#") {
			t.Errorf("%s = %q, want hex color starting with #", name, color)
		}
	}
}
