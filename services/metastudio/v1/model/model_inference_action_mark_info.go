package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InferenceActionMarkInfo 分身数字人推理预处理动作标记信息。
type InferenceActionMarkInfo struct {

	// 动作列表。
	ActionInfo *[]ActionMarkItem `json:"action_info,omitempty"`
}

func (o InferenceActionMarkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InferenceActionMarkInfo struct{}"
	}

	return strings.Join([]string{"InferenceActionMarkInfo", string(data)}, " ")
}
