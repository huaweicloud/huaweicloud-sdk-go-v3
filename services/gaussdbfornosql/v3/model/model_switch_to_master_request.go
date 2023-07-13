package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchToMasterRequest Request Object
type SwitchToMasterRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SwitchToMasterDisasterRecoveryBody `json:"body,omitempty"`
}

func (o SwitchToMasterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchToMasterRequest struct{}"
	}

	return strings.Join([]string{"SwitchToMasterRequest", string(data)}, " ")
}
