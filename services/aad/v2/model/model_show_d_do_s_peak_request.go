package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDDoSPeakRequest Request Object
type ShowDDoSPeakRequest struct {

	// 开始时间（毫秒时间戳）
	StartTime string `json:"start_time"`

	// 结束时间（毫秒时间戳）
	EndTime string `json:"end_time"`

	// 实例id
	InstanceId string `json:"instance_id"`

	// 高防IP
	Ip string `json:"ip"`
}

func (o ShowDDoSPeakRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDDoSPeakRequest struct{}"
	}

	return strings.Join([]string{"ShowDDoSPeakRequest", string(data)}, " ")
}
