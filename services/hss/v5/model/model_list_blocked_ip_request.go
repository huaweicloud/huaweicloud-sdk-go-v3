package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBlockedIpRequest Request Object
type ListBlockedIpRequest struct {

	// Region ID
	Region *string `json:"region,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 查询时间范围天数，与自定义查询时间begin_time，end_time互斥
	LastDays *int32 `json:"last_days,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 攻击源IP
	SrcIp *string `json:"src_ip,omitempty"`

	// 拦截状态，包含如下:   - intercepted : 已拦截   - canceled : 已解除拦截   - cancelling : 待解除拦截
	InterceptStatus *string `json:"intercept_status,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListBlockedIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlockedIpRequest struct{}"
	}

	return strings.Join([]string{"ListBlockedIpRequest", string(data)}, " ")
}
