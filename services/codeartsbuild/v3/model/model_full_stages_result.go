package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullStagesResult 结果
type FullStagesResult struct {

	// 构建步骤
	BuildStages map[string]BuildStageRecord `json:"build_stages,omitempty"`
}

func (o FullStagesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullStagesResult struct{}"
	}

	return strings.Join([]string{"FullStagesResult", string(data)}, " ")
}
