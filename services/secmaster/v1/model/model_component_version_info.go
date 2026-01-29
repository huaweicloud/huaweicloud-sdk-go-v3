package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentVersionInfo 插件版本信息
type ComponentVersionInfo struct {

	// 版本ID
	Id *string `json:"id,omitempty"`

	// 版本号
	VersionNum *string `json:"version_num,omitempty"`

	// 版本描述
	VersionDesc *string `json:"version_desc,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 插件所属包
	PackageName *string `json:"package_name,omitempty"`

	// 插件的附件id
	ComponentAttachmentId *string `json:"component_attachment_id,omitempty"`

	// 订阅版本id
	SubVersionId *string `json:"sub_version_id,omitempty"`

	// 操作连接配置List
	ConnectionConfig *string `json:"connection_config,omitempty"`
}

func (o ComponentVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentVersionInfo struct{}"
	}

	return strings.Join([]string{"ComponentVersionInfo", string(data)}, " ")
}
