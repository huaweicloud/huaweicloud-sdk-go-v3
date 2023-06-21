package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 习题详细内容
type ExerciseCodeResource struct {

	// 习题内容存储id
	Id *string `json:"id,omitempty"`

	// 资源聚合id
	PolymericResourceId *string `json:"polymeric_resource_id,omitempty"`

	// 习题内容
	Content *string `json:"content,omitempty"`

	// 参考答案
	CodeAnswer *string `json:"code_answer,omitempty"`
}

func (o ExerciseCodeResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExerciseCodeResource struct{}"
	}

	return strings.Join([]string{"ExerciseCodeResource", string(data)}, " ")
}
