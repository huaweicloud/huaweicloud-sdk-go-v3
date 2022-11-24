package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CbcOrderResult struct {

	// 云服务ID
	CloudServiceId *string `json:"cloudServiceId,omitempty"`

	// 订单ID
	OrderId string `json:"orderId"`

	// 订购结果，1：成功；0：失败
	SubscribeResult int32 `json:"subscribeResult"`

	// 包周期资源预生成资源id。
	ResourceId *string `json:"resourceId,omitempty"`
}

func (o CbcOrderResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbcOrderResult struct{}"
	}

	return strings.Join([]string{"CbcOrderResult", string(data)}, " ")
}
