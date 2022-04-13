package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPoolsRequest struct {
	// 分页查询中每页的监听器个数

	Limit *int32 `json:"limit,omitempty"`
	// 分页查询的起始的资源id，表示上一页最后一条查询记录的负载均衡器的id。不指定时表示查询第一页。

	Marker *string `json:"marker,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 后端云服务器组ID。

	Id *string `json:"id,omitempty"`
	// 后端云服务器组名称。

	Name *string `json:"name,omitempty"`
	// 后端云服务器组的描述信息。

	Description *string `json:"description,omitempty"`
	// 后端云服务器组关联的健康检查的ID。

	HealthmonitorId *string `json:"healthmonitor_id,omitempty"`
	// 后端云服务器组关联的负载均衡器ID。

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	// 后端云服务器组的后端协议。支持TCP、UDP和HTTP。

	Protocol *string `json:"protocol,omitempty"`
	// 后端云服务器组的负载均衡算法。取值范围：ROUND_ROBIN：加权轮询算法。LEAST_CONNECTIONS：加权最少连接算法。SOURCE_IP：源IP算法。当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。

	LbAlgorithm *string `json:"lb_algorithm,omitempty"`
	// 后端云服务器组关联的后端云服务器IP。

	MemberAddress *string `json:"member_address,omitempty"`
	// 后端云服务器组关联的后端云服务器对应的弹性云服务器的ID。

	MemberDeviceId *string `json:"member_device_id,omitempty"`
	// 企业项目ID，仅用于基于企业项目的细粒度鉴权使用；如果参数中传递了loadbalancer_id，则用该负载均衡器对应企业项目ID鉴权；如果参数中没有传递loadbalancer_id，而传递了healthmonitor_id，则使用健康检查器对应的企业项目id鉴权。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListPoolsRequest", string(data)}, " ")
}
