package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTransitIpRequest Request Object
type DeleteTransitIpRequest struct {

	// 中转IP的ID。
	TransitIpId string `json:"transit_ip_id"`
}

func (o DeleteTransitIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitIpRequest struct{}"
	}

	return strings.Join([]string{"DeleteTransitIpRequest", string(data)}, " ")
}
