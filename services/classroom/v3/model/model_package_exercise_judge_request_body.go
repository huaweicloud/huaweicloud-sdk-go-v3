package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 习题库里习题判题调用参数
type PackageExerciseJudgeRequestBody struct {

	// 判题结束后的回调url
	NotifyUrl string `json:"notify_url"`

	// 代码超时时间
	Timeout *int32 `json:"timeout,omitempty"`

	// 结果返回类型
	OutputType string `json:"output_type"`

	// 习题作答（需Base64编码）
	CodeAnswer string `json:"code_answer"`
}

func (o PackageExerciseJudgeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageExerciseJudgeRequestBody struct{}"
	}

	return strings.Join([]string{"PackageExerciseJudgeRequestBody", string(data)}, " ")
}
