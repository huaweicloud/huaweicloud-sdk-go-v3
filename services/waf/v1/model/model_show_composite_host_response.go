package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCompositeHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty" xml:"id"`

	// 域名id
	Hostid *string `json:"hostid,omitempty" xml:"hostid"`

	// 创建的云模式防护域名
	Hostname *string `json:"hostname,omitempty" xml:"hostname"`

	// 策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// cname前缀
	AccessCode *string `json:"access_code,omitempty" xml:"access_code"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty" xml:"protect_status"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty" xml:"access_status"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty" xml:"proxy"`

	// 创建防护域名的时间
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 套餐付费模式，目前只支持prePaid预付款模式
	PaidType *string `json:"paid_type,omitempty" xml:"paid_type"`

	Flag *Flag `json:"flag,omitempty" xml:"flag"`

	// 域名所属WAF模式,cloud为云模式，premium为独享模式
	WafType *string `json:"waf_type,omitempty" xml:"waf_type"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty" xml:"web_tag"`

	// 接入进度，仅用于新版console(前端)使用
	AccessProgress *[]AccessProgress `json:"access_progress,omitempty" xml:"access_progress"`

	// 租户引擎实例信息列表
	PremiumWafInstances *[]PremiumWafInstances `json:"premium_waf_instances,omitempty" xml:"premium_waf_instances"`

	// 域名描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 是否使用独享ip   - true：使用独享ip   - false：不实用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty" xml:"exclusive_ip"`

	// 华为云区域ID，控制台创建的域名会携带此参数，api调用创建的域名此参数为空，可以通过华为云上地区和终端节点文档查询区域ID对应的中文名称
	Region         *string `json:"region,omitempty" xml:"region"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCompositeHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompositeHostResponse struct{}"
	}

	return strings.Join([]string{"ShowCompositeHostResponse", string(data)}, " ")
}
