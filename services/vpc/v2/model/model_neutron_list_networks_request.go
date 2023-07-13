package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronListNetworksRequest Request Object
type NeutronListNetworksRequest struct {

	// 每页返回的个数
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时查询第一页
	Marker *string `json:"marker,omitempty"`

	// 按照网络对应的ID过滤查询
	Id *string `json:"id,omitempty"`

	// 按照网络的名称过滤查询
	Name *string `json:"name,omitempty"`

	// 按照网络的状态过滤查询，取值范围：ACTIVE、ERROR、DOWN
	Status *string `json:"status,omitempty"`

	// 按照网络是否支持跨租户共享过滤查询，取值范围：true or false
	Shared *bool `json:"shared,omitempty"`

	// 按照网络是否外部网络过滤查询，取值范围：true or false
	Routerexternal *bool `json:"router:external,omitempty"`

	// 按照网络的管理状态过滤查询，取值范围：true or false
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 按照网络的类型过滤查询
	ProvidernetworkType *string `json:"provider:network_type,omitempty"`

	// 按照network所属的项目ID过滤
	TenantId *string `json:"tenant_id,omitempty"`
}

func (o NeutronListNetworksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListNetworksRequest struct{}"
	}

	return strings.Join([]string{"NeutronListNetworksRequest", string(data)}, " ")
}
