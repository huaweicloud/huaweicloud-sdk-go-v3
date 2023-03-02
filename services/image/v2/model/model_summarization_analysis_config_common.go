package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SummarizationAnalysisConfigCommon struct {
	Inference *SummarizationAnalysisInference `json:"inference"`
}

func (o SummarizationAnalysisConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SummarizationAnalysisConfigCommon struct{}"
	}

	return strings.Join([]string{"SummarizationAnalysisConfigCommon", string(data)}, " ")
}
