package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowServiceDetailRequest struct {
	// 服务ID

	ServiceId string `json:"service_id"`
	// 铂金版实例ID

	IefInstanceId string `json:"ief-instance-id"`
}

func (o ShowServiceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServiceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowServiceDetailRequest", string(data)}, " ")
}
