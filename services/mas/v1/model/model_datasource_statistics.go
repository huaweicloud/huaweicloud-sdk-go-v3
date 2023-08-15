package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatasourceStatistics struct {
	Count *int32 `json:"count,omitempty"`

	MysqlCount *int32 `json:"mysql_count,omitempty"`
}

func (o DatasourceStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasourceStatistics struct{}"
	}

	return strings.Join([]string{"DatasourceStatistics", string(data)}, " ")
}
