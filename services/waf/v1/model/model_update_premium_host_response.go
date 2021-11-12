package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePremiumHostResponse struct {
	// 域名id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 创建的云模式防护域名

	Hostname *string `json:"hostname,omitempty"`
	// 用户Domain ID

	Domainid *string `json:"domainid,omitempty"`
	// 用户的project_id

	ProjectId *string `json:"project_id,omitempty"`
	// cname前缀

	AccessCode *string `json:"access_code,omitempty"`
	// http协议类型

	Protocol *string `json:"protocol,omitempty"`
	// 源站信息

	Server *[]PremiumWafServer `json:"server,omitempty"`
	// 返回的证书id

	Certificateid *string `json:"certificateid,omitempty"`
	// 证书名称

	Certificatename *string `json:"certificatename,omitempty"`
	// 支持最低的TLS版本

	Tls *string `json:"tls,omitempty"`
	// 加密套件代码

	Cipher *string `json:"cipher,omitempty"`
	// 是否开启了代理

	Proxy *bool `json:"proxy,omitempty"`
	// 锁定状态

	Locked *int32 `json:"locked,omitempty"`
	// 防护状态

	ProtectStatus *int32 `json:"protect_status,omitempty"`
	// 接入状态

	AccessStatus *int32 `json:"access_status,omitempty"`
	// 创建防护域名的时间

	Timestamp *int64 `json:"timestamp,omitempty"`

	BlockPage *BlockPage `json:"block_page,omitempty"`
	// 可扩展属性

	Extend map[string]string `json:"extend,omitempty"`

	TrafficMark *TrafficMark `json:"traffic_mark,omitempty"`
	// 域名特殊标记

	Flag map[string]string `json:"flag,omitempty"`
	// 独享模式特殊域名模式（仅特殊模式需要，如elb）

	Mode *string `json:"mode,omitempty"`
	// 域名关联的组ID（仅特殊模式需要，如elb）

	PoolIds        *[]string `json:"pool_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdatePremiumHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostResponse struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostResponse", string(data)}, " ")
}
