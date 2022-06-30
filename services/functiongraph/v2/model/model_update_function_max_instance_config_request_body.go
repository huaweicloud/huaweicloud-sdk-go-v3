package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新函数最大实例数请求体
type UpdateFunctionMaxInstanceConfigRequestBody struct {

	// 最大实例数
	MaxInstanceNum *int32 `json:"max_instance_num,omitempty"`
}

func (o UpdateFunctionMaxInstanceConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionMaxInstanceConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionMaxInstanceConfigRequestBody", string(data)}, " ")
}
