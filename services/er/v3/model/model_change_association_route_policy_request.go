package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAssociationRoutePolicyRequest Request Object
type ChangeAssociationRoutePolicyRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	// 关联ID
	AssociationId string `json:"association_id"`

	Body *AssociationRoutePolicyRequestBody `json:"body,omitempty"`
}

func (o ChangeAssociationRoutePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAssociationRoutePolicyRequest struct{}"
	}

	return strings.Join([]string{"ChangeAssociationRoutePolicyRequest", string(data)}, " ")
}
