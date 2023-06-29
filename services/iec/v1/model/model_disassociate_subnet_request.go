package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateSubnetRequest Request Object
type DisassociateSubnetRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id"`

	Body *DisassociateSubnetRequestBody `json:"body,omitempty"`
}

func (o DisassociateSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateSubnetRequest struct{}"
	}

	return strings.Join([]string{"DisassociateSubnetRequest", string(data)}, " ")
}
