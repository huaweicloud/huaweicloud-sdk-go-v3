package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDDosRequest struct {

	// 用户EIP对应的ID
	FloatingIpId string `json:"floating_ip_id" xml:"floating_ip_id"`

	// ip
	Ip *string `json:"ip,omitempty" xml:"ip"`

	Body *UpdateAntiDDosServiceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateDDosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDDosRequest struct{}"
	}

	return strings.Join([]string{"UpdateDDosRequest", string(data)}, " ")
}
