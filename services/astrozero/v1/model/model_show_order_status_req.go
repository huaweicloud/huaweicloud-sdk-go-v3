package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderStatusReq 查询活动时间被租户开通信息
type ShowOrderStatusReq struct {

	// 租户id
	DomainId string `json:"domain_id"`

	// 活动开始时间,必填
	StartTime string `json:"start_time"`

	// 活动结束时间,必填
	EndTime string `json:"end_time"`
}

func (o ShowOrderStatusReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderStatusReq struct{}"
	}

	return strings.Join([]string{"ShowOrderStatusReq", string(data)}, " ")
}
