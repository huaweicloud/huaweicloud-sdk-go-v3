package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransitSubnetRequest Request Object
type ShowTransitSubnetRequest struct {

	// 中转子网ID。
	TransitSubnetId string `json:"transit_subnet_id"`
}

func (o ShowTransitSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitSubnetRequest struct{}"
	}

	return strings.Join([]string{"ShowTransitSubnetRequest", string(data)}, " ")
}
