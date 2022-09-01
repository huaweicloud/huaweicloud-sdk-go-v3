package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEvaluationProjectDetailRequest struct {

	// 评估项目ID。
	EvaluationProjectId string `json:"evaluation_project_id" xml:"evaluation_project_id"`
}

func (o ShowEvaluationProjectDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEvaluationProjectDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowEvaluationProjectDetailRequest", string(data)}, " ")
}
