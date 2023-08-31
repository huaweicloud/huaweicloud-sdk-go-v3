package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EquipmentDnsItem struct {

	// 主DNS
	MasterDns *string `json:"master_dns,omitempty"`

	// 备DNS
	SlaveDns *string `json:"slave_dns,omitempty"`
}

func (o EquipmentDnsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EquipmentDnsItem struct{}"
	}

	return strings.Join([]string{"EquipmentDnsItem", string(data)}, " ")
}
