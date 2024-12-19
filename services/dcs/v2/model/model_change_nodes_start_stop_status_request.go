package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeNodesStartStopStatusRequest Request Object
type ChangeNodesStartStopStatusRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ChangeNodesStartStopStatusBody `json:"body,omitempty"`
}

func (o ChangeNodesStartStopStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeNodesStartStopStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeNodesStartStopStatusRequest", string(data)}, " ")
}
