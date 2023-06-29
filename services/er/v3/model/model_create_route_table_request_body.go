package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRouteTableRequestBody This is a auto create Body Object
type CreateRouteTableRequestBody struct {
	RouteTable *CreateRouteTable `json:"route_table,omitempty"`
}

func (o CreateRouteTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRouteTableRequestBody", string(data)}, " ")
}
