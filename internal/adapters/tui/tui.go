package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// App represents the TUI application
type App struct {
	program *tea.Program
}

// New creates a new TUI application
func New() *App {
	return &App{}
}

// Run starts the TUI application
func (a *App) Run() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	a.program = p
	_, err := p.Run()
	return err
}

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
	width    int
	height   int
}

func initialModel() model {
	return model{
		choices: []string{
			"List Docker Resources",
			"Analyze Unused Resources",
			"Prune Resources",
			"View History",
			"Settings",
			"Quit",
		},
		selected: make(map[int]struct{}),
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
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			if m.cursor == len(m.choices)-1 {
				return m, tea.Quit
			}
			// Handle other menu selections in the future
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	// Styles
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		Padding(1, 0).
		MarginBottom(1)

	menuItemStyle := lipgloss.NewStyle().
		PaddingLeft(2)

	selectedMenuItemStyle := lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.Color("#7D56F4")).
		Bold(true)

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Padding(1, 0).
		MarginTop(1)

	// Build the view
	var s string

	// Title
	s += titleStyle.Render("🐳 docktidy - Docker Resource Manager")
	s += "\n\n"

	// Menu items
	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "▶ "
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "✓"
		}

		if m.cursor == i {
			s += selectedMenuItemStyle.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, choice))
		} else {
			s += menuItemStyle.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, choice))
		}
		s += "\n"
	}

	// Help text
	s += "\n"
	s += helpStyle.Render("↑/k up • ↓/j down • enter/space select • q/ctrl+c quit")

	return s
}
