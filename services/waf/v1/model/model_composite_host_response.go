package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompositeHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名id
	Hostid *string `json:"hostid,omitempty"`

	// 创建的云模式防护域名
	Hostname *string `json:"hostname,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// cname前缀
	AccessCode *string `json:"access_code,omitempty"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty"`

	// 创建防护域名的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 套餐付费模式，默认值为prePaid。prePaid：包周期款模式；postPaid：按需模式。
	PaidType *string `json:"paid_type,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 域名所属WAF模式,cloud为云模式，premium为独享模式
	WafType *string `json:"waf_type,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`

	// 接入进度，仅用于新版console(前端)使用
	AccessProgress *[]AccessProgress `json:"access_progress,omitempty"`

	// 租户引擎实例信息列表
	PremiumWafInstances *[]PremiumWafInstances `json:"premium_waf_instances,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 是否使用独享ip   - true：使用独享ip   - false：不实用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty"`

	// 华为云区域ID，控制台创建的域名会携带此参数，api调用创建的域名此参数为空，可以通过华为云上地区和终端节点文档查询区域ID对应的中文名称
	Region *string `json:"region,omitempty"`

	// 防护域名的源站服务器配置信息，只有独享模式域名才返回vpc_id
	Server *[]WafServer `json:"server,omitempty"`
}

func (o CompositeHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompositeHostResponse struct{}"
	}

	return strings.Join([]string{"CompositeHostResponse", string(data)}, " ")
}
