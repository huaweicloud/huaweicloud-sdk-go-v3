package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGdgwRouteTableRequest Request Object
type UpdateGdgwRouteTableRequest struct {

	// 全域接入网关ID
	GdgwId string `json:"gdgw_id"`

	Body *UpdateGdgwRoutetableRequestBody `json:"body,omitempty"`
}

func (o UpdateGdgwRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGdgwRouteTableRequest struct{}"
	}

	return strings.Join([]string{"UpdateGdgwRouteTableRequest", string(data)}, " ")
}
