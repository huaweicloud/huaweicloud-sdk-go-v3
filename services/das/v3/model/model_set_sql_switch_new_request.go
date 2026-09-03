package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSqlSwitchNewRequest Request Object
type SetSqlSwitchNewRequest struct {
	Body *SetSqlSwitchNewRequestBody `json:"body,omitempty"`
}

func (o SetSqlSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSqlSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"SetSqlSwitchNewRequest", string(data)}, " ")
}
