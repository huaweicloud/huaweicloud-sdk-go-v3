package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainConfigs **参数解释：** 创建域名时同步配置的配置项，创建域名后可通过\"修改域名全量配置\"接口再次调整 **约束限制：** 不涉及
type CreateDomainConfigs struct {

	// **参数解释：** 将源站返回的错误状态码缓存到CDN节点，用户再次请求时由CDN直接响应给用户错误码 **约束限制：** 不涉及
	ErrorCodeCache *[]ErrorCodeCache `json:"error_code_cache,omitempty"`

	IpFilter *IpFilter `json:"ip_filter,omitempty"`

	// **参数解释：** 设置用量封顶阈值，当实际用量大于阈值时停用域名，有效预防流量盗刷或恶意攻击带来的高额账单。  > 由于监控数据存在时延，域名将在用量达到阈值后的10分钟左右被停用  **约束限制：** 不涉及
	FlowLimitStrategy *[]FlowLimitStrategy `json:"flow_limit_strategy,omitempty"`

	BandwidthAlarmStrategy *BandwidthAlarmStrategy `json:"bandwidth_alarm_strategy,omitempty"`
}

func (o CreateDomainConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainConfigs struct{}"
	}

	return strings.Join([]string{"CreateDomainConfigs", string(data)}, " ")
}
