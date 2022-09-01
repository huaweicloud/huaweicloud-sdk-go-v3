package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateRouteTableRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id" xml:"routetable_id"`

	Body *RoutetableAssociateReqbody `json:"body,omitempty" xml:"body"`
}

func (o DisassociateRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouteTableRequest struct{}"
	}

	return strings.Join([]string{"DisassociateRouteTableRequest", string(data)}, " ")
}
