package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentDnsInfoResponse Response Object
type UpdateEquipmentDnsInfoResponse struct {

	// 主DNS
	MasterDns *string `json:"master_dns,omitempty"`

	// 备DNS
	SlaveDns       *string `json:"slave_dns,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEquipmentDnsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentDnsInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentDnsInfoResponse", string(data)}, " ")
}
