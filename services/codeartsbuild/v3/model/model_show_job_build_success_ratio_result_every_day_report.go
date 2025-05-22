package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowJobBuildSuccessRatioResultEveryDayReport struct {

	// 日期
	Date *string `json:"date,omitempty"`

	// 成功率
	SuccessRatio *int32 `json:"success_ratio,omitempty"`
}

func (o ShowJobBuildSuccessRatioResultEveryDayReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobBuildSuccessRatioResultEveryDayReport struct{}"
	}

	return strings.Join([]string{"ShowJobBuildSuccessRatioResultEveryDayReport", string(data)}, " ")
}
