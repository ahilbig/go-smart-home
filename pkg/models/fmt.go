package models

import "fmt"

func (m *Switch) String() string {
	desc := "<nil>"
	if m.Description != nil {
		desc = *m.Description
	}
	return fmt.Sprintf("{id=%d, state=%d, desc=%s}", m.ID, m.Status, desc)
}
