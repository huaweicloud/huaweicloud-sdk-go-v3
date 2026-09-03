package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetSqlSwitchNewRequest Request Object
type BatchSetSqlSwitchNewRequest struct {
	Body *BatchSetSqlSwitchNewRequestBody `json:"body,omitempty"`
}

func (o BatchSetSqlSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSqlSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"BatchSetSqlSwitchNewRequest", string(data)}, " ")
}
