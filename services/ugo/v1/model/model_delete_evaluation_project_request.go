package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEvaluationProjectRequest Request Object
type DeleteEvaluationProjectRequest struct {

	// 评估项目ID。
	EvaluationProjectId string `json:"evaluation_project_id"`
}

func (o DeleteEvaluationProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEvaluationProjectRequest struct{}"
	}

	return strings.Join([]string{"DeleteEvaluationProjectRequest", string(data)}, " ")
}
