package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryShipperRequest Request Object
type RetryShipperRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 投递ID
	ShipperId string `json:"shipper_id"`
}

func (o RetryShipperRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryShipperRequest struct{}"
	}

	return strings.Join([]string{"RetryShipperRequest", string(data)}, " ")
}
