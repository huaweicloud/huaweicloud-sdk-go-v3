package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ServerCreateRequest struct {

	// **参数解释**：用于登录服务器的密码。admin_pass和key_pair_name必须二选一。密码规则： - 长度为8至26个字符 - 至少包含大写字母、小写字母、数字及特殊符号(!@%-_=+[{}]:,./?)中的3种 - 不能与用户名或倒序的用户名相同 - 不能包含root或administrator及其逆序 **约束限制**：admin_pass和key_pair_name不能同时存在。 **取值范围**：长度为8至26个字符，满足上述密码规则。 **默认取值**：不涉及。
	AdminPass *string `json:"admin_pass,omitempty"`

	// **参数解释**：服务器规格架构类型。 **约束限制**：不涉及。 **取值范围**： - X86：CPU架构为X86 - ARM：CPU架构为ARM **默认取值**：不涉及。
	Arch *ServerCreateRequestArch `json:"arch,omitempty"`

	// **参数解释**：服务器所在的可用区。 **约束限制**：不涉及。 **取值范围**：长度为1至256个字符，只能包含字母、数字、中划线、下划线和点。 **默认取值**：不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	ChargingInfo *ChargingInfo `json:"charging_info,omitempty"`

	// **参数解释**：单次购买的服务器数量。 **约束限制**：不支持超节点。 **取值范围**：1至10。 **默认取值**：不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：企业ID。 **约束限制**：不涉及。 **取值范围**：长度为1至36个字符，只能包含字母、数字、中划线、下划线和点。 **默认取值**：不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：服务器规格名称。 **约束限制**：flavor和resource_flavor二选一。 **取值范围**：长度为1至128个字符。 **默认取值**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：服务器资源规格名称。 **约束限制**：flavor和resource_flavor二选一。 **取值范围**：长度为1至256个字符，只能包含字母、数字、中划线、下划线和点。 **默认取值**：不涉及。
	ResourceFlavor *string `json:"resource_flavor,omitempty"`

	// **参数解释**：服务器镜像ID。 **约束限制**：不涉及。 **取值范围**：长度为36个字符，符合UUID格式。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`

	// **参数解释**：服务器登录密钥对名称。admin_pass和key_pair_name必须二选一。注意超节点仅支持使用密钥对创建。 **约束限制**：admin_pass和key_pair_name不能同时存在。 **取值范围**：长度为1至64个字符，只能包含字母、数字、中划线、下划线和点。 **默认取值**：不涉及。
	KeyPairName *string `json:"key_pair_name,omitempty"`

	// **参数解释**：服务器名称。 **约束限制**：不涉及。 **取值范围**：长度为1至64个字符，只能包含字母、数字、中划线、下划线和点。 **默认取值**：不涉及。
	Name string `json:"name"`

	Network *ServerNetwork `json:"network"`

	RootVolume *EvsVolume `json:"root_volume,omitempty"`

	DataVolume *ServerDataVolume `json:"data_volume,omitempty"`

	// **参数解释**：服务器类型。 **约束限制**：不涉及。 **取值范围**： - BMS：裸金属服务 - ECS：弹性云服务 - HPS：超节点服务 **默认取值**：不涉及。
	ServerType *ServerCreateRequestServerType `json:"server_type,omitempty"`

	// **参数解释**： 创建云服务器过程中待注入实例自定义数据。支持注入文本、文本文件。 示例： base64编码前： * Linux服务器： ```bash #!/bin/bash echo user_test > /home/user.txt ``` base64编码后： * Linux服务器： ```bash IyEvYmluL2Jhc2gKZWNobyB1c2VyX3Rlc3QgPiAvaG9tZS91c2VyLnR4dA== ``` 了解更多实例自定义数据注入请参考[[用户数据注入](https://support.huaweicloud.com/usermanual-ecs/zh-cn_topic_0032380449.html)](tag:hc)[[用户数据注入](https://support.huaweicloud.com/intl/zh-cn/usermanual-ecs/zh-cn_topic_0032380449.html)](tag:hk)[ECS服务“通过实例自定义数据配置ECS实例”章节](tag:fcs,hcso)。 用户需明确user_data的使用效果，可能产生的影响和风险由用户自行承担。 **约束限制**： - user_data的值为base64编码之后的内容。 - 注入内容（编码之前的内容）最大长度为32K。  **取值范围**：不涉及。 **默认取值**：不涉及。
	UserData *string `json:"user_data,omitempty"`

	// **参数解释**：超节点集群网络信息。仅在创建超节点时需要该参数。 **约束限制**：仅用于创建HPS类型的服务器。 **取值范围**：长度为36个字符，符合UUID格式。 **默认取值**：不涉及。
	HpsClusterId *string `json:"hps_cluster_id,omitempty"`
}

func (o ServerCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerCreateRequest struct{}"
	}

	return strings.Join([]string{"ServerCreateRequest", string(data)}, " ")
}

type ServerCreateRequestArch struct {
	value string
}

type ServerCreateRequestArchEnum struct {
	ARM ServerCreateRequestArch
	X86 ServerCreateRequestArch
}

func GetServerCreateRequestArchEnum() ServerCreateRequestArchEnum {
	return ServerCreateRequestArchEnum{
		ARM: ServerCreateRequestArch{
			value: "ARM",
		},
		X86: ServerCreateRequestArch{
			value: "X86",
		},
	}
}

func (c ServerCreateRequestArch) Value() string {
	return c.value
}

func (c ServerCreateRequestArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerCreateRequestArch) UnmarshalJSON(b []byte) error {
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

type ServerCreateRequestServerType struct {
	value string
}

type ServerCreateRequestServerTypeEnum struct {
	BMS ServerCreateRequestServerType
	ECS ServerCreateRequestServerType
	HPS ServerCreateRequestServerType
}

func GetServerCreateRequestServerTypeEnum() ServerCreateRequestServerTypeEnum {
	return ServerCreateRequestServerTypeEnum{
		BMS: ServerCreateRequestServerType{
			value: "BMS",
		},
		ECS: ServerCreateRequestServerType{
			value: "ECS",
		},
		HPS: ServerCreateRequestServerType{
			value: "HPS",
		},
	}
}

func (c ServerCreateRequestServerType) Value() string {
	return c.value
}

func (c ServerCreateRequestServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerCreateRequestServerType) UnmarshalJSON(b []byte) error {
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
