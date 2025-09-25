package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageDetailRsp 镜像详情返回体
type ImageDetailRsp struct {

	// 镜像名称
	Name *string `json:"name,omitempty"`

	// 镜像ID
	Id *string `json:"id,omitempty"`

	// **参数解释**： 作业所属空间ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	EihealthProjectId *string `json:"eihealth_project_id,omitempty"`

	// **参数解释**： 作业所属空间名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	EihealthProjectName *string `json:"eihealth_project_name,omitempty"`

	// 镜像类型
	Type *string `json:"type,omitempty"`

	// 镜像芯片类型
	ChipType *string `json:"chip_type,omitempty"`

	// 镜像描述
	Description *string `json:"description,omitempty"`

	// 镜像版本列表
	Tags *[]string `json:"tags,omitempty"`

	// **参数解释**： 创建镜像的用户名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 创建镜像的用户ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 创建方式。 **约束限制**： 不涉及 **取值范围**： * PLATFORM_CREATED: 平台创建 * SWR_SYNC: swr同步 **默认取值**： 不涉及
	CreateType *string `json:"create_type,omitempty"`

	// 镜像创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 镜像更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 源项目名称
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 源项目id
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// 源资源id
	SourceResourceId *string `json:"source_resource_id,omitempty"`
}

func (o ImageDetailRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetailRsp struct{}"
	}

	return strings.Join([]string{"ImageDetailRsp", string(data)}, " ")
}
