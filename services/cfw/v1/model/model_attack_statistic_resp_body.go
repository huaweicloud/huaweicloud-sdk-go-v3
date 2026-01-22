package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttackStatisticRespBody struct {

	// **参数解释**： 应用列表 **取值范围**： 不涉及
	Apps *[]TopInfo `json:"apps,omitempty"`

	// **参数解释**： 绑定资源名称 **取值范围**： 不涉及
	AssociatedName *string `json:"associated_name,omitempty"`

	// **参数解释**： 绑定资源类型 **取值范围**： PORT：IPV4云服务器 IPV6_PORT：IPV6云服务器
	AssociatedType *string `json:"associated_type,omitempty"`

	// **参数解释**： 攻击次数 **取值范围**： 不涉及
	AttackCount *int64 `json:"attack_count,omitempty"`

	// **参数解释**： 攻击类型 **取值范围**： Access Control：访问控制 Vulnerability scanning：漏洞扫描 Email attack：邮件攻击 Vulnerability Attack：漏洞攻击 Web attack：Web攻击 password attack：密码攻击 Hijacking attack：劫持攻击 Protocol exception：协议异常 Trojan horse：特洛伊木马
	AttackType *string `json:"attack_type,omitempty"`

	// **参数解释**： 拦截次数 **取值范围**： 不涉及
	DenyCount *int64 `json:"deny_count,omitempty"`

	// **参数解释**： 目的端口列表 **取值范围**： 不涉及
	DstPorts *[]TopInfo `json:"dst_ports,omitempty"`

	// **参数解释**： IP地址 **取值范围**： 不涉及
	Ip *string `json:"ip,omitempty"`

	// **参数解释**： 最近攻击时间 **取值范围**： 不涉及
	LatestTime *int64 `json:"latest_time,omitempty"`

	// **参数解释**： 地区ID **取值范围**： 不涉及
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释**： 地区名称 **取值范围**： 不涉及
	RegionName *string `json:"region_name,omitempty"`

	// **参数解释**： 攻击来源类型 **取值范围**： 不涉及
	SrcType *string `json:"src_type,omitempty"`

	// **参数解释**： vgw Id **取值范围**： 不涉及
	VgwId *string `json:"vgw_id,omitempty"`
}

func (o AttackStatisticRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackStatisticRespBody struct{}"
	}

	return strings.Join([]string{"AttackStatisticRespBody", string(data)}, " ")
}
