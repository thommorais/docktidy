// Package tui implements the terminal user interface adapter using Bubbletea.
package tui

import (
	"fmt"
	"strconv"
	"strings"

	units "github.com/docker/go-units"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thommorais/docktidy/internal/domain"
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

type screenState int

const (
	screenWelcome screenState = iota
	screenDashboard
)

type appConfig struct {
	dockerStatus StatusMessage
	usage        domain.DiskUsage
}

// AppOption configures an App instance.
type AppOption func(*appConfig)

// WithDockerStatus overrides the Docker status message shown in the welcome screen.
func WithDockerStatus(status StatusMessage) AppOption {
	return func(cfg *appConfig) {
		cfg.dockerStatus = status
	}
}

// WithDiskUsage injects disk usage data for the dashboard view.
func WithDiskUsage(usage domain.DiskUsage) AppOption {
	return func(cfg *appConfig) {
		cfg.usage = usage
	}
}

func defaultAppConfig() appConfig {
	txt := text.Default()
	return appConfig{
		dockerStatus: StatusMessage{
			Message: txt.Get(text.KeyDockerStatusUnknown),
			Level:   StatusLevelUnknown,
		},
		usage: domain.DiskUsage{},
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
	usage        domain.DiskUsage
	screen       screenState
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
		usage:        cfg.usage,
		screen:       screenWelcome,
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
		case "enter":
			if m.screen == screenWelcome {
				m.screen = screenDashboard
			}
		case "b":
			if m.screen == screenDashboard {
				m.screen = screenWelcome
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colorPrimary)).
		MarginBottom(2)

	messageStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorText)).
		Padding(0, 2).
		MarginBottom(1)

	featureStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorText)).
		PaddingLeft(4)

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

	var builder strings.Builder

	builder.WriteString(titleStyle.Render(m.text.Get(text.KeyAppTagline)))
	builder.WriteString("\n\n")

	switch m.screen {
	case screenDashboard:
		builder.WriteString(m.renderDashboard(messageStyle))
	default:
		builder.WriteString(m.renderWelcome(messageStyle, featureStyle))
	}

	builder.WriteString("\n")
	builder.WriteString(m.renderStatusLine(statusStyle))
	builder.WriteString("\n\n")

	builder.WriteString(helpStyle.Render(m.footerMessage()))

	return builder.String()
}

func (m model) renderWelcome(messageStyle lipgloss.Style, featureStyle lipgloss.Style) string {
	var b strings.Builder

	b.WriteString(messageStyle.Render(m.text.Get(text.KeyWelcomeMessage)))
	b.WriteString("\n\n")

	features := []string{
		m.text.Get(text.KeyWelcomeFeature1),
		m.text.Get(text.KeyWelcomeFeature2),
		m.text.Get(text.KeyWelcomeFeature3),
		m.text.Get(text.KeyWelcomeFeature4),
	}

	for _, feature := range features {
		b.WriteString(featureStyle.Render(fmt.Sprintf("  * %s", feature)))
		b.WriteString("\n")
	}

	return strings.TrimRight(b.String(), "\n")
}

func (m model) renderDashboard(messageStyle lipgloss.Style) string {
	var b strings.Builder

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colorAccent)).
		Padding(0, 2).
		MarginBottom(1)

	b.WriteString(headerStyle.Render(m.text.Get(text.KeyDashboardTitle)))
	b.WriteString("\n")

	if len(m.usage.Rows) == 0 {
		b.WriteString(messageStyle.Render(m.text.Get(text.KeyDashboardEmpty)))
		return b.String()
	}

	table := renderUsageTable(m.usage.Rows)
	b.WriteString(table)

	return strings.TrimRight(b.String(), "\n")
}

func (m model) renderStatusLine(style lipgloss.Style) string {
	rendered := style.Render(m.dockerStatus.Message)
	if m.width > 0 {
		return lipgloss.PlaceHorizontal(m.width, lipgloss.Right, rendered)
	}
	return rendered
}

func (m model) footerMessage() string {
	var parts []string
	switch m.screen {
	case screenDashboard:
		parts = append(parts, m.text.Get(text.KeyDashboardBack))
	default:
		parts = append(parts, m.text.Get(text.KeyWelcomeContinue))
	}
	parts = append(parts, m.text.Get(text.KeyHelpQuit))
	return strings.Join(parts, " • ")
}

func renderUsageTable(rows []domain.DiskUsageRow) string {
	headers := []string{"TYPE", "TOTAL", "ACTIVE", "SIZE", "RECLAIMABLE"}
	table := make([][]string, 0, len(rows)+1)
	table = append(table, headers)

	for _, row := range rows {
		table = append(table, []string{
			row.Type,
			strconv.Itoa(row.Total),
			strconv.Itoa(row.Active),
			formatBytes(row.SizeBytes),
			formatBytes(row.ReclaimableBytes),
		})
	}

	if len(table) == 1 {
		return ""
	}

	widths := make([]int, len(headers))
	for _, row := range table {
		for i, col := range row {
			if len(col) > widths[i] {
				widths[i] = len(col)
			}
		}
	}

	totalWidth := 0
	for _, w := range widths {
		totalWidth += w
	}
	totalWidth += (len(widths) - 1) * 3

	var b strings.Builder
	for i, row := range table {
		for j, col := range row {
			fmt.Fprintf(&b, "%-*s", widths[j], col)
			if j < len(row)-1 {
				b.WriteString("   ")
			}
		}
		b.WriteString("\n")
		if i == 0 {
			b.WriteString(strings.Repeat("-", totalWidth))
			b.WriteString("\n")
		}
	}

	return b.String()
}

func formatBytes(size int64) string {
	if size <= 0 {
		return "0 B"
	}
	return units.HumanSizeWithPrecision(float64(size), 1)
}
