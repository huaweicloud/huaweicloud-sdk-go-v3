package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTransitSubnetRequest Request Object
type UpdateTransitSubnetRequest struct {

	// 中转子网ID。
	TransitSubnetId string `json:"transit_subnet_id"`

	Body *UpdateTransitSubnetRequestBody `json:"body,omitempty"`
}

func (o UpdateTransitSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransitSubnetRequest struct{}"
	}

	return strings.Join([]string{"UpdateTransitSubnetRequest", string(data)}, " ")
}
