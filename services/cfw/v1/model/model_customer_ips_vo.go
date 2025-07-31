package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CustomerIpsVo **参数解释**： 自定义IPS规则返回数据 **取值范围**： 不涉及
type CustomerIpsVo struct {

	// **参数解释**： 动作 **取值范围**： 0：只记录日志，1：重置/拦截
	Action *CustomerIpsVoAction `json:"action,omitempty"`

	// **参数解释**： 影响操作系统 **取值范围**： 0 any、1 Windows、2 Linux、3 FreeBSD、4 Solaris、5 other Unix、6 网络设备、7 Mac OS、8 ios、9 android、10 others
	AffectedOs *CustomerIpsVoAffectedOs `json:"affected_os,omitempty"`

	// **参数解释**： 攻击类型 **取值范围**： 1：访问控制、2：漏洞扫描、3：邮件攻击、4：漏洞攻击、5：Web攻击、6：密码攻击、7：劫持攻击、8：协议异常、9：特洛伊木马、10：蠕虫、11：缓冲区溢出、12：黑客工具、13：间谍软件、14：DDos泛洪、15：应用层DDos攻击、16：其他可疑行为、17：可疑DNS活动、18：网络钓鱼、19：垃圾邮件、20：其他攻击
	AttackType *int32 `json:"attack_type,omitempty"`

	// **参数解释**： 规则状态 **取值范围**： 0：初始化，1：配置中，2：配置成功，3：配置失败
	ConfigStatus *int32 `json:"config_status,omitempty"`

	// **参数解释**： 匹配IPS攻击的内容 **取值范围**：
	Contents *[]IpsContent `json:"contents,omitempty"`

	// **参数解释**： 默认：null，0：客户端到服务端，1：服务端到客户端 **取值范围**： 不涉及
	Direction *int32 `json:"direction,omitempty"`

	DstPort *Port `json:"dst_port,omitempty"`

	// **参数解释**： 防火墙集群id **取值范围**： 不涉及
	GroupId *string `json:"group_id,omitempty"`

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

	SrcPort *Port `json:"src_port,omitempty"`
}

func (o CustomerIpsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerIpsVo struct{}"
	}

	return strings.Join([]string{"CustomerIpsVo", string(data)}, " ")
}

type CustomerIpsVoAction struct {
	value int32
}

type CustomerIpsVoActionEnum struct {
	E_0 CustomerIpsVoAction
	E_1 CustomerIpsVoAction
}

func GetCustomerIpsVoActionEnum() CustomerIpsVoActionEnum {
	return CustomerIpsVoActionEnum{
		E_0: CustomerIpsVoAction{
			value: 0,
		}, E_1: CustomerIpsVoAction{
			value: 1,
		},
	}
}

func (c CustomerIpsVoAction) Value() int32 {
	return c.value
}

func (c CustomerIpsVoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomerIpsVoAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CustomerIpsVoAffectedOs struct {
	value int32
}

type CustomerIpsVoAffectedOsEnum struct {
	E_0  CustomerIpsVoAffectedOs
	E_1  CustomerIpsVoAffectedOs
	E_2  CustomerIpsVoAffectedOs
	E_3  CustomerIpsVoAffectedOs
	E_4  CustomerIpsVoAffectedOs
	E_5  CustomerIpsVoAffectedOs
	E_6  CustomerIpsVoAffectedOs
	E_7  CustomerIpsVoAffectedOs
	E_8  CustomerIpsVoAffectedOs
	E_9  CustomerIpsVoAffectedOs
	E_10 CustomerIpsVoAffectedOs
}

func GetCustomerIpsVoAffectedOsEnum() CustomerIpsVoAffectedOsEnum {
	return CustomerIpsVoAffectedOsEnum{
		E_0: CustomerIpsVoAffectedOs{
			value: 0,
		}, E_1: CustomerIpsVoAffectedOs{
			value: 1,
		}, E_2: CustomerIpsVoAffectedOs{
			value: 2,
		}, E_3: CustomerIpsVoAffectedOs{
			value: 3,
		}, E_4: CustomerIpsVoAffectedOs{
			value: 4,
		}, E_5: CustomerIpsVoAffectedOs{
			value: 5,
		}, E_6: CustomerIpsVoAffectedOs{
			value: 6,
		}, E_7: CustomerIpsVoAffectedOs{
			value: 7,
		}, E_8: CustomerIpsVoAffectedOs{
			value: 8,
		}, E_9: CustomerIpsVoAffectedOs{
			value: 9,
		}, E_10: CustomerIpsVoAffectedOs{
			value: 10,
		},
	}
}

func (c CustomerIpsVoAffectedOs) Value() int32 {
	return c.value
}

func (c CustomerIpsVoAffectedOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomerIpsVoAffectedOs) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
