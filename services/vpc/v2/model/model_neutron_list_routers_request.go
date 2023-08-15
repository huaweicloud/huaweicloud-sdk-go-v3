package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronListRoutersRequest Request Object
type NeutronListRoutersRequest struct {

	// 每页返回的个数
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时查询第一页
	Marker *string `json:"marker,omitempty"`

	// 按照路由器的ID过滤查询
	Id *string `json:"id,omitempty"`

	// 按照路由器的状态过滤查询，取值范围：ACTIVE， DOWN，ERROR
	Status *string `json:"status,omitempty"`

	// 按照路由器所属的项目ID过滤查询
	TenantId *string `json:"tenant_id,omitempty"`

	// 按照路由器的管理状态过滤查询，取值范围：true or false
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

func (o NeutronListRoutersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListRoutersRequest struct{}"
	}

	return strings.Join([]string{"NeutronListRoutersRequest", string(data)}, " ")
}
