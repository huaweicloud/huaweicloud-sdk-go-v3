package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSqlLimitingSwitchNewRequest Request Object
type SetSqlLimitingSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SetSqlLimitingSwitchNewRequestBody `json:"body,omitempty"`
}

func (o SetSqlLimitingSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSqlLimitingSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"SetSqlLimitingSwitchNewRequest", string(data)}, " ")
}
