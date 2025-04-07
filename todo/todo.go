package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

// ByPri implements sort.Interface for []Item based on
// the Priority & position field
type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}
	return s[i].Priority < s[j].Priority
}

func SaveItems(filname string, items []Item) error {

	b, err := json.Marshal(items)

	if err != nil {
		return err
	}
	err = os.WriteFile(filname, b, 0644)

	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func ReadItems(filname string) ([]Item, error) {

	var items []Item

	b, err := os.ReadFile(filname)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &items); err != nil {
		return nil, err
	}

	for i, _ := range items {
		items[i].position = i + 1
	}

	return items, nil

}

func (i *Item) PrettyP() string {
	switch i.Priority {
	case 1:
		return "(1)"
	case 3:
		return "(3)"
	default:
		return " "
	}
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) MarkDone() {
	i.Done = true
}

func (i *Item) ToggleDone() {
	i.Done = !i.Done
}
