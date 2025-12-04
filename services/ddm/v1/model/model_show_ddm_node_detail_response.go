package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdmNodeDetailResponse Response Object
type ShowDdmNodeDetailResponse struct {

	// 节点id。
	Id *string `json:"id,omitempty"`

	// 私有ip。
	PrivateIp *string `json:"private_ip,omitempty"`

	// 节点状态。
	Status *string `json:"status,omitempty"`

	// 节点名称。
	Name *string `json:"name,omitempty"`

	// 子网id。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 可用区。
	AzCode *string `json:"az_code,omitempty"`

	// 组id。
	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDdmNodeDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdmNodeDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDdmNodeDetailResponse", string(data)}, " ")
}
