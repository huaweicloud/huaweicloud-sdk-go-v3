package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTransitSubnetRequest Request Object
type DeleteTransitSubnetRequest struct {

	// 中转子网ID。
	TransitSubnetId string `json:"transit_subnet_id"`
}

func (o DeleteTransitSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitSubnetRequest struct{}"
	}

	return strings.Join([]string{"DeleteTransitSubnetRequest", string(data)}, " ")
}
