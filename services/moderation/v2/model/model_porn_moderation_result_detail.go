package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 涉黄场景审核结果
type PornModerationResultDetail struct {
	// 置信度，取值范围 0-1。

	Confidence float32 `json:"confidence,omitempty"`
	// 当前支持label列表如下： - normal：正常 - porn：色情

	Label *string `json:"label,omitempty"`
}

func (o PornModerationResultDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PornModerationResultDetail struct{}"
	}

	return strings.Join([]string{"PornModerationResultDetail", string(data)}, " ")
}
