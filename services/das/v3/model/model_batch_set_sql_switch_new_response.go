package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetSqlSwitchNewResponse Response Object
type BatchSetSqlSwitchNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchSetSqlSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSqlSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"BatchSetSqlSwitchNewResponse", string(data)}, " ")
}
