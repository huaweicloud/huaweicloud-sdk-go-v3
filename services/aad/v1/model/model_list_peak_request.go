package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPeakRequest Request Object
type ListPeakRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// 单个 IP
	Ip string `json:"ip"`

	// 开始时间，毫秒时间戳
	StartTime string `json:"start_time"`

	// 结束时间，毫秒时间戳
	EndTime string `json:"end_time"`
}

func (o ListPeakRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPeakRequest struct{}"
	}

	return strings.Join([]string{"ListPeakRequest", string(data)}, " ")
}
