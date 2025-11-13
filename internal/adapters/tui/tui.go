// Package tui implements the terminal user interface adapter using Bubbletea.
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
	colorDanger    = "#EF476F"
)

// StatusLevel indicates the severity of a status message.
type StatusLevel string

const (
	// StatusLevelUnknown is used when status has not been checked.
	StatusLevelUnknown StatusLevel = "unknown"
	// StatusLevelHealthy indicates everything is working as expected.
	StatusLevelHealthy StatusLevel = "healthy"
	// StatusLevelDegraded indicates an error or degraded state.
	StatusLevelDegraded StatusLevel = "degraded"
)

// StatusMessage represents a message shown in the status area.
type StatusMessage struct {
	Message string
	Level   StatusLevel
}

type appConfig struct {
	dockerStatus StatusMessage
}

// AppOption configures an App instance.
type AppOption func(*appConfig)

// WithDockerStatus overrides the Docker status message shown in the welcome screen.
func WithDockerStatus(status StatusMessage) AppOption {
	return func(cfg *appConfig) {
		cfg.dockerStatus = status
	}
}

func defaultAppConfig() appConfig {
	txt := text.Default()
	return appConfig{
		dockerStatus: StatusMessage{
			Message: txt.Get(text.KeyDockerStatusUnknown),
			Level:   StatusLevelUnknown,
		},
	}
}

// App is the TUI application wrapper for the Bubbletea program.
type App struct {
	program *tea.Program
	config  appConfig
}

// New creates a new TUI application instance.
func New(opts ...AppOption) *App {
	cfg := defaultAppConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	return &App{
		config: cfg,
	}
}

// Run starts the TUI application and blocks until it exits.
func (a *App) Run() error {
	p := tea.NewProgram(initialModel(a.config), tea.WithAltScreen())
	a.program = p
	_, err := p.Run()
	return err
}

type model struct {
	width        int
	height       int
	text         *text.Text
	dockerStatus StatusMessage
}

func initialModel(cfg appConfig) model {
	status := cfg.dockerStatus
	if status.Message == "" {
		status = defaultAppConfig().dockerStatus
	}
	if status.Level == "" {
		status.Level = StatusLevelUnknown
	}

	return model{
		text:         text.Default(),
		dockerStatus: status,
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

	statusStyle := lipgloss.NewStyle().
		Padding(0, 2).
		Faint(true)

	switch m.dockerStatus.Level {
	case StatusLevelHealthy:
		statusStyle = statusStyle.Foreground(lipgloss.Color(colorAccent))
	case StatusLevelDegraded:
		statusStyle = statusStyle.Foreground(lipgloss.Color(colorDanger))
	default:
		statusStyle = statusStyle.Foreground(lipgloss.Color(colorSecondary))
	}

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
	content += "\n"

	renderedStatus := statusStyle.Render(m.dockerStatus.Message)
	if m.width > 0 {
		renderedStatus = lipgloss.PlaceHorizontal(m.width, lipgloss.Right, renderedStatus)
	}
	content += renderedStatus
	content += "\n\n"

	content += helpStyle.Render(m.text.Get(text.KeyHelpQuit))

	return content
}
