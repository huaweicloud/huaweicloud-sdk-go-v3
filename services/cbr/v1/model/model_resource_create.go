package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceCreate struct {
	ExtraInfo *ResourceExtraInfo `json:"extra_info,omitempty"`

	// **参数解释：** 待备份资源ID，获取方法请参见[查询存储库绑定资源信息](ShowVault.xml)，[查询资源可保护性](ShowProtectable.xml) **约束限制：** 需要该资源暂未绑定到存储库中，且属于可备份的状态 **取值范围：** 不涉及 **默认取值：** 不涉及
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
