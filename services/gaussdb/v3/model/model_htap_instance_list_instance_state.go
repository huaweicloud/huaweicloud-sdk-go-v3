package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HtapInstanceListInstanceState HTAP实例状态信息。
type HtapInstanceListInstanceState struct {

	// HTAP实例状态。
	InstanceStatus *string `json:"instance_status,omitempty"`

	// HTAP实例创建失败错误码。
	CreateFailErrorCode *string `json:"create_fail_error_code,omitempty"`

	// HTAP实例创建失败错误信息。
	FailMessage *string `json:"fail_message,omitempty"`

	// 是否需要重启更新参数。
	WaitRestartForParams *bool `json:"wait_restart_for_params,omitempty"`
}

func (o HtapInstanceListInstanceState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapInstanceListInstanceState struct{}"
	}

	return strings.Join([]string{"HtapInstanceListInstanceState", string(data)}, " ")
}
