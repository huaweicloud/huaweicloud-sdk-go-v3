package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpandDdmInstanceNodesRequestBody struct {

	// 组id，指定当前进行节点扩容的组。当实例的组>1时，必填。
	GroupId *string `json:"group_id,omitempty"`

	// 子网ID，当组内节点的subnetId>1时，必填。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 变更包年包月实例规格时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。true，表示自动从账户中支付。false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// 节点信息列表。最小1，最大32
	Nodes []EnlargeNodeInfo `json:"nodes"`
}

func (o ExpandDdmInstanceNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDdmInstanceNodesRequestBody struct{}"
	}

	return strings.Join([]string{"ExpandDdmInstanceNodesRequestBody", string(data)}, " ")
}
