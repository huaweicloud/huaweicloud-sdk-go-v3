package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RaspProtectHistoryResponseInfo struct {

	// **参数解释** 应用防护事件所属云服务器的名称，用于标识事件来源主机 **取值范围** 字符长度1-64位，支持中文、英文、数字、短横线、下划线，符合华为云ECS命名规范
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 应用防护事件所属云服务器的私有IP地址，用于定位事件来源主机的网络位置 **取值范围** 符合IPv4格式的字符串（如192.168.0.97），支持多个私有IP用逗号分隔
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释** 应用防护事件的发生时间，以Unix时间戳（毫秒级）表示 **时间格式** 可转换为YYYY-MM-DD HH:MM:SS格式（如1736414463000对应2024-12-10 10:41:03） **取值范围** Unix时间戳（毫秒级），取值0-为当前系统时间戳
	AlarmTime *int64 `json:"alarm_time,omitempty"`

	// **参数解释** 应用防护事件的具体名称，标识事件对应的攻击类型（如ExpressionInject表示表达式注入攻击） **取值范围** 字符长度1-128位，支持英文、数字、下划线，为系统预定义的攻击类型标识
	EventName *string `json:"event_name,omitempty"`

	// **参数解释** 应用防护事件的告警级别，用于筛选指定严重程度的事件 **约束限制** 取值必须在指定范围内，否则返回空结果 **取值范围** - Security：信息级 - Low：低危 - Medium：中危 - High：高危 - Critical：紧急 **默认取值** 无
	Severity *string `json:"severity,omitempty"`

	// **参数解释** 发起攻击的源IP地址，可能是公网IP或内网IP，用于定位攻击来源 **取值范围** 符合IPv4或IPv6格式的字符串，支持单个IP或IP段（如127.0.0.1、2001:db8::1）
	ReqSrcIp *string `json:"req_src_ip,omitempty"`

	// **参数解释** 应用防护事件发生时的应用程序调用堆栈信息，用于定位漏洞触发点 **取值范围** 字符长度0-4096位，支持英文、数字、符号等堆栈信息常见字符，为空表示无堆栈数据
	AppStack *string `json:"app_stack,omitempty"`

	// **参数解释** 攻击请求中的附属字段名称（如请求头字段、表单字段等），用于标识攻击载荷的传入字段 **取值范围** 字符长度0-256位，支持英文、数字、符号等HTTP请求字段常见字符，为空表示无相关字段
	AttackInputName *string `json:"attack_input_name,omitempty"`

	// **参数解释** 攻击请求中包含的恶意载荷数据（如注入脚本、恶意命令等），用于分析攻击手段 **取值范围** 字符长度0-2048位，支持各类字符（含特殊字符），为空表示无恶意载荷
	AttackInputValue *string `json:"attack_input_value,omitempty"`

	// **参数解释** 攻击请求的URL查询字符串部分（即?后的参数），用于分析攻击请求的参数传递方式 **取值范围** 字符长度0-1024位，支持URL编码后的字符，为空表示无查询字符串
	QueryString *string `json:"query_string,omitempty"`

	// **参数解释** 攻击请求的HTTP请求头信息，以JSON格式存储，包含User-Agent、Host等字段 **取值范围** 字符长度0-4096位，为JSON格式字符串，字段名和值支持常见HTTP头字符，为空表示无请求头信息
	ReqHeaders *string `json:"req_headers,omitempty"`

	// **参数解释** 攻击请求使用的HTTP方法（如GET、POST），用于分析攻击的请求类型 **取值范围** 字符长度3-10位，支持标准HTTP方法（GET、POST、PUT、DELETE等），区分大小写
	ReqMethod *string `json:"req_method,omitempty"`

	// **参数解释** 攻击请求的请求体参数（如POST请求的表单数据），用于分析攻击的参数传递内容 **取值范围** 字符长度0-2048位，支持表单编码或JSON格式字符，为空表示无请求体参数
	ReqParams *string `json:"req_params,omitempty"`

	// **参数解释** 攻击请求的URL路径部分（不含查询字符串），用于定位攻击的目标接口 **取值范围** 字符长度0-512位，支持URL路径字符（如/、字母、数字、短横线、下划线），为空表示根路径
	ReqPath *string `json:"req_path,omitempty"`

	// **参数解释** 攻击请求使用的HTTP协议版本（如HTTP/1.1），用于分析攻击的协议环境 **取值范围** 字符长度5-10位，支持HTTP/1.0、HTTP/1.1、HTTP/2等标准协议版本
	ReqProtocol *string `json:"req_protocol,omitempty"`

	// **参数解释** 攻击请求的完整URL地址（含协议、主机、路径、查询字符串），用于完整还原攻击请求 **取值范围** 字符长度0-1024位，符合URL格式规范，为空表示无完整URL信息
	ReqUrl *string `json:"req_url,omitempty"`

	// **参数解释** 应用防护事件的攻击类型标识，与请求参数的攻击标识对应（格式为小写下划线） **取值范围** - Attack Success：攻击成功 - Attack Attempt：攻击尝试 - Attack Blocked：攻击被阻断 - Abnormal Behavior：异常行为 - Collapsible Host：主机失陷 - System Vulnerability：系统脆弱性
	AttackTag *string `json:"attack_tag,omitempty"`

	// **参数解释**: 检测到该攻击事件的RASP探针标识，用于定位探针类型和检测模块 **取值范围**: 字符长度1-128位，支持英文、数字、点号、短横线、下划线，为系统预定义的探针标识
	ChkProbe *string `json:"chk_probe,omitempty"`

	// **参数解释** 触发该防护事件的检测规则唯一标识，用于关联具体的防护规则配置 **取值范围** 字符长度1-64位，支持英文、数字、下划线，为系统预定义的规则标识（如ExpressionInject）
	ChkRule *string `json:"chk_rule,omitempty"`

	// **参数解释** 触发该防护事件的检测规则详细描述，说明规则的检测逻辑和目的 **取值范围** 字符长度0-512位，支持中文、英文、数字、常用标点符号，为空表示无规则描述
	ChkRuleDesc *string `json:"chk_rule_desc,omitempty"`

	// **参数解释** 标识该防护事件是否因应用存在漏洞导致（yes表示存在漏洞，no表示不存在） **取值范围** - yes：存在漏洞 - no：不存在漏洞 - unknown：未知
	ExistBug *string `json:"exist_bug,omitempty"`
}

func (o RaspProtectHistoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RaspProtectHistoryResponseInfo struct{}"
	}

	return strings.Join([]string{"RaspProtectHistoryResponseInfo", string(data)}, " ")
}
