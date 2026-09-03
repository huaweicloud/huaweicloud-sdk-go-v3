package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DasCommonInstanceNodeDto 实例节点信息
type DasCommonInstanceNodeDto struct {

	// 节点ID
	Id *string `json:"id,omitempty"`

	// 节点名称
	Name *string `json:"name,omitempty"`

	// 节点状态
	Status *string `json:"status,omitempty"`

	// 节点角色
	Role *string `json:"role,omitempty"`

	// 节点私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 节点公共IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 节点组ID
	GroupId *string `json:"group_id,omitempty"`

	// 节点组名称
	GroupName *string `json:"group_name,omitempty"`
}

func (o DasCommonInstanceNodeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DasCommonInstanceNodeDto struct{}"
	}

	return strings.Join([]string{"DasCommonInstanceNodeDto", string(data)}, " ")
}
