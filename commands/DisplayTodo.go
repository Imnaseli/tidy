package commands

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sijirama/tidy/database"
	"os"
	"time"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type tmodel struct {
	table table.Model
}

func (m tmodel) Init() tea.Cmd { return nil }

func (m tmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tmodel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func DisplayTodos(todos []database.TodoComplete) {
	columns := []table.Column{
		{Title: "Id", Width: 2},
		{Title: "Title", Width: 20},
		{Title: "Description", Width: 30},
		{Title: "Status", Width: 6},
		{Title: "Created at", Width: 20},
	}

	rows := mapTodosToRows(todos)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := tmodel{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func mapTodosToRows(todos []database.TodoComplete) []table.Row {
	var rows []table.Row
	for _, todo := range todos {
		completed := fmt.Sprintf("%s", formatCompleted(todo.Completed))
		createdAt := todo.CreatedAt.Format(time.RFC822)
		rows = append(rows, table.Row{
			fmt.Sprint(todo.ID),
			todo.Title,
			todo.Description,
			completed,
			createdAt,
		})
	}
	return rows
}

func formatCompleted(completed bool) string {
	if completed == true {
		return "Done"
	}
	return "Nah"
}
