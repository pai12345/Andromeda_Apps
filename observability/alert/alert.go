package alert

import "fmt"

func (alr Alert) FetchAlert() *Alert {
	fmt.Println(alr)
	return &alr
}

func (alr Alert) AlertLevels() *Levels {
	levels := Levels{
		"Emergency":   true,
		"Critical":    true,
		"Error":       true,
		"Warning":     true,
		"Information": true,
	}
	return &levels
}

func (alr Alert) Page() bool {
	status := false
	result := *alr.AlertLevels()
	if result[alr.Severity] {
		fmt.Printf("Paged squad %v", alr.Squad)
		status = true
	}
	return status
}

func CollectivePage(page Pager) bool {
	status := page.Page()
	return status
}
