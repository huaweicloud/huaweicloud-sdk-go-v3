package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMembersRequest Request Object
type ListMembersRequest struct {

	// 后端云服务器组id
	PoolId string `json:"pool_id"`

	// 分页查询中每页的后端服务器个数
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始的资源id，表示上一页最后一条查询记录的后端服务器的id。不指定时表示查询第一页。
	Marker *string `json:"marker,omitempty"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 后端云服务器的ID。
	Id *string `json:"id,omitempty"`

	// 后端云服务器的名称。
	Name *string `json:"name,omitempty"`

	// 后端云服务器对应的IP地址。
	Address *string `json:"address,omitempty"`

	// 后端云服务器后端端口的协议号。
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// 后端云服务器所在的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 后端云服务器的管理状态。取值范围：true/false。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 后端云服务器的权重。
	Weight *int32 `json:"weight,omitempty"`

	// 企业项目ID。 不传时查询default企业项目（即enterprise_project_id=0）下的资源，鉴权按照default企业项目鉴权。 如果传值enterprise_project_id=all_granted_eps，则表示查询所有有权限的企业项目下的资源。 其他情况则传已存在的企业项目ID。此时会校验ID，若不存在或格式错误则报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMembersRequest struct{}"
	}

	return strings.Join([]string{"ListMembersRequest", string(data)}, " ")
}
