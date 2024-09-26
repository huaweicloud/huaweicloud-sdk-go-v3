package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigurationPort 访问方式配置端口、协议、证书等信息。
type AccessConfigurationPort struct {

	// 配置模式。 - 如果operator值为空，则表示使用全量覆盖模式进行配置，否则表示使用增删改模式进行配置。且此级列表的所有元素的operator值必须同时全为空或者非空。 - 当使用增删改模式时，operator取值支持\"add\",\"copy\",\"modify\",\"delete\"，分别表示新增，复制指定uid的元素修改后新增，修改指定uid的元素，删除指定uid的元素。 - 当operator取值为\"copy\",\"modify\",\"delete\"时，uid的值必须为非空，且存在于最后一次生效的配置中。 - 当operator取值为\"copy\",\"modify\"时，与operator同级别的字段中除uid外的所有字段如不写，置空或者为空列表，则表示保留在最后一次生效配置中指定uid的元素的同一字段的值。
	Operator *string `json:"operator,omitempty"`

	// 端口配置的uid。
	Uid *string `json:"uid,omitempty"`

	// 监听端口。
	TargetPort *int32 `json:"target_port,omitempty"`

	// 访问端口。
	Port *int32 `json:"port,omitempty"`

	// 协议，负载均衡支持TCP，负载均衡与路由配置支持HTTP、HTTPS。
	Protocol *AccessConfigurationPortProtocol `json:"protocol,omitempty"`

	// 默认证书，访问方式配置为转发策略且协议为HTTPS时配置，未配置域名证书对时使用默认证书。
	DefaultCertificate *string `json:"default_certificate,omitempty"`

	// 证书。
	Certificate *string `json:"certificate,omitempty"`

	// 安全策略。
	Policy *AccessConfigurationPortPolicy `json:"policy,omitempty"`

	Paths *[]AccessConfigurationHttpPath `json:"paths,omitempty"`

	// 用户选择的elb的ID。
	ElbId *string `json:"elb_id,omitempty"`
}

func (o AccessConfigurationPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigurationPort struct{}"
	}

	return strings.Join([]string{"AccessConfigurationPort", string(data)}, " ")
}

type AccessConfigurationPortProtocol struct {
	value string
}

type AccessConfigurationPortProtocolEnum struct {
	TCP   AccessConfigurationPortProtocol
	HTTP  AccessConfigurationPortProtocol
	HTTPS AccessConfigurationPortProtocol
}

func GetAccessConfigurationPortProtocolEnum() AccessConfigurationPortProtocolEnum {
	return AccessConfigurationPortProtocolEnum{
		TCP: AccessConfigurationPortProtocol{
			value: "TCP",
		},
		HTTP: AccessConfigurationPortProtocol{
			value: "HTTP",
		},
		HTTPS: AccessConfigurationPortProtocol{
			value: "HTTPS",
		},
	}
}

func (c AccessConfigurationPortProtocol) Value() string {
	return c.value
}

func (c AccessConfigurationPortProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigurationPortProtocol) UnmarshalJSON(b []byte) error {
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

type AccessConfigurationPortPolicy struct {
	value string
}

type AccessConfigurationPortPolicyEnum struct {
	TLS_1_2_STRICT AccessConfigurationPortPolicy
	TLS_1_2        AccessConfigurationPortPolicy
	TLS_1_1        AccessConfigurationPortPolicy
	TLS_1_0        AccessConfigurationPortPolicy
}

func GetAccessConfigurationPortPolicyEnum() AccessConfigurationPortPolicyEnum {
	return AccessConfigurationPortPolicyEnum{
		TLS_1_2_STRICT: AccessConfigurationPortPolicy{
			value: "tls-1-2-strict",
		},
		TLS_1_2: AccessConfigurationPortPolicy{
			value: "tls-1-2",
		},
		TLS_1_1: AccessConfigurationPortPolicy{
			value: "tls-1-1",
		},
		TLS_1_0: AccessConfigurationPortPolicy{
			value: "tls-1-0",
		},
	}
}

func (c AccessConfigurationPortPolicy) Value() string {
	return c.value
}

func (c AccessConfigurationPortPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigurationPortPolicy) UnmarshalJSON(b []byte) error {
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
