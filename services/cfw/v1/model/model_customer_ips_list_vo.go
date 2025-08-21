package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerIpsListVo struct {

	// **参数解释**： 动作 **取值范围**： 0：只记录日志，1：重置/拦截
	Action *int32 `json:"action,omitempty"`

	// **参数解释**： 影响操作系统 **取值范围**： 0 any、1 Windows、2 Linux、3 FreeBSD、4 Solaris、5 other Unix、6 网络设备、7 Mac OS、8 ios、9 android、10 others
	AffectedOs *int32 `json:"affected_os,omitempty"`

	// **参数解释**： 攻击类型 **约束限制**： 不涉及 **取值范围**： 1：访问控制、2：漏洞扫描、3：邮件攻击、4：漏洞攻击、5：Web攻击、6：密码攻击、7：劫持攻击、8：协议异常、9：特洛伊木马、10：蠕虫、11：缓冲区溢出、12：黑客工具、13：间谍软件、14：DDos泛洪、15：应用层DDos攻击、16：其他可疑行为、17：可疑DNS活动、18：网络钓鱼、19：垃圾邮件、20：其他攻击 **默认取值**： 不涉及
	AttackType *int32 `json:"attack_type,omitempty"`

	// **参数解释**： 规则状态 **取值范围**： 0：初始化，1：配置中，2：配置成功，3：配置失败
	ConfigStatus *int32 `json:"config_status,omitempty"`

	// **参数解释**： 匹配IPS攻击的内容 **取值范围**： 不涉及
	Content *string `json:"content,omitempty"`

	// **参数解释**： 端口类型 **取值范围**： -1 Any，0 包含，1 排除
	DstPortType *int32 `json:"dst_port_type,omitempty"`

	// **参数解释**： 端口 **取值范围**： 1 - 65535
	DstPorts *string `json:"dst_ports,omitempty"`

	// **参数解释**： cfw侧自定义IPS规则id **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	IpsCfwId *string `json:"ips_cfw_id,omitempty"`

	// **参数解释**： ips规则id **取值范围**： 不涉及
	IpsId *string `json:"ips_id,omitempty"`

	// **参数解释**： ips规则名称 **取值范围**： 不涉及
	IpsName *string `json:"ips_name,omitempty"`

	// **参数解释**： 协议类型 **取值范围**： 1 FTP、2 TELNET、3 SMTP、4 DNS_TCP、5 DNS_UDP、6 DHCP、7 TFTP、8 FINGER、9 HTTP、10 POP3、11 SUNRPC_TCP、12 SUNRPC_UDP、13 NNTP、14 MSRPC_TCP、15 MSRPC_UDP、16 NETBIOS_NAME_TCP、17 NETBIOS_NAME_UDP、18 NETBIOS_SMB、19 NETBIOS_DATAGRAM、20 IMAP4、21 SNMP、22 LDAP、23 MSSQL、24 ORACLE
	Protocol *int32 `json:"protocol,omitempty"`

	// **参数解释**： 严重程度 **取值范围**： critical：致命，high：高危，medium:中危，low:低危
	Severity *int32 `json:"severity,omitempty"`

	// **参数解释**： 影响软件 **取值范围**： 0 ANY、1 ADOBE、2 APACHE、3 APPLE、4 CA、5 CISCO、6 GOOGLE_CHROME、7 HP、8 IBM、9 IE、10 IIS、11 MC_AFEE、12 MEDIA_PLAYER、13 MICROSOFT_NET、14 MICROSOFT_EDGE、15 MICROSOFT_EXCHANGE、16 MICROSOFT_OFFICE、17 MICROSOFT_OUTLOOK、18 MICROSOFT_SHARE_POINT、19 MICROSOFT_WINDOWS、20 MOZILLA、21 MSSQL、22 MYSQL、23 NOVELL、24 ORACLE、25 SAMBA、26 SAMSUNG、27 SAP、28 SCADA、29 SQUID、30 SUN、31 SYMANTEC、32 TREND_MICRO、33 VMWARE、34 WORD_PRESS、35 Others
	Software *int32 `json:"software,omitempty"`

	// **参数解释**： 端口类型，-1 Any，0 包含，1 排除 **取值范围**： 不涉及
	SrcPortType *int32 `json:"src_port_type,omitempty"`

	// **参数解释**： 端口 **取值范围**： 不涉及
	SrcPorts *string `json:"src_ports,omitempty"`

	// **参数解释**： 集群ID **取值范围**： 不涉及
	GroupId *string `json:"group_id,omitempty"`
}

func (o CustomerIpsListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerIpsListVo struct{}"
	}

	return strings.Join([]string{"CustomerIpsListVo", string(data)}, " ")
}
