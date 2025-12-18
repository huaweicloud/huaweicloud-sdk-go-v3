package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DdmNodeInfo struct {

	// 节点ID。
	Id *string `json:"id,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 可用区编码。
	AzCode *string `json:"az_code,omitempty"`

	// 锁状态。
	Actions *[]ActionInfo `json:"actions,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o DdmNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdmNodeInfo struct{}"
	}

	return strings.Join([]string{"DdmNodeInfo", string(data)}, " ")
}
