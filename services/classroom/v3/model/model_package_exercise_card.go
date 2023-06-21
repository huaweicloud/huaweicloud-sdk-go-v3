package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 习题简单信息返回体，用于列表返回
type PackageExerciseCard struct {

	// 习题id
	Id *string `json:"id,omitempty"`

	// 习题名称
	Name *string `json:"name,omitempty"`

	Difficult *DifficultInfo `json:"difficult,omitempty"`

	// 习题类型编号
	ExerciseType *int32 `json:"exercise_type,omitempty"`

	// 习题类型名称
	ExerciseTypeName *string `json:"exercise_type_name,omitempty"`

	// 习题库里习题编号
	OrderCount *int32 `json:"order_count,omitempty"`

	// 相关知识点
	KnowledgePoint *[]KnowledgePointInfo `json:"knowledge_point,omitempty"`
}

func (o PackageExerciseCard) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageExerciseCard struct{}"
	}

	return strings.Join([]string{"PackageExerciseCard", string(data)}, " ")
}
