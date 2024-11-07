package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogtankRequestBody 创建云日志请求体。
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
