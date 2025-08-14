package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssessTaskRequestBody 创建应用评估任务请求体
type CreateAssessTaskRequestBody struct {

	// 应用ID
	ApplicationId string `json:"application_id"`
}

func (o CreateAssessTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssessTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAssessTaskRequestBody", string(data)}, " ")
}
