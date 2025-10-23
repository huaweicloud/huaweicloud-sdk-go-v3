package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatisticsParamsDto **参数解释：** 统计任务信息
type StatisticsParamsDto struct {

	// **参数解释：** 分支名
	BranchName *string `json:"branch_name,omitempty"`
}

func (o StatisticsParamsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsParamsDto struct{}"
	}

	return strings.Join([]string{"StatisticsParamsDto", string(data)}, " ")
}
