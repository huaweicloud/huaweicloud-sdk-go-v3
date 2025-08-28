package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogtankRequestBody **参数解释**：创建云日志请求参数。  **约束限制**：不涉及
type CreateLogtankRequestBody struct {
	Logtank *CreateLogtankOption `json:"logtank"`
}

func (o CreateLogtankRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLogtankRequestBody", string(data)}, " ")
}
