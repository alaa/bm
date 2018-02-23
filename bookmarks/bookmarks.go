package bookmarks

import "fmt"

type Bookmarks struct {
	Entries map[string][]string `yaml:""`
}

func (b *Bookmarks) List() {
	for k, v := range b.Entries {
		fmt.Printf("%s | %v", k, v)
	}
}
