package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransitIpRequest Request Object
type ShowTransitIpRequest struct {

	// 中转IP的ID。
	TransitIpId string `json:"transit_ip_id"`
}

func (o ShowTransitIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitIpRequest struct{}"
	}

	return strings.Join([]string{"ShowTransitIpRequest", string(data)}, " ")
}
