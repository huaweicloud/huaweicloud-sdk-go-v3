package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PcrTestRecordResult
type PcrTestRecordResult struct {

	// 姓名
	Name string `json:"name"`

	// 核酸检测采样时间
	SamplingTime string `json:"sampling_time"`

	// 核酸检测结果更新时间
	TestTime string `json:"test_time"`

	// 核酸检测结果，可选值包括：  - \"positive\",即阳性  - \"negative\",即阴性  - \"unknown\",未知
	TestResult string `json:"test_result"`

	Confidence *PcrTestRecordConfidence `json:"confidence"`

	// 代表检测识别出来的文字块数目。
	WordsBlockCount int32 `json:"words_block_count"`

	// 识别文字块列表，输出顺序从左到右，从上到下。
	WordsBlockList []PcrTestRecordWordsBlockList `json:"words_block_list"`
}

func (o PcrTestRecordResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PcrTestRecordResult struct{}"
	}

	return strings.Join([]string{"PcrTestRecordResult", string(data)}, " ")
}
