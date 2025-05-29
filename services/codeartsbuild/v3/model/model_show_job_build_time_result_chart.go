package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowJobBuildTimeResultChart struct {

	// 构建每日编号
	DailyBuildNumber *string `json:"daily_build_number,omitempty"`

	// 构建编号
	BuildNumber *int32 `json:"build_number,omitempty"`

	// 步骤执行时长，单位s
	BuildTime *int32 `json:"build_time,omitempty"`

	// 构建结果
	BuildResult *string `json:"build_result,omitempty"`
}

func (o ShowJobBuildTimeResultChart) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobBuildTimeResultChart struct{}"
	}

	return strings.Join([]string{"ShowJobBuildTimeResultChart", string(data)}, " ")
}
