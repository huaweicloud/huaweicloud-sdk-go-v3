package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPoolsRequest struct {
	// 后端云服务器组的管理状态；该字段为预留字段，暂未启用。只支持更新为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 后端云服务器组的描述信息。

	Description *[]string `json:"description,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 后端云服务器组关联的健康检查的ID。

	HealthmonitorId *[]string `json:"healthmonitor_id,omitempty"`
	// 后端云服务器组的ID。

	Id *[]string `json:"id,omitempty"`
	// 后端云服务器组所在ip_version。 取值范围：dualstack,v4,v6。

	IpVersion *[]string `json:"ip_version,omitempty"`
	// 后端云服务器组的负载均衡算法，取值：ROUND_ROBIN：加权轮询算法；LEAST_CONNECTIONS：加权最少连接算法；SOURCE_IP：源IP算法；当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。

	LbAlgorithm *[]string `json:"lb_algorithm,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 后端云服务器组绑定的负载均衡器ID。

	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 后端云服务器的IP地址。

	MemberAddress *[]string `json:"member_address,omitempty"`
	// 是否开启删除保护，false不开启，默认为空查询全部

	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`
	// 后端云服务器对应的弹性云服务器的ID。

	MemberDeviceId *[]string `json:"member_device_id,omitempty"`
	// 后端云服务器组的名称。

	Name *[]string `json:"name,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 后端云服务器组的后端协议。

	Protocol *[]string `json:"protocol,omitempty"`
}

func (o ListPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListPoolsRequest", string(data)}, " ")
}
