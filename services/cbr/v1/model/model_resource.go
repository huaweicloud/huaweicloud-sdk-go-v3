package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {
	ExtraInfo *ResourceExtraInfo `json:"extra_info,omitempty" xml:"extra_info"`

	// 待备份资源id
	Id string `json:"id" xml:"id"`

	// 待备份资源名称，长度限制：0-255
	Name *string `json:"name,omitempty" xml:"name"`

	// 待备份资源的类型, 云服务器: OS::Nova::Server, 云硬盘: OS::Cinder::Volume, 裸金属服务器: OS::Ironic::BareMetalServer, 线下本地服务器: OS::Native::Server, 弹性文件系统: OS::Sfs::Turbo
	Type string `json:"type" xml:"type"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
