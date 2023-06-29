package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Propagation 传播详情
type Propagation struct {

	// 关联唯一标识
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 企业路由器id
	ErId *string `json:"er_id,omitempty"`

	// 路由表唯一标识
	RouteTableId *string `json:"route_table_id,omitempty"`

	// 连接唯一标识
	AttachmentId *string `json:"attachment_id,omitempty"`

	// 连接的类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 连接的资源唯一标识
	ResourceId *string `json:"resource_id,omitempty"`

	RoutePolicy *ImportRoutePolicy `json:"route_policy,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 资源创建时间  采用UTC时间  格式：YYYY-MM-DDTHH:MM:SS
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 资源更新时间  采用UTC时间  格式：YYYY-MM-DDTHH:MM:SS
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o Propagation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Propagation struct{}"
	}

	return strings.Join([]string{"Propagation", string(data)}, " ")
}
