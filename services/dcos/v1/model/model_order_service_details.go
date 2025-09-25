package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderServiceDetails 服务单服务详情
type OrderServiceDetails struct {

	// 服务设备对象
	Devices *[]DeviceServiceDetails `json:"devices,omitempty"`

	// 资产对象
	Assets *[]AssetServiceDetails `json:"assets,omitempty"`

	// 人员信息
	Contacts *[]ContactServiceDetails `json:"contacts,omitempty"`
}

func (o OrderServiceDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderServiceDetails struct{}"
	}

	return strings.Join([]string{"OrderServiceDetails", string(data)}, " ")
}
