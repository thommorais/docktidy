package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
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
}

func TestInitialModel(t *testing.T) {
	m := initialModel()

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
}

func TestModelInit(t *testing.T) {
	m := initialModel()
	cmd := m.Init()

	if cmd != nil {
		t.Errorf("Init() returned non-nil cmd, want nil")
	}
}

func TestModelUpdate_WindowSizeMsg(t *testing.T) {
	m := initialModel()

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
			m := initialModel()

			_, cmd := m.Update(tt.keyMsg)

			if cmd == nil {
				t.Errorf("Update(%s) returned nil cmd, want tea.Quit", tt.name)
			}
		})
	}
}

func TestModelView(t *testing.T) {
	m := initialModel()
	view := m.View()

	if view == "" {
		t.Error("View() returned empty string")
	}

	expectedContent := []string{
		"docktidy - Docker Resource Manager",
		"Spark joy in your Docker environment",
		"Welcome to docktidy",
		"Interactive resource selection with risk levels",
		"Usage history tracking to protect active resources",
		"Dry-run mode to preview changes before applying",
		"Detailed cleanup history and recovery commands",
		"spark joy",
		"Press 'q', 'esc', or ctrl+c to quit",
	}

	for _, content := range expectedContent {
		if !strings.Contains(view, content) {
			t.Errorf("View() missing expected content: %q", content)
		}
	}
}

func TestModelView_UsesTextPackage(t *testing.T) {
	m := initialModel()

	title := m.text.Get(text.KeyAppTitle)
	if title == "" {
		t.Error("text.Get(KeyAppTitle) returned empty string")
	}

	view := m.View()
	if !strings.Contains(view, title) {
		t.Errorf("View() does not contain title from text package: %q", title)
	}
}

func TestModelView_AfterResize(t *testing.T) {
	m := initialModel()

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
