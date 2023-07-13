package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExerciseDetailRequest Request Object
type ShowExerciseDetailRequest struct {

	// 需查询的习题id
	ExerciseId string `json:"exercise_id"`
}

func (o ShowExerciseDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExerciseDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowExerciseDetailRequest", string(data)}, " ")
}
