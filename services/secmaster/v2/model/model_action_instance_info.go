package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionInstanceInfo Action Instance
type ActionInstanceInfo struct {
	Action *ActionInfo `json:"action,omitempty"`

	InstanceLog *AuditLogInfo `json:"instance_log,omitempty"`
}

func (o ActionInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionInstanceInfo struct{}"
	}

	return strings.Join([]string{"ActionInstanceInfo", string(data)}, " ")
}
