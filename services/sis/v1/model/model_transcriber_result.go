package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TranscriberResult struct {
	// 识别结果文本。

	Text string `json:"text"`

	AnalysisInfo *AnalysisInfoResult `json:"analysis_info,omitempty"`
	// 分词输出列表

	WordInfo *[]WordInfo `json:"word_info,omitempty"`
}

func (o TranscriberResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscriberResult struct{}"
	}

	return strings.Join([]string{"TranscriberResult", string(data)}, " ")
}
