package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTriggerUpgradeRequestBody 触发升级请求
type ExecuteTriggerUpgradeRequestBody struct {

	// 是否强制升级：0-否 1-是
	IsForceExecute int32 `json:"is_force_execute"`

	// 升级目标版本
	TargetVersion string `json:"target_version"`

	// 升级任务描述
	Description *string `json:"description,omitempty"`

	// 通知开启：0-未开启 1-开启
	IsNotify int32 `json:"is_notify"`

	// 扩展参数（JSON格式）
	ExtraParams *string `json:"extra_params,omitempty"`

	// 桌面sids列表
	DesktopSids *[]string `json:"desktop_sids,omitempty"`
}

func (o ExecuteTriggerUpgradeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTriggerUpgradeRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteTriggerUpgradeRequestBody", string(data)}, " ")
}
