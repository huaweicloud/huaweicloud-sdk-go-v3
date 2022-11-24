package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateRouteTableRequestBody struct {
	RouteTable *UpdateRouteTable `json:"route_table,omitempty"`
}

func (o UpdateRouteTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteTableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRouteTableRequestBody", string(data)}, " ")
}
