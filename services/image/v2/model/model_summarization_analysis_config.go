package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SummarizationAnalysisConfig struct {
	Common *SummarizationAnalysisConfigCommon `json:"common"`
}

func (o SummarizationAnalysisConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SummarizationAnalysisConfig struct{}"
	}

	return strings.Join([]string{"SummarizationAnalysisConfig", string(data)}, " ")
}
