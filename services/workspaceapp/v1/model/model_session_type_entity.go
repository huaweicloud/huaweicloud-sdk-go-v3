package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SessionTypeEntity struct {

	// 资源规格编码。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 会话类型： - CPU - GPU
	SessionType *string `json:"session_type,omitempty"`

	// 资源类型: -hws.resource.type.workspace.volume： 云办公桌面磁盘 -hws.resource.type.workspace.desktop： 云办公桌面 -hws.resource.type.workspace.appstream： 云应用 -hws.resource.type.workspace.appstreamsession： 云应用多会话
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源所属云服务类型编码： - hws.service.type.vdi - hws.service.type.marketplace
	CloudServiceType *string `json:"cloud_service_type,omitempty"`
}

func (o SessionTypeEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionTypeEntity struct{}"
	}

	return strings.Join([]string{"SessionTypeEntity", string(data)}, " ")
}
