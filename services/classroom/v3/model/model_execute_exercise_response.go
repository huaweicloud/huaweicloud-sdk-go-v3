package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteExerciseResponse struct {

	// 判题任务ID
	JudgementId    *string `json:"judgement_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteExerciseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteExerciseResponse struct{}"
	}

	return strings.Join([]string{"ExecuteExerciseResponse", string(data)}, " ")
}
