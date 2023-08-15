package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExercisesListRequestBody 获取习题库里习题调用参数
type ExercisesListRequestBody struct {
	Filter *ExerciseFilter `json:"filter,omitempty"`

	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`

	// 起始页
	StartIndex *int32 `json:"start_index,omitempty"`
}

func (o ExercisesListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExercisesListRequestBody struct{}"
	}

	return strings.Join([]string{"ExercisesListRequestBody", string(data)}, " ")
}
