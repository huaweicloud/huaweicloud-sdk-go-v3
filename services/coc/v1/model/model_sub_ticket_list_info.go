package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubTicketListInfo struct {

	// 企业项目ID。
	EpId *string `json:"ep_id,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源类型。
	Type *string `json:"type,omitempty"`

	// 资源名称。
	Name *string `json:"name,omitempty"`

	// 资源类型。
	Provider *string `json:"provider,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// Region信息。
	RegionId *string `json:"region_id,omitempty"`

	// 主机ID。
	HostingId *string `json:"hosting_id,omitempty"`

	// 资源属性信息。
	PropertiesJson *string `json:"properties_json,omitempty"`

	// 资源标签信息。
	TagsJson *string `json:"tags_json,omitempty"`

	// 是否已删除。
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 变更工单ID
	Id *string `json:"id,omitempty"`

	// 变更工单主键ID。
	MainTicketId *string `json:"main_ticket_id,omitempty"`

	// 父级工单ID。
	ParentTicketId *string `json:"parent_ticket_id,omitempty"`

	// 问题单ID。
	TicketId *string `json:"ticket_id,omitempty"`

	// 问题单单号。
	RealTicketId *string `json:"real_ticket_id,omitempty"`

	// 工单路径。
	TicketPath *string `json:"ticket_path,omitempty"`

	// region信息。
	TargetValue *string `json:"target_value,omitempty"`

	// 子单类型。
	TargetType *string `json:"target_type,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 创单人。
	Creator *string `json:"creator,omitempty"`

	// 操作人。
	Operator *string `json:"operator,omitempty"`
}

func (o SubTicketListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTicketListInfo struct{}"
	}

	return strings.Join([]string{"SubTicketListInfo", string(data)}, " ")
}
