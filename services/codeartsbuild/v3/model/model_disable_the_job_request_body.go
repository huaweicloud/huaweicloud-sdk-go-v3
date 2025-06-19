package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableTheJobRequestBody 禁用构建任务请求体
type DisableTheJobRequestBody struct {

	// 禁用原因
	Reason *string `json:"reason,omitempty"`

	// 是否禁用
	Disabled bool `json:"disabled"`
}

func (o DisableTheJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableTheJobRequestBody struct{}"
	}

	return strings.Join([]string{"DisableTheJobRequestBody", string(data)}, " ")
}
