package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateSubnetRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id" xml:"routetable_id"`

	Body *AssociateSubnetRequestBody `json:"body,omitempty" xml:"body"`
}

func (o AssociateSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSubnetRequest struct{}"
	}

	return strings.Join([]string{"AssociateSubnetRequest", string(data)}, " ")
}
