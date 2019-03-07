package raspi

import (
	"go-smart-home/pkg/models"
	"sync"
)

var confInstance *RaspConf
var once sync.Once

func GetConfInstance() *RaspConf {
	once.Do(func() {
		confInstance = &RaspConf{Switches: make(map[int64]*models.Switch)}
	})
	return confInstance
}
