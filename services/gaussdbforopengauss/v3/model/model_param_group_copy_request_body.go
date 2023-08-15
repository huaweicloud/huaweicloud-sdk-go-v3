package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParamGroupCopyRequestBody struct {

	// 复制后的参数模板名称。
	Name string `json:"name"`

	// 参数模板描述。
	Description *string `json:"description,omitempty"`
}

func (o ParamGroupCopyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamGroupCopyRequestBody struct{}"
	}

	return strings.Join([]string{"ParamGroupCopyRequestBody", string(data)}, " ")
}
