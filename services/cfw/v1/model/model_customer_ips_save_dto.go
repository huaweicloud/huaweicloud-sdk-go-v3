package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CustomerIpsSaveDto struct {

	// **参数解释**： 动作 **取值范围**： 0：只记录日志，1：重置/拦截
	ActionType CustomerIpsSaveDtoActionType `json:"action_type"`

	// **参数解释**： 影响操作系统 **取值范围**： 0 any、1 Windows、2 Linux、3 FreeBSD、4 Solaris、5 other Unix、6 网络设备、7 Mac OS、8 ios、9 android、10 others
	AffectedOs CustomerIpsSaveDtoAffectedOs `json:"affected_os"`

	// **参数解释**： 攻击类型 **约束限制**： 不涉及 **取值范围**： 1：访问控制、2：漏洞扫描、3：邮件攻击、4：漏洞攻击、5：Web攻击、6：密码攻击、7：劫持攻击、8：协议异常、9：特洛伊木马、10：蠕虫、11：缓冲区溢出、12：黑客工具、13：间谍软件、14：DDos泛洪、15：应用层DDos攻击、16：其他可疑行为、17：可疑DNS活动、18：网络钓鱼、19：垃圾邮件、20：其他攻击 **默认取值**： 不涉及
	AttackType int32 `json:"attack_type"`

	// **参数解释**： 匹配IPS攻击的内容 **取值范围**：
	Contents []IpsContent `json:"contents"`

	// **参数解释**： 默认：null，0：客户端到服务端，1：服务端到客户端 **取值范围**： 不涉及
	Direction int32 `json:"direction"`

	DstPort *Port `json:"dst_port"`

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： ips规则名称 **取值范围**： 不涉及
	IpsName string `json:"ips_name"`

	// **参数解释**： 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志ID，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得 **约束限制**： type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得 **取值范围**： 32位UUID **默认取值**： 不涉及
	ObjectId string `json:"object_id"`

	// **参数解释**： 协议类型 **取值范围**： 1 FTP、2 TELNET、3 SMTP、4 DNS_TCP、5 DNS_UDP、6 DHCP、7 TFTP、8 FINGER、9 HTTP、10 POP3、11 SUNRPC_TCP、12 SUNRPC_UDP、13 NNTP、14 MSRPC_TCP、15 MSRPC_UDP、16 NETBIOS_NAME_TCP、17 NETBIOS_NAME_UDP、18 NETBIOS_SMB、19 NETBIOS_DATAGRAM、20 IMAP4、21 SNMP、22 LDAP、23 MSSQL、24 ORACLE
	Protocol int32 `json:"protocol"`

	// **参数解释**： 严重程度 **取值范围**： critical：致命，high：高危，medium:中危，low:低危
	Severity int32 `json:"severity"`

	// **参数解释**： 影响软件 **取值范围**： 0 ANY、1 ADOBE、2 APACHE、3 APPLE、4 CA、5 CISCO、6 GOOGLE_CHROME、7 HP、8 IBM、9 IE、10 IIS、11 MC_AFEE、12 MEDIA_PLAYER、13 MICROSOFT_NET、14 MICROSOFT_EDGE、15 MICROSOFT_EXCHANGE、16 MICROSOFT_OFFICE、17 MICROSOFT_OUTLOOK、18 MICROSOFT_SHARE_POINT、19 MICROSOFT_WINDOWS、20 MOZILLA、21 MSSQL、22 MYSQL、23 NOVELL、24 ORACLE、25 SAMBA、26 SAMSUNG、27 SAP、28 SCADA、29 SQUID、30 SUN、31 SYMANTEC、32 TREND_MICRO、33 VMWARE、34 WORD_PRESS、35 Others
	Software int32 `json:"software"`

	SrcPort *Port `json:"src_port"`
}

func (o CustomerIpsSaveDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerIpsSaveDto struct{}"
	}

	return strings.Join([]string{"CustomerIpsSaveDto", string(data)}, " ")
}

type CustomerIpsSaveDtoActionType struct {
	value int32
}

type CustomerIpsSaveDtoActionTypeEnum struct {
	E_0 CustomerIpsSaveDtoActionType
	E_1 CustomerIpsSaveDtoActionType
}

func GetCustomerIpsSaveDtoActionTypeEnum() CustomerIpsSaveDtoActionTypeEnum {
	return CustomerIpsSaveDtoActionTypeEnum{
		E_0: CustomerIpsSaveDtoActionType{
			value: 0,
		}, E_1: CustomerIpsSaveDtoActionType{
			value: 1,
		},
	}
}

func (c CustomerIpsSaveDtoActionType) Value() int32 {
	return c.value
}

func (c CustomerIpsSaveDtoActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomerIpsSaveDtoActionType) UnmarshalJSON(b []byte) error {
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

type CustomerIpsSaveDtoAffectedOs struct {
	value int32
}

type CustomerIpsSaveDtoAffectedOsEnum struct {
	E_0  CustomerIpsSaveDtoAffectedOs
	E_1  CustomerIpsSaveDtoAffectedOs
	E_2  CustomerIpsSaveDtoAffectedOs
	E_3  CustomerIpsSaveDtoAffectedOs
	E_4  CustomerIpsSaveDtoAffectedOs
	E_5  CustomerIpsSaveDtoAffectedOs
	E_6  CustomerIpsSaveDtoAffectedOs
	E_7  CustomerIpsSaveDtoAffectedOs
	E_8  CustomerIpsSaveDtoAffectedOs
	E_9  CustomerIpsSaveDtoAffectedOs
	E_10 CustomerIpsSaveDtoAffectedOs
}

func GetCustomerIpsSaveDtoAffectedOsEnum() CustomerIpsSaveDtoAffectedOsEnum {
	return CustomerIpsSaveDtoAffectedOsEnum{
		E_0: CustomerIpsSaveDtoAffectedOs{
			value: 0,
		}, E_1: CustomerIpsSaveDtoAffectedOs{
			value: 1,
		}, E_2: CustomerIpsSaveDtoAffectedOs{
			value: 2,
		}, E_3: CustomerIpsSaveDtoAffectedOs{
			value: 3,
		}, E_4: CustomerIpsSaveDtoAffectedOs{
			value: 4,
		}, E_5: CustomerIpsSaveDtoAffectedOs{
			value: 5,
		}, E_6: CustomerIpsSaveDtoAffectedOs{
			value: 6,
		}, E_7: CustomerIpsSaveDtoAffectedOs{
			value: 7,
		}, E_8: CustomerIpsSaveDtoAffectedOs{
			value: 8,
		}, E_9: CustomerIpsSaveDtoAffectedOs{
			value: 9,
		}, E_10: CustomerIpsSaveDtoAffectedOs{
			value: 10,
		},
	}
}

func (c CustomerIpsSaveDtoAffectedOs) Value() int32 {
	return c.value
}

func (c CustomerIpsSaveDtoAffectedOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomerIpsSaveDtoAffectedOs) UnmarshalJSON(b []byte) error {
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
