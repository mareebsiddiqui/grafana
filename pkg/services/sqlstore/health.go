package sqlstore

import (
	"github.com/mareebsiddiqui/grafana/pkg/bus"
	m "github.com/mareebsiddiqui/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
