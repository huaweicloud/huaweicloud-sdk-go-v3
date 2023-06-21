package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 获取习题库里习题过滤字段
type ExerciseFilter struct {

	// 需查询的习题名称
	Name *string `json:"name,omitempty"`

	// 习题类型列表
	ExerciseType *[]int32 `json:"exercise_type,omitempty"`

	// 难度id列表
	DifficultIds *[]string `json:"difficult_ids,omitempty"`

	// 知识点id列表
	KnowledgePointIds *[]string `json:"knowledge_point_ids,omitempty"`
}

func (o ExerciseFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExerciseFilter struct{}"
	}

	return strings.Join([]string{"ExerciseFilter", string(data)}, " ")
}
