package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobSuccessRatioResponse struct {
	// 任务成功构建次数

	SuccessCount *int32 `json:"success_count,omitempty"`
	// 任务构建总次数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 任务成功率,精确到小数点后两位

	SuccessRatio   *float64 `json:"success_ratio,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowJobSuccessRatioResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobSuccessRatioResponse struct{}"
	}

	return strings.Join([]string{"ShowJobSuccessRatioResponse", string(data)}, " ")
}
