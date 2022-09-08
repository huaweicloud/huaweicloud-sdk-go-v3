package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务执行结果
type JudgementResult struct {

	// 标准类型输出结果
	Output string `json:"output"`

	// 文件形式输出的文件id，可根据文件id下载详情
	FileId string `json:"file_id"`

	// 图片形式输出的图片id，可根据图片id下载详情
	ImageId string `json:"image_id"`

	// 用例形式输出的用例总个数
	CaseCount int32 `json:"case_count"`

	// 用例形式输出的已执行用例的个数
	ExecutedCount int32 `json:"executed_count"`

	// 用例形式输出的已执行用例的结果
	Testcases []JudgementCaseResult `json:"testcases"`
}

func (o JudgementResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JudgementResult struct{}"
	}

	return strings.Join([]string{"JudgementResult", string(data)}, " ")
}
