package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecordSet DNS记录集合。
type RecordSet struct {

	// 记录集id
	Id string `json:"id"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 可用区id
	ZoneId *string `json:"zone_id,omitempty"`

	// 可用区名称
	ZoneName *string `json:"zone_name,omitempty"`

	// 记录类型
	Type *string `json:"type,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 记录集名称
	Name string `json:"name"`

	// 生存时间
	Ttl int32 `json:"ttl"`

	// 记录信息
	Records []string `json:"records"`

	// 比例
	Weight int32 `json:"weight"`

	// 解析线路ID
	Line string `json:"line"`

	// 创建时间
	CreateAt *string `json:"create_at,omitempty"`

	// 更新时间
	UpdateAt *string `json:"update_at,omitempty"`

	// 是否为默认记录集合
	Default *bool `json:"default,omitempty"`
}

func (o RecordSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordSet struct{}"
	}

	return strings.Join([]string{"RecordSet", string(data)}, " ")
}
