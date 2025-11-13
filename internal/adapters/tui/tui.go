package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thommorais/docktidy/pkg/text"
)

const (
	colorPrimary   = "#7D56F4"
	colorSecondary = "#626262"
	colorAccent    = "#04B575"
	colorText      = "#FAFAFA"
)

// App is the TUI application wrapper for the Bubbletea program.
type App struct {
	program *tea.Program
}

// New creates a new TUI application instance.
func New() *App {
	return &App{}
}

// Run starts the TUI application and blocks until it exits.
func (a *App) Run() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	a.program = p
	_, err := p.Run()
	return err
}

type model struct {
	width  int
	height int
	text   *text.Text
}

func initialModel() model {
	return model{
		text: text.Default(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colorPrimary)).
		Padding(1, 0, 0, 0).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorAccent)).
		Italic(true).
		MarginBottom(2)

	messageStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorText)).
		Padding(0, 2).
		MarginBottom(1)

	featureStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorText)).
		PaddingLeft(4)

	philosophyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorText)).
		Padding(0, 2).
		MarginBottom(2)

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorSecondary)).
		Padding(1, 0, 0, 0)

	var content string

	content += titleStyle.Render(m.text.Get(text.KeyAppTitle))
	content += "\n"

	content += subtitleStyle.Render(m.text.Get(text.KeyAppSubtitle))
	content += "\n\n"

	content += messageStyle.Render(m.text.Get(text.KeyWelcomeMessage))
	content += "\n\n"

	content += featureStyle.Render(fmt.Sprintf("  * %s", m.text.Get(text.KeyWelcomeFeature1)))
	content += "\n"
	content += featureStyle.Render(fmt.Sprintf("  * %s", m.text.Get(text.KeyWelcomeFeature2)))
	content += "\n"
	content += featureStyle.Render(fmt.Sprintf("  * %s", m.text.Get(text.KeyWelcomeFeature3)))
	content += "\n"
	content += featureStyle.Render(fmt.Sprintf("  * %s", m.text.Get(text.KeyWelcomeFeature4)))
	content += "\n\n"

	content += philosophyStyle.Render(m.text.Get(text.KeyWelcomePhilosophy))
	content += "\n\n"

	content += helpStyle.Render(m.text.Get(text.KeyHelpQuit))

	return content
}
