package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeliveryInfo 交付信息
type DeliveryInfo struct {

	// 类型，收货/发货
	Type *string `json:"type,omitempty"`

	// 起始时间，UTC格式
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，UTC格式
	EndTime *string `json:"end_time,omitempty"`

	// 机房编码
	RoomCode *string `json:"room_code,omitempty"`

	// 联系人
	Contacts *[]ContactInformation `json:"contacts,omitempty"`

	// 资产清单
	Assets *[]AssetInfo `json:"assets,omitempty"`
}

func (o DeliveryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeliveryInfo struct{}"
	}

	return strings.Join([]string{"DeliveryInfo", string(data)}, " ")
}
