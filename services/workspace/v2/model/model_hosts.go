package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Hosts 主机信息
type Hosts struct {

	// 主机所属AZ。
	AvailabilityZone string `json:"availability_zone"`

	// 主机名称。
	Name string `json:"name"`

	// 在创建云服务器时（未指定主机ID），是否允许云服务器自动分配在一台可用的主机上。取值范围：“on”或“off”默认为 “on”。
	AutoPlacement *string `json:"auto_placement,omitempty"`

	// 主机类型。
	HostType string `json:"host_type"`

	// 待分配的主机数量
	Quantity int32 `json:"quantity"`

	// 主机产品Id
	ProductId string `json:"product_id"`
}

func (o Hosts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Hosts struct{}"
	}

	return strings.Join([]string{"Hosts", string(data)}, " ")
}
