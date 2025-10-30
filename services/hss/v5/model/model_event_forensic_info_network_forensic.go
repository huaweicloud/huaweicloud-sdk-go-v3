package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfoNetworkForensic **参数解释**： 网络取证信息 **取值范围**： 不涉及
type EventForensicInfoNetworkForensic struct {

	// **参数解释**： 本地ip地址 **取值范围**： 不涉及
	LocalAddress *string `json:"local_address,omitempty"`

	// **参数解释**： 本地端口 **取值范围**： 不涉及
	LocalPort *int32 `json:"local_port,omitempty"`

	// **参数解释**： 源ip **取值范围**： 不涉及
	SrcIp *string `json:"src_ip,omitempty"`

	// **参数解释**： 远端ip地址(攻击者ip) **取值范围**： 不涉及
	RemoteAddress *string `json:"remote_address,omitempty"`

	// **参数解释**： 远程端口 **取值范围**： 不涉及
	RemotePort *int32 `json:"remote_port,omitempty"`

	// **参数解释**： 协议 **取值范围**： 不涉及
	Protocol *string `json:"protocol,omitempty"`

	// **参数解释**： 应用层协议 **取值范围**： 不涉及
	AppProtocol *string `json:"app_protocol,omitempty"`

	// **参数解释**： 流量方向 **取值范围**： 不涉及
	FlowDirection *string `json:"flow_direction,omitempty"`

	// **参数解释**： 连接次数 **取值范围**： 不涉及
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 首次连接时间(毫秒) **取值范围**： 不涉及
	FirstTime *int64 `json:"first_time,omitempty"`

	// **参数解释**： 最后一连接时间(毫秒) **取值范围**： 不涉及
	LastTime *int64 `json:"last_time,omitempty"`

	// **参数解释**： 请求方法 **取值范围**： 不涉及
	RequestMethod *string `json:"request_method,omitempty"`

	// **参数解释**： 请求地址 **取值范围**： 不涉及
	RequestUrl *string `json:"request_url,omitempty"`

	// **参数解释**： 查询字符串 **取值范围**： 不涉及
	QueryString *string `json:"query_string,omitempty"`

	// **参数解释**： 请求参数 **取值范围**： 不涉及
	RequestParams *string `json:"request_params,omitempty"`

	// **参数解释**： 请求头信息 **取值范围**： 不涉及
	RequestHeader *string `json:"request_header,omitempty"`
}

func (o EventForensicInfoNetworkForensic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfoNetworkForensic struct{}"
	}

	return strings.Join([]string{"EventForensicInfoNetworkForensic", string(data)}, " ")
}
