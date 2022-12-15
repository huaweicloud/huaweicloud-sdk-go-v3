package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceCreate struct {
	ExtraInfo *ResourceExtraInfo `json:"extra_info,omitempty"`

	// 待备份资源id
	Id string `json:"id"`

	// 待备份资源的类型, 云服务器: OS::Nova::Server, 云硬盘: OS::Cinder::Volume, 裸金属服务器: OS::Ironic::BareMetalServer, 线下本地服务器: OS::Native::Server, 弹性文件系统: OS::Sfs::Turbo, 云桌面：OS::Workspace::DesktopV2
	Type string `json:"type"`

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o ResourceCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceCreate struct{}"
	}

	return strings.Join([]string{"ResourceCreate", string(data)}, " ")
}
