package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExerciseDetailData 习题详细内容及测试用例信息
type ExerciseDetailData struct {
	ExerciseCodeResource *ExerciseCodeResource `json:"exercise_code_resource,omitempty"`

	// 测试用例信息
	ExerciseCaseResource *[]ExerciseCaseResource `json:"exercise_case_resource,omitempty"`
}

func (o ExerciseDetailData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExerciseDetailData struct{}"
	}

	return strings.Join([]string{"ExerciseDetailData", string(data)}, " ")
}
