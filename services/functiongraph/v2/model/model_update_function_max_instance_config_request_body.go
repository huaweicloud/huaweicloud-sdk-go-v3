package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFunctionMaxInstanceConfigRequestBody 更新函数最大实例数请求体
type UpdateFunctionMaxInstanceConfigRequestBody struct {

	// 最大实例数；-1代表该函数实例数无限制，0代表该函数被禁用
	MaxInstanceNum *int32 `json:"max_instance_num,omitempty"`
}

func (o UpdateFunctionMaxInstanceConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionMaxInstanceConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionMaxInstanceConfigRequestBody", string(data)}, " ")
}
