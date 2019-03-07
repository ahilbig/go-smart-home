// Raspberry specific implementation

package raspi

import (
	"fmt"
	"go-smart-home/pkg/models"
)

type RaspConf struct {
	Switches map[int64]*models.Switch
}

func (r *RaspConf) CreateSwitch(s *models.Switch) (*models.Switch, error) {
	if _, exists := r.Switches[s.ID]; exists {
		return nil, fmt.Errorf("switch with id=%d already exists", s.ID)
	}
	r.Switches[s.ID] = s
	return s, nil
}

func (r *RaspConf) UpdateSwitch(s *models.Switch) error {
	if _, exists := r.Switches[s.ID]; !exists {
		return fmt.Errorf("switch with id=%d doesn't exist", s.ID)
	}
	r.Switches[s.ID] = s
	return nil
}

func (r *RaspConf) GetSwitches() []*models.Switch {
	arr := make([]*models.Switch, 0, len(r.Switches))
	for _, v := range r.Switches {
		if v != nil {
			arr = append(arr, v)
		}
	}
	return arr
}
