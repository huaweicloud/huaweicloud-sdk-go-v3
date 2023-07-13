package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFunctionTagsRequestBody struct {

	// action名称
	Action *string `json:"action,omitempty"`

	// 标签列表
	Tags *[]map[string]string `json:"tags,omitempty"`

	// 系统标签列表
	SysTags *[]map[string]string `json:"sys_tags,omitempty"`
}

func (o UpdateFunctionTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionTagsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionTagsRequestBody", string(data)}, " ")
}
