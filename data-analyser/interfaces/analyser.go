package interfaces

import "data-monitor/data-common/models"

type Analyser interface {
	Analyser(models.Packet)
}
