package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExerciseCaseResource 测试用例信息
type ExerciseCaseResource struct {

	// 测试用例存储id
	Id *string `json:"id,omitempty"`

	// 资源聚合id
	PolymericResourceId *string `json:"polymeric_resource_id,omitempty"`

	// 用例输入
	InputFile *string `json:"input_file,omitempty"`

	// 用例输出
	OutputFile *string `json:"output_file,omitempty"`

	// 用例顺序位置
	Index *int32 `json:"index,omitempty"`

	// 用例类型
	InputType *string `json:"input_type,omitempty"`
}

func (o ExerciseCaseResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExerciseCaseResource struct{}"
	}

	return strings.Join([]string{"ExerciseCaseResource", string(data)}, " ")
}
