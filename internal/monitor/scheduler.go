package monitor

import (
	"fmt"
	"time"

	"Project/internal/repository"

	"github.com/charmbracelet/lipgloss"
)

func StartScheduler() {
	ticker := time.NewTicker(15 * time.Second)
	counter := 0

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")). // Purple | blue
		Padding(0, 2).
		MarginBottom(1).
		Width(45)

	labelStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")) // pink
	idStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("220"))               // yellow
	urlStyle := lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("43"))  // Cyan

	for range ticker.C {
		fmt.Printf("Running checks. Program is running for %d seconds.\n", counter)
		counter += 15

		targets, err := repository.GetAllTargets()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, target := range targets {
			content := fmt.Sprintf(
				"%s %s\n%s %s\n%s %s",
				labelStyle.Render("ID:  "), idStyle.Render(fmt.Sprintf("%v", target.ID)),
				labelStyle.Render("Name:"), target.Name,
				labelStyle.Render("URL: "), urlStyle.Render(target.URL),
			)
			fmt.Println(boxStyle.Render(content))
		}
		for _, target := range targets {
			go StartTargetWorker(target)
		}
	}
}
