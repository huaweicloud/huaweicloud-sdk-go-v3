package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectRsp 项目信息
type ProjectRsp struct {

	// 项目id
	Id *string `json:"id,omitempty"`

	// 项目名称
	Name *string `json:"name,omitempty"`

	// 项目所有者
	Creator *string `json:"creator,omitempty"`

	// 当前用户在该项目上的角色
	Role *string `json:"role,omitempty"`

	// 项目桶存储量
	Size *int64 `json:"size,omitempty"`

	// 项目状态
	Status *string `json:"status,omitempty"`

	// 标签列表
	Tags *[]string `json:"tags,omitempty"`

	// 项目描述
	Description *string `json:"description,omitempty"`

	// 项目创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 项目更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 空间置顶时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TopTime *string `json:"top_time,omitempty"`

	// **参数解释**： 请求删除时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DeleteTime *string `json:"delete_time,omitempty"`

	// **参数解释**： 核心空间标记。 **约束限制**： 不涉及 **取值范围**： true 或 false。 **默认取值**： 不涉及
	IsCore *bool `json:"is_core,omitempty"`

	// **参数解释**： 空间桶来源租户。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	StorageTenant *string `json:"storage_tenant,omitempty"`
}

func (o ProjectRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectRsp struct{}"
	}

	return strings.Join([]string{"ProjectRsp", string(data)}, " ")
}
