package license

import "fmt"

type License struct {
	ID    uint32
	Title string
}

func (l *License) String() string {
	return fmt.Sprintf("Lic: [ID: %d, Title: - %s]",
		l.ID,
		l.Title)
}
