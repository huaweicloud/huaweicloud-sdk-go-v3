package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronRouter router对象模型
type NeutronRouter struct {

	// 功能说明：资源的管理状态。只支持true。 取值范围：true、false 约束：只支持true
	AdminStateUp bool `json:"admin_state_up"`

	ExternalGatewayInfo *ExternalGatewayInfo `json:"external_gateway_info"`

	// 路由器ID
	Id string `json:"id"`

	// 功能说明：路由器的名称 取值范围：0-64个字符，仅支持数字、字母、中文、_(下划线)、-（中划线）、.（点）。 约束：如果name非空，则name不能重复。
	Name string `json:"name"`

	// 功能说明：路由信息，见Route对象
	Routes []Route `json:"routes"`

	// 功能说明：路由器的状态 取值范围：ACTIVE， DOWN，ERROR
	Status string `json:"status"`

	// 项目ID
	TenantId string `json:"tenant_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 功能说明：资源创建UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 功能说明：资源更新UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o NeutronRouter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronRouter struct{}"
	}

	return strings.Join([]string{"NeutronRouter", string(data)}, " ")
}
