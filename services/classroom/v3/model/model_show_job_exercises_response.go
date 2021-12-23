package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobExercisesResponse struct {
	// 作业下习题列表

	GroupExercises *[]ExerciseGroup `json:"group_exercises,omitempty"`
	// 作业下习题总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowJobExercisesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobExercisesResponse struct{}"
	}

	return strings.Join([]string{"ShowJobExercisesResponse", string(data)}, " ")
}
