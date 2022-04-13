package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建云日志请求体
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
