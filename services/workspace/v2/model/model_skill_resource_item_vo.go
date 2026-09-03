package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SkillResourceItemVo 技能绑定的资源列表项响应。
type SkillResourceItemVo struct {

	// 资源记录ID。
	Id *string `json:"id,omitempty"`

	// 资源类型，DESKTOP或DESKTOP_TAG。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源ID（DESKTOP时为instance_id，DESKTOP_TAG时为key:value）。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称（DESKTOP时为实例名称，DESKTOP_TAG时为标签）。
	ResourceName *string `json:"resource_name,omitempty"`

	// 技能版本号。
	SkillVersion *string `json:"skill_version,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o SkillResourceItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkillResourceItemVo struct{}"
	}

	return strings.Join([]string{"SkillResourceItemVo", string(data)}, " ")
}
