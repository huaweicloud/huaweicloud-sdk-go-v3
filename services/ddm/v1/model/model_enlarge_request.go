package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnlargeRequest This is a auto request Object
type EnlargeRequest struct {

	// 当前进行节点扩容的DDM实例底层虚机规格id
	FlavorId string `json:"flavor_id"`

	// 需要扩容的节点个数
	NodeNumber int32 `json:"node_number"`

	// 组id，指定当前进行节点扩容的组。当实例的组>1时，必填。
	GroupId *string `json:"group_id,omitempty"`

	// 变更包年包月实例规格时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。true，表示自动从账户中支付。false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// 可用区Code，仅包年包月实例传递该参数，个数需与node_number一致。请参见地区和终端节点(https://developer.huaweicloud.com/endpoint?DDM)。
	AvailableZones *[]string `json:"available_zones,omitempty"`
}

func (o EnlargeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeRequest struct{}"
	}

	return strings.Join([]string{"EnlargeRequest", string(data)}, " ")
}
