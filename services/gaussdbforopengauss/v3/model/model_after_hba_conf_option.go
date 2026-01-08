package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AfterHbaConfOption **参数解释**: 修改后的hba配置信息。 **约束限制**: 不涉及。
type AfterHbaConfOption struct {

	// **参数解释** 客户端连接类型。 **约束限制**: 不涉及。 **取值范围** - host：表示这条记录既接受一个普通的TCP/IP套接字连接，也接受一个经过SSL加密的TCP/IP套接字连接。 - hostssl：表示这条记录只接受一个经过SSL加密的TCP/IP套接字连接。 - hostnossl：表示这条记录只接受一个普通的TCP/IP套接字连接。  **默认取值**: 不涉及。
	Type AfterHbaConfOptionType `json:"type"`

	// **参数解释** 声明记录所匹配且允许访问的数据库，多租特性场景下该参数声明记录所匹配且允许访问的PDB。 - 值replication表示如果请求一个复制连接，则匹配，但复制连接不表示任何特定的数据库。如需使用名为replication的数据库，需在database列使用记录“replication”作为数据库名。 - 多租数据库下， 值replication_pdb1表示如果请求一个名为pdb1数据库的复制连接，则匹配成功。值replication方式只生效Non-PDB。 - PDB复制连接生效配置方式为replication_[pdbname]，pdbname为用户创建PDB数据库时候的名字。 - 如需使用名为replication_pdb1的数据库，需在database列使用记录“replication_pdb1”作为数据库名。  **约束限制**: 不涉及。 **取值范围** - all：表示该记录匹配所有数据库。 - 特定的数据库名称或者用逗号分隔的数据库列表。  **默认取值**: 不涉及。
	Database string `json:"database"`

	// **参数解释** 声明记录所匹配且允许访问的数据库用户。 **约束限制** 不支持系统用户。 **取值范围** - all：表明该记录匹配所有用户。 - 特定的数据库用户名或者用逗号分隔的用户列表。 **默认取值**: 不涉及。
	User string `json:"user"`

	// **参数解释** 指定与记录匹配且允许访问的IP地址范围。 **约束限制** 当前仅支持IP地址/掩码长度格式。 数据库引擎版本为V2.0-8.1.0及以上支持address配置IPv6的IP。 **取值范围** 支持IPv4和IPv6，可以使用如下形式来表示： IP地址/掩码长度。例如，10.10.0.0/24、2001:250:250:250:250:250:250:175/128。 **默认取值**: 不涉及。
	Address string `json:"address"`

	// **参数解释** 声明连接时使用的认证方法。 **约束限制**: 不涉及。 **取值范围** - reject：无条件地拒绝连接。常用于过滤某些主机。 - md5：MD5加密算法安全性低，存在安全风险，不推荐使用，建议使用更安全的加密算法。默认不支持，可通过password_encryption_type参数配置。 - sha256：要求客户端提供一个sha256算法加密的口令进行认证，该口令在传送过程中结合salt（服务器发送给客户端的随机数）的单向sha256加密，增强了安全性。 - sm3：要求客户端提供一个sm3算法加密口令进行认证，该口令在传送过程中结合salt（服务器发送给客户端的随机数）的单向sm3加密，增加了安全性。 - cert：客户端证书认证模式，此模式需进行SSL连接配置且需要客户端提供有效的SSL证书，不需要提供用户密码。cert认证方式只支持hostssl类型的规则。  **默认取值**: 不涉及。
	Method AfterHbaConfOptionMethod `json:"method"`
}

func (o AfterHbaConfOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AfterHbaConfOption struct{}"
	}

	return strings.Join([]string{"AfterHbaConfOption", string(data)}, " ")
}

type AfterHbaConfOptionType struct {
	value string
}

type AfterHbaConfOptionTypeEnum struct {
	HOST      AfterHbaConfOptionType
	HOSTNOSSL AfterHbaConfOptionType
	HOSTSSL   AfterHbaConfOptionType
}

func GetAfterHbaConfOptionTypeEnum() AfterHbaConfOptionTypeEnum {
	return AfterHbaConfOptionTypeEnum{
		HOST: AfterHbaConfOptionType{
			value: "host",
		},
		HOSTNOSSL: AfterHbaConfOptionType{
			value: "hostnossl",
		},
		HOSTSSL: AfterHbaConfOptionType{
			value: "hostssl",
		},
	}
}

func (c AfterHbaConfOptionType) Value() string {
	return c.value
}

func (c AfterHbaConfOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AfterHbaConfOptionType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type AfterHbaConfOptionMethod struct {
	value string
}

type AfterHbaConfOptionMethodEnum struct {
	REJECT AfterHbaConfOptionMethod
	MD5    AfterHbaConfOptionMethod
	SHA256 AfterHbaConfOptionMethod
	SM3    AfterHbaConfOptionMethod
	CERT   AfterHbaConfOptionMethod
}

func GetAfterHbaConfOptionMethodEnum() AfterHbaConfOptionMethodEnum {
	return AfterHbaConfOptionMethodEnum{
		REJECT: AfterHbaConfOptionMethod{
			value: "reject",
		},
		MD5: AfterHbaConfOptionMethod{
			value: "md5",
		},
		SHA256: AfterHbaConfOptionMethod{
			value: "sha256",
		},
		SM3: AfterHbaConfOptionMethod{
			value: "sm3",
		},
		CERT: AfterHbaConfOptionMethod{
			value: "cert",
		},
	}
}

func (c AfterHbaConfOptionMethod) Value() string {
	return c.value
}

func (c AfterHbaConfOptionMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AfterHbaConfOptionMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
