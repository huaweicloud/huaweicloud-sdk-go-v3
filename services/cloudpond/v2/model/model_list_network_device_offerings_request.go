package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworkDeviceOfferingsRequest Request Object
type ListNetworkDeviceOfferingsRequest struct {

	// 地区编码
	ZoneCode *string `json:"zone_code,omitempty"`

	// 存储类型
	StorageType *[]string `json:"storage_type,omitempty"`

	// 付费模式
	PayMode *[]PayMode `json:"pay_mode,omitempty"`

	// 购买时长
	PeriodNum *[]int32 `json:"period_num,omitempty"`

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`
}

func (o ListNetworkDeviceOfferingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkDeviceOfferingsRequest struct{}"
	}

	return strings.Join([]string{"ListNetworkDeviceOfferingsRequest", string(data)}, " ")
}
