package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetIndexUsageSwitchNewRequest Request Object
type SetIndexUsageSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SetIndexUsageSwitchNewRequestBody `json:"body,omitempty"`
}

func (o SetIndexUsageSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetIndexUsageSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"SetIndexUsageSwitchNewRequest", string(data)}, " ")
}
