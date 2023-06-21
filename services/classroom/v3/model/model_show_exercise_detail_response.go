package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowExerciseDetailResponse struct {

	// 习题id
	Id *string `json:"id,omitempty"`

	// 习题名称
	Name *string `json:"name,omitempty"`

	// 习题描述
	Description *string `json:"description,omitempty"`

	Difficult *DifficultInfo `json:"difficult,omitempty"`

	// 习题类型编号
	ExerciseType *int32 `json:"exercise_type,omitempty"`

	// 习题类型名称
	ExerciseTypeName *string `json:"exercise_type_name,omitempty"`

	// 习题库里习题编号
	OrderCount *int32 `json:"order_count,omitempty"`

	// 测试用例描述
	TestCaseDescription *string `json:"test_case_description,omitempty"`

	// 相关知识点
	KnowledgePoint *[]KnowledgePointInfo `json:"knowledge_point,omitempty"`

	// 判题类型
	JudgeType *int32 `json:"judge_type,omitempty"`

	ExerciseData   *ExerciseDetailData `json:"exercise_data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowExerciseDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExerciseDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowExerciseDetailResponse", string(data)}, " ")
}
