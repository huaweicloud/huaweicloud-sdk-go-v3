package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcRdsHa 备机同步参数
type HwcRdsHa struct {

	// 备机同步参数。  取值：非空。  RDS for MySQL为“async”或“semisync”。 RDS for PostgreSQL为“async”或“sync”。 RDS for Microsoft SQL Server为“sync”。 说明： “async”为异步模式。 “semisync”为半同步模式。 “sync”为同步模式。
	ReplicationMode string `json:"replication_mode"`
}

func (o HwcRdsHa) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcRdsHa struct{}"
	}

	return strings.Join([]string{"HwcRdsHa", string(data)}, " ")
}
