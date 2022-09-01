package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 攻击类型
type ShowEventItems struct {

	// 攻击发生时的时间戳（毫秒）
	Time *int64 `json:"time,omitempty" xml:"time"`

	// 策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 源ip，Web访问者的IP地址（攻击者IP地址）
	Sip *string `json:"sip,omitempty" xml:"sip"`

	// 域名
	Host *string `json:"host,omitempty" xml:"host"`

	// 攻击的url链接
	Url *string `json:"url,omitempty" xml:"url"`

	// 攻击类型
	Attack *string `json:"attack,omitempty" xml:"attack"`

	// 命中的规则id
	Rule *string `json:"rule,omitempty" xml:"rule"`

	// 防护动作
	Action *string `json:"action,omitempty" xml:"action"`

	// 攻击请求的cookie
	Cookie *string `json:"cookie,omitempty" xml:"cookie"`

	// 攻击请求的headers
	Headers *interface{} `json:"headers,omitempty" xml:"headers"`

	// 被攻击的域名id
	HostId *string `json:"host_id,omitempty" xml:"host_id"`

	// 防护事件id
	Id *string `json:"id,omitempty" xml:"id"`

	// 恶意负载
	Payload *string `json:"payload,omitempty" xml:"payload"`

	// 恶意负载位置
	PayloadLocation *string `json:"payload_location,omitempty" xml:"payload_location"`

	// 源ip地理位置
	Region *string `json:"region,omitempty" xml:"region"`

	// 处理时长
	ProcessTime *int32 `json:"process_time,omitempty" xml:"process_time"`

	// 攻击请求的请求行
	RequestLine *string `json:"request_line,omitempty" xml:"request_line"`

	// 返回大小（字节）
	ResponseSize *int32 `json:"response_size,omitempty" xml:"response_size"`

	// 响应时间（毫秒）
	ResponseTime *int64 `json:"response_time,omitempty" xml:"response_time"`

	// 响应码
	Status *string `json:"status,omitempty" xml:"status"`

	// 请求体
	RequestBody *string `json:"request_body,omitempty" xml:"request_body"`
}

func (o ShowEventItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventItems struct{}"
	}

	return strings.Join([]string{"ShowEventItems", string(data)}, " ")
}
