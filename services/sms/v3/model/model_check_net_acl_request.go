package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckNetAclRequest Request Object
type CheckNetAclRequest struct {

	// 目的虚拟机所属project_id
	TProjectId string `json:"t_project_id"`

	// 目的端子网ID
	TNetworkId string `json:"t_network_id"`

	// 区域ID
	RegionId string `json:"region_id"`

	// 操作系统类型
	OsType string `json:"os_type"`
}

func (o CheckNetAclRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNetAclRequest struct{}"
	}

	return strings.Join([]string{"CheckNetAclRequest", string(data)}, " ")
}
