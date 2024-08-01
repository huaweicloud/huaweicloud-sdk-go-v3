package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAnalyzerReqBody struct {
	Configuration *AnalyzerConfiguration `json:"configuration,omitempty"`

	// 分析器的名称。
	Name string `json:"name"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	Type *AnalyzerType `json:"type"`
}

func (o CreateAnalyzerReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnalyzerReqBody struct{}"
	}

	return strings.Join([]string{"CreateAnalyzerReqBody", string(data)}, " ")
}
