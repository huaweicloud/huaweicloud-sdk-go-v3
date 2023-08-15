package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProtectedInstanceParams 查询保护实例数据结构
type ShowProtectedInstanceParams struct {

	// 保护实例的ID。
	Id string `json:"id"`

	// 保护实例的名称。
	Name string `json:"name"`

	// 保护实例的描述。
	Description string `json:"description"`

	// 保护实例的状态。
	Status string `json:"status"`

	// 生产站点云服务器ID。
	SourceServer string `json:"source_server"`

	// 容灾站点云服务器ID。
	TargetServer string `json:"target_server"`

	// 保护组的ID。
	ServerGroupId string `json:"server_group_id"`

	// 创建时间。默认格式为：\"yyyy-MM-dd HH:mm:ss.SSS\"，例：\"2019-04-01 12:00:00.000\"。
	CreatedAt string `json:"created_at"`

	// 更新时间。默认格式为：\"yyyy-MM-dd HH:mm:ss.SSS\"，例：\"2019-04-01 12:00:00.000\"。
	UpdatedAt string `json:"updated_at"`

	Metadata *MetadataParams `json:"metadata"`

	// 挂载的复制对列表。
	Attachment []ProtectedInstanceAttachment `json:"attachment"`

	// 标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 保护实例的同步进度。单位：百分比（%）。
	Progress int32 `json:"progress"`

	// 标识保护实例所在保护组的当前生产站点可用区。source：表示当前生产站点可用区为保护组source_availability_zone的值。target：表示当前生产站点可用区为保护组的target_availability_zone的值。
	PriorityStation string `json:"priority_station"`
}

func (o ShowProtectedInstanceParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectedInstanceParams struct{}"
	}

	return strings.Join([]string{"ShowProtectedInstanceParams", string(data)}, " ")
}
