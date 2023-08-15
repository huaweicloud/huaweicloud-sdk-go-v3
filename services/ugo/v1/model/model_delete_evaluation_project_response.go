package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEvaluationProjectResponse Response Object
type DeleteEvaluationProjectResponse struct {

	// 评估项目ID。
	EvaluationProjectId *int32 `json:"evaluation_project_id,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o DeleteEvaluationProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEvaluationProjectResponse struct{}"
	}

	return strings.Join([]string{"DeleteEvaluationProjectResponse", string(data)}, " ")
}
