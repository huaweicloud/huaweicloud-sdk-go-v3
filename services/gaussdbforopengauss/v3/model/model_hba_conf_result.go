package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HbaConfResult struct {

	// **参数解释**: 客户端连接类型。 **取值范围**: - host：表示这条记录既接受一个普通的TCP/IP套接字连接，也接受一个经过SSL加密的TCP/IP套接字连接。 - hostssl：表示这条记录只接受一个经过SSL加密的TCP/IP套接字连接。 - hostnossl：表示这条记录只接受一个普通的TCP/IP套接字连接。
	Type *HbaConfResultType `json:"type,omitempty"`

	// **参数解释**: 声明记录所匹配且允许访问的数据库，多租特性场景下该参数声明记录所匹配且允许访问的PDB。 - 值replication表示如果请求一个复制连接，则匹配，但复制连接不表示任何特定的数据库。如需使用名为replication的数据库，需在database列使用记录“replication”作为数据库名。 - 多租数据库下， 值replication_pdb1表示如果请求一个名为pdb1数据库的复制连接，则匹配成功。值replication方式只生效Non-PDB。 - PDB复制连接生效配置方式为replication_[pdbname]，pdbname为用户创建PDB数据库时候的名字。 - 如需使用名为replication_pdb1的数据库，需在database列使用记录“replication_pdb1”作为数据库名。  **取值范围**: - all：表示该记录匹配所有数据库。 - 特定的数据库名称或者用逗号分隔的数据库列表。
	Database *string `json:"database,omitempty"`

	// **参数解释**: 声明记录所匹配且允许访问的数据库用户。 **取值范围**: - all：表明该记录匹配所有用户。 - 特定的数据库用户名或者用逗号分隔的用户列表。
	User *string `json:"user,omitempty"`

	// **参数解释**: 指定与记录匹配且允许访问的IP地址范围。 **取值范围**: 支持IPv4和IPv6，可以使用如下形式来表示： IP地址/掩码长度。例如，10.10.0.0/24、2001:250:250:250:250:250:250:175/128。
	Address *string `json:"address,omitempty"`

	// **参数解释**: 声明连接时使用的认证方法。 **取值范围**: - reject：无条件地拒绝连接。常用于过滤某些主机。 - md5：MD5加密算法安全性低，存在安全风险，不推荐使用，建议使用更安全的加密算法。默认不支持，可通过password_encryption_type参数配置。 - sha256：要求客户端提供一个sha256算法加密的口令进行认证，该口令在传送过程中结合salt（服务器发送给客户端的随机数）的单向sha256加密，增强了安全性。 - sm3：要求客户端提供一个sm3算法加密口令进行认证，该口令在传送过程中结合salt（服务器发送给客户端的随机数）的单向sm3加密，增加了安全性。 - cert：客户端证书认证模式，此模式需进行SSL连接配置且需要客户端提供有效的SSL证书，不需要提供用户密码。cert认证方式只支持hostssl类型的规则。
	Method *string `json:"method,omitempty"`
}

func (o HbaConfResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HbaConfResult struct{}"
	}

	return strings.Join([]string{"HbaConfResult", string(data)}, " ")
}

type HbaConfResultType struct {
	value string
}

type HbaConfResultTypeEnum struct {
	HOST      HbaConfResultType
	HOSTNOSSL HbaConfResultType
	HOSTSSL   HbaConfResultType
}

func GetHbaConfResultTypeEnum() HbaConfResultTypeEnum {
	return HbaConfResultTypeEnum{
		HOST: HbaConfResultType{
			value: "host",
		},
		HOSTNOSSL: HbaConfResultType{
			value: "hostnossl",
		},
		HOSTSSL: HbaConfResultType{
			value: "hostssl",
		},
	}
}

func (c HbaConfResultType) Value() string {
	return c.value
}

func (c HbaConfResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HbaConfResultType) UnmarshalJSON(b []byte) error {
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
