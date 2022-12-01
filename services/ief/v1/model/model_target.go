package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量处理对象详情
type Target struct {

	// 批量处理对象基本信息
	ExtensionInfo *[]Extension `json:"extension_info,omitempty"`

	// 批量处理对象ID
	TargetId string `json:"target_id"`
}

func (o Target) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Target struct{}"
	}

	return strings.Join([]string{"Target", string(data)}, " ")
}
