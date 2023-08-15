package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociatedTransitIp 关联的中转IP。
type AssociatedTransitIp struct {

	// 中转IP的ID。
	TransitIpId *string `json:"transit_ip_id,omitempty"`

	// 中转IP地址。
	TransitIpAddress *string `json:"transit_ip_address,omitempty"`
}

func (o AssociatedTransitIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatedTransitIp struct{}"
	}

	return strings.Join([]string{"AssociatedTransitIp", string(data)}, " ")
}
