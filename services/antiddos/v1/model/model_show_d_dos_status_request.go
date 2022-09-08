package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDDosStatusRequest struct {

	// 用户EIP对应的ID
	FloatingIpId string `json:"floating_ip_id"`

	// 用户EIP
	Ip *string `json:"ip,omitempty"`
}

func (o ShowDDosStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDDosStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowDDosStatusRequest", string(data)}, " ")
}
