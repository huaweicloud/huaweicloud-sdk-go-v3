package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RaspProtectHistoryResponseInfo struct {

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 主机私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 告警时间(ms)
	AlarmTime *interface{} `json:"alarm_time,omitempty"`

	// 告警名称
	EventName *string `json:"event_name,omitempty"`

	// 告警级别
	Severity *string `json:"severity,omitempty"`

	// 源IP
	ReqSrcIp *string `json:"req_src_ip,omitempty"`

	// 应用程序调用堆栈信息
	AppStack *string `json:"app_stack,omitempty"`

	// 攻击附属字段
	AttackInputName *string `json:"attack_input_name,omitempty"`

	// 攻击负载内容
	AttackInputValue *string `json:"attack_input_value,omitempty"`

	// 查询字符串
	QueryString *string `json:"query_string,omitempty"`

	// web请求头信息
	ReqHeaders *string `json:"req_headers,omitempty"`

	// WEB请求方法
	ReqMethod *string `json:"req_method,omitempty"`

	// WEB请求参数
	ReqParams *string `json:"req_params,omitempty"`

	// WEB请求路径
	ReqPath *string `json:"req_path,omitempty"`

	// WEB请求协议
	ReqProtocol *string `json:"req_protocol,omitempty"`

	// WEB请求URL地址
	ReqUrl *string `json:"req_url,omitempty"`

	// 攻击标识
	AttackTag *string `json:"attack_tag,omitempty"`

	// 探针标识
	ChkProbe *string `json:"chk_probe,omitempty"`

	// 检测规则标识
	ChkRule *string `json:"chk_rule,omitempty"`

	// 规则描述
	ChkRuleDesc *string `json:"chk_rule_desc,omitempty"`

	// 应用是否存在bug
	ExistBug *string `json:"exist_bug,omitempty"`
}

func (o RaspProtectHistoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RaspProtectHistoryResponseInfo struct{}"
	}

	return strings.Join([]string{"RaspProtectHistoryResponseInfo", string(data)}, " ")
}
