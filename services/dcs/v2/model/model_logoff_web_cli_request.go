package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoffWebCliRequest Request Object
type LogoffWebCliRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *LogoutWebCliBody `json:"body,omitempty"`
}

func (o LogoffWebCliRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoffWebCliRequest struct{}"
	}

	return strings.Join([]string{"LogoffWebCliRequest", string(data)}, " ")
}
