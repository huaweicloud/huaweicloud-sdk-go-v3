package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestartPolicyResponse Response Object
type ShowRestartPolicyResponse struct {

	// 是否重启虚拟机。
	RestartServer *bool `json:"restart_server,omitempty"`

	// 重启策略。
	RestartPolicy  *interface{} `json:"restart_policy,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowRestartPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestartPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowRestartPolicyResponse", string(data)}, " ")
}
