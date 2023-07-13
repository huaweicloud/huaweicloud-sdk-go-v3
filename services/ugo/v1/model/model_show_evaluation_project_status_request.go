package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEvaluationProjectStatusRequest Request Object
type ShowEvaluationProjectStatusRequest struct {

	// 评估项目ID。
	EvaluationProjectId string `json:"evaluation_project_id"`
}

func (o ShowEvaluationProjectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEvaluationProjectStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowEvaluationProjectStatusRequest", string(data)}, " ")
}
