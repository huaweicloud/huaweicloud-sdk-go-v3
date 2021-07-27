package model

import (
	"encoding/json"

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
}

func (o JudgementResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JudgementResult struct{}"
	}

	return strings.Join([]string{"JudgementResult", string(data)}, " ")
}
