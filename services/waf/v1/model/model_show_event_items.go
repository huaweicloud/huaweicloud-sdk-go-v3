package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 攻击类型
type ShowEventItems struct {

	// 攻击发生时的时间戳（毫秒）
	Time *int64 `json:"time,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 源ip
	Sip *string `json:"sip,omitempty"`

	// 域名
	Host *string `json:"host,omitempty"`

	// 攻击的url链接
	Url *string `json:"url,omitempty"`

	// 攻击类型
	Attack *string `json:"attack,omitempty"`

	// 命中的规则id
	Rule *string `json:"rule,omitempty"`

	// 防护动作
	Action *string `json:"action,omitempty"`

	// 攻击请求的cookie
	Cookie *string `json:"cookie,omitempty"`

	// 攻击请求的headers
	Headers *interface{} `json:"headers,omitempty"`

	// 被攻击的域名id
	HostId *string `json:"host_id,omitempty"`

	// 防护事件id
	Id *string `json:"id,omitempty"`

	// 恶意负载
	Payload *string `json:"payload,omitempty"`

	// 恶意负载位置
	PayloadLocation *string `json:"payload_location,omitempty"`

	// 源ip地理位置
	Region *string `json:"region,omitempty"`

	// 处理时长
	ProcessTime *int32 `json:"process_time,omitempty"`

	// 攻击请求的请求行
	RequestLine *string `json:"request_line,omitempty"`

	// 返回大小（字节）
	ResponseSize *int32 `json:"response_size,omitempty"`

	// 响应时间（毫秒）
	ResponseTime *int64 `json:"response_time,omitempty"`

	// 响应码
	Status *string `json:"status,omitempty"`

	// 请求体
	RequestBody *string `json:"request_body,omitempty"`
}

func (o ShowEventItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventItems struct{}"
	}

	return strings.Join([]string{"ShowEventItems", string(data)}, " ")
}
