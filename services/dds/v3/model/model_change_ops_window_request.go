package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeOpsWindowRequest Request Object
type ChangeOpsWindowRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *OpsWindowRequestBody `json:"body,omitempty"`
}

func (o ChangeOpsWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOpsWindowRequest struct{}"
	}

	return strings.Join([]string{"ChangeOpsWindowRequest", string(data)}, " ")
}
