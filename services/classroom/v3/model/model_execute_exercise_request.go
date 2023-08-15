package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteExerciseRequest Request Object
type ExecuteExerciseRequest struct {

	// 具体调用者的用户token
	UserToken *string `json:"user-token,omitempty"`

	// 需判题的习题id
	ExerciseId string `json:"exercise_id"`

	Body *PackageExerciseJudgeRequestBody `json:"body,omitempty"`
}

func (o ExecuteExerciseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteExerciseRequest struct{}"
	}

	return strings.Join([]string{"ExecuteExerciseRequest", string(data)}, " ")
}
