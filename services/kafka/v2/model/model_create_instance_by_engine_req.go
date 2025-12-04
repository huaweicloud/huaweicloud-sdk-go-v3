package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceByEngineReq 创建实例请求体。
type CreateInstanceByEngineReq struct {

	// **参数解释**： 实例名称。 **约束限制**： 由英文字符开头，只能由英文字母、数字、中划线、下划线组成，长度为4~64的字符。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 实例的描述信息。 **约束限制**： 长度不超过1024的字符串。[且字符串不能包含\">\"与\"<\"，字符串首字符不能为\"=\",\"+\",\"-\",\"@\"的全角和半角字符。](tag:hcs,fcs) \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\\\\\或者\\\\\"。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 消息引擎。 **约束限制**： 不涉及 **取值范围**： kafka **默认取值**： 不涉及。
	Engine CreateInstanceByEngineReqEngine `json:"engine"`

	// **参数解释**： 消息引擎的版本。 **约束限制**： 不涉及 **取值范围**： [- 1.1.0](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,sbc,cmcc,ax) [- 2.3.0](tag:g42,tm,hk_g42,ctc,hk_tm,dt,sbc,cmcc,hws_eu) - 2.7 [- 3.x](tag:hws,hws_hk,dt,sbc,hcs,fcs,ctc,tm,hk_tm,hws_eu,ax) **默认取值**： 不涉及。
	EngineVersion string `json:"engine_version"`

	// **参数解释**： 代理个数。 **约束限制**： 不涉及。 **取值范围**： [- Kafka实例规格为kafka.2u4g.cluster.small时，代理数取值范围3-30。](tag:hws,hws_hk,hws_eu,dt,ax) - Kafka实例规格为kafka.2u4g.cluster时，代理数取值范围3-30。 - Kafka实例规格为kafka.4u8g.cluster时，代理数取值范围3-30。 - Kafka实例规格为kafka.8u16g.cluster时，代理数取值范围3-50。 - Kafka实例规格为kafka.12u24g.cluster时，代理数取值范围3-50。 - Kafka实例规格为kafka.16u32g.cluster时，代理数取值范围3-50。 **默认取值**： 不涉及。
	BrokerNum int32 `json:"broker_num"`

	// **参数解释**： 消息存储空间，单位GB。 **约束限制**： 不涉及。 **取值范围**： [- Kafka实例规格为c6.2u4g.cluster时，存储空间取值范围300GB ~ 300000GB。 - Kafka实例规格为c6.4u8g.cluster时，存储空间取值范围300GB ~ 600000GB。 - Kafka实例规格为c6.8u16g.cluster时，存储空间取值范围300GB ~ 1500000GB。 - Kafka实例规格为c6.12u24g.cluster时，存储空间取值范围300GB ~ 1500000GB。 - Kafka实例规格为c6.16u32g.cluster时，存储空间取值范围300GB ~ 1500000GB。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax) [- Kafka实例规格为s6.2u4g.cluster.small时，存储空间取值范围300GB~300000GB。](tag:hws,hws_hk,hws_eu,dt,ax) [- Kafka实例规格为kafka.2u8g.cluster时，存储空间取值范围300GB~300000GB。](tag:fcs) [- Kafka实例规格为kafka.4u16g.cluster时，存储空间取值范围300GB~600000GB。 - Kafka实例规格为kafka.8u32g.cluster时，存储空间取值范围300GB~1500000GB。 - Kafka实例规格为kafka.16u64g.cluster时，存储空间取值范围300GB~1500000GB。 - Kafka实例规格为kafka.32u128g.cluster时，存储空间取值范围300GB~1500000GB。](tag:hcs,fcs) **默认取值**： 不涉及。
	StorageSpace int32 `json:"storage_space"`

	// **参数解释**：  认证用户名。 **约束限制**： 只能由英文字母开头且由英文字母、数字、中划线、下划线组成，长度为4~64的字符。当ssl_enable为true时，该参数必选，ssl_enable为false时，该参数无效。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AccessUser *string `json:"access_user,omitempty"`

	// **参数解释**： 实例的认证密码。 **约束限制**： - 当ssl_enable为true时，该参数必选，ssl_enable为false时，该参数无效。 - 输入长度为8到32位的字符串。 - 必须包含如下四种字符中的三种组合：   - 小写字母。   - 大写字母。   - 数字。   - 特殊字符包括（`~!@#$%^&*()-_=+\\|[{}]:'\",<.>/?）和空格，并且不能以-开头。    **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Password *string `json:"password,omitempty"`

	// **参数解释**：  虚拟私有云ID。获取方法如下：参考《虚拟私有云 API参考》，调用“查询VPC列表”接口，从响应体中获取VPC ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 指定实例所属的安全组。获取方法如下：参考《虚拟私有云 API参考》，调用“查询安全组列表”接口，从响应体中获取安全组ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityGroupId string `json:"security_group_id"`

	// **参数解释**： 子网信息。获取方法如下：参考《虚拟私有云 API参考》，调用“查询子网列表”接口，从响应体中获取子网ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubnetId string `json:"subnet_id"`

	// **参数解释**： 创建节点到指定且有资源的可用区ID。请参考[查询可用区信息](ListAvailableZones.xml)获取可用区ID。 **约束限制**： 该参数不能为空数组或者数组的值为空。  创建Kafka实例，节点需要部署在1个或3个及以上可用区中。如果部署在多个可用区中，以英文逗号隔开多个可用区ID。
	AvailableZones []string `json:"available_zones"`

	// **参数解释**： 产品ID。产品ID可以从[查询产品规格列表](ListEngineProducts.xml)获取。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProductId string `json:"product_id"`

	// **参数解释**： 维护时间窗开始时间。 **约束限制**： 格式为HH:mm。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// **参数解释**： 维护时间窗结束时间。 **约束限制**： 格式为HH:mm。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// **参数解释**： 是否开启公网访问功能。 **约束限制**： 不涉及。 **取值范围**： - true：开启。 - false：不开启。 **默认取值**： false。
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// **参数解释**： 创建实例时可以手动指定实例节点的内网IP地址，仅支持指定IPv4地址。 **约束限制**： 指定内网IP地址数量必须小于等于创建的节点数量。  如果指定的内网IP地址数量小于创建的节点数量时，系统会自动为剩余的节点随机分配内网IP地址。
	TenantIps *[]string `json:"tenant_ips,omitempty"`

	// **参数解释**： 实例绑定的弹性IP地址的ID。 **约束限制**： 以英文逗号隔开多个弹性IP地址的ID。  如果开启了公网访问功能（即enable_publicip为true），该字段为必选。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublicipId *string `json:"publicip_id,omitempty"`

	// **参数解释**： 是否开启SASL加密访问。  [实例创建后将不支持动态开启和关闭。](tag:ocb,hws_ocb,hcs) **约束限制**： 不涉及。 **取值范围**： - true：开启SASL加密访问。 - false：关闭SASL加密访问。 **默认取值**： 不涉及。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// **参数解释**： 开启SASL后使用的安全协议。 **约束限制**： 若该字段值为空，默认开启SASL_SSL认证机制。实例创建后，此参数不支持动态修改。 若创建实例时，使用了port_protocol参数，则Kafka的内网访问安全协议以及公网访问安全协议会使用port_protocol中的值，则此参数无效。 **取值范围**： - SASL_SSL：使用SSL证书加密传输，支持账号密码认证，安全性更高。 - SASL_PLAINTEXT：通过明文传输，支持账号密码认证，性能更好。 **默认取值**： 不涉及。
	KafkaSecurityProtocol *string `json:"kafka_security_protocol,omitempty"`

	// **参数解释**： 开启SASL后使用的认证机制。 **约束限制**： 如果开启了SASL认证功能（即ssl_enable=true），该字段为必选。若该字段值为空，默认开启PLAIN认证机制。
	SaslEnabledMechanisms *[]CreateInstanceByEngineReqSaslEnabledMechanisms `json:"sasl_enabled_mechanisms,omitempty"`

	PortProtocol *PortProtocol `json:"port_protocol,omitempty"`

	// **参数解释**： 磁盘的容量到达容量阈值后，对于消息的处理策略。 **约束限制**： 不涉及。 **取值范围**： - produce_reject：表示拒绝消息写入。 - time_base：表示自动删除最老消息。 **默认取值**： 不涉及。
	RetentionPolicy *CreateInstanceByEngineReqRetentionPolicy `json:"retention_policy,omitempty"`

	// **参数解释**： 是否开启IPv6。 **约束限制**： 仅在虚拟私有云支持IPv6时生效。 **取值范围**： - true：开启。 - false：不开启。 **默认取值**： false。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// **参数解释**： 是否开启磁盘加密。 **约束限制**： 不涉及。 **取值范围**： - true：开启。 - false：不开启。 **默认取值**： false。
	DiskEncryptedEnable *bool `json:"disk_encrypted_enable,omitempty"`

	// **参数解释**： 磁盘加密key，未开启磁盘加密时为空。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DiskEncryptedKey *string `json:"disk_encrypted_key,omitempty"`

	// **参数解释**： 是否开启Smart Connect功能。Smart Connect用于Kafka实例和其他云服务之间的数据同步，或者两个Kafka实例之间的数据同步，实现数据的备份或迁移。 **约束限制**： 不涉及。 **取值范围**： - true：开启。 - false：不开启。 **默认取值**： false。
	ConnectorEnable *bool `json:"connector_enable,omitempty"`

	// **参数解释**： 是否开启kafka自动创建Topic功能。当您选择开启，向一个未创建的Topic生产或消费消息时，系统会自动创建此Topic。 **约束限制**： 不涉及。 **取值范围**： - true：开启。 - false：不开启。 **默认取值**： false。
	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`

	// **参数解释**： 云硬盘类型。[如何选择磁盘类型请参考《云硬盘 [产品介绍](tag:hws,hws_hk,hws_eu,cmcc)[用户指南](tag:dt,g42,hk_g42,ctc,tm,hk_tm,sbc,ocb,hws_ocb,ax)》的“磁盘类型及性能介绍”。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax) **约束限制**： 不涉及。 **取值范围**： - dms.physical.storage.high.v2：高IO云硬盘。 - dms.physical.storage.ultra.v2：超高IO云硬盘。 - [dms.physical.storage.general：通用型SSD云硬盘。](tag:hws,hws_hk,dt,ax) - [dms.physical.storage.extreme：极速型SSD云硬盘。](tag:hws,hws_hk,dt,ax) **默认取值**： 不涉及。
	StorageSpecCode CreateInstanceByEngineReqStorageSpecCode `json:"storage_spec_code"`

	// **参数解释**： 企业项目ID。 **约束限制**： 若为企业项目账号，该参数必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： Kafka实例的标签信息。 **约束限制**： 不涉及。
	Tags *[]TagEntity `json:"tags,omitempty"`

	// **参数解释**： CPU架构。 **约束限制**： 不涉及。 **取值范围**： - X86 [- ARM](tag:hcs,fcs,ctc) **默认取值**： 不涉及。
	ArchType *string `json:"arch_type,omitempty"`

	// **参数解释**： VPC内网明文访问。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VpcClientPlain *bool `json:"vpc_client_plain,omitempty"`

	BssParam *BssParam `json:"bss_param,omitempty"`
}

func (o CreateInstanceByEngineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceByEngineReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceByEngineReq", string(data)}, " ")
}

type CreateInstanceByEngineReqEngine struct {
	value string
}

type CreateInstanceByEngineReqEngineEnum struct {
	KAFKA CreateInstanceByEngineReqEngine
}

func GetCreateInstanceByEngineReqEngineEnum() CreateInstanceByEngineReqEngineEnum {
	return CreateInstanceByEngineReqEngineEnum{
		KAFKA: CreateInstanceByEngineReqEngine{
			value: "kafka",
		},
	}
}

func (c CreateInstanceByEngineReqEngine) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqEngine) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqSaslEnabledMechanisms struct {
	value string
}

type CreateInstanceByEngineReqSaslEnabledMechanismsEnum struct {
	PLAIN         CreateInstanceByEngineReqSaslEnabledMechanisms
	SCRAM_SHA_512 CreateInstanceByEngineReqSaslEnabledMechanisms
}

func GetCreateInstanceByEngineReqSaslEnabledMechanismsEnum() CreateInstanceByEngineReqSaslEnabledMechanismsEnum {
	return CreateInstanceByEngineReqSaslEnabledMechanismsEnum{
		PLAIN: CreateInstanceByEngineReqSaslEnabledMechanisms{
			value: "PLAIN",
		},
		SCRAM_SHA_512: CreateInstanceByEngineReqSaslEnabledMechanisms{
			value: "SCRAM-SHA-512",
		},
	}
}

func (c CreateInstanceByEngineReqSaslEnabledMechanisms) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqSaslEnabledMechanisms) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqSaslEnabledMechanisms) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqRetentionPolicy struct {
	value string
}

type CreateInstanceByEngineReqRetentionPolicyEnum struct {
	TIME_BASE      CreateInstanceByEngineReqRetentionPolicy
	PRODUCE_REJECT CreateInstanceByEngineReqRetentionPolicy
}

func GetCreateInstanceByEngineReqRetentionPolicyEnum() CreateInstanceByEngineReqRetentionPolicyEnum {
	return CreateInstanceByEngineReqRetentionPolicyEnum{
		TIME_BASE: CreateInstanceByEngineReqRetentionPolicy{
			value: "time_base",
		},
		PRODUCE_REJECT: CreateInstanceByEngineReqRetentionPolicy{
			value: "produce_reject",
		},
	}
}

func (c CreateInstanceByEngineReqRetentionPolicy) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqRetentionPolicy) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqStorageSpecCode struct {
	value string
}

type CreateInstanceByEngineReqStorageSpecCodeEnum struct {
	DMS_PHYSICAL_STORAGE_HIGH_V2                  CreateInstanceByEngineReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA_V2                 CreateInstanceByEngineReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_GENERAL_TAGHWSHWS_HKDTAX CreateInstanceByEngineReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_EXTREME_TAGHWSHWS_HKDTAX CreateInstanceByEngineReqStorageSpecCode
}

func GetCreateInstanceByEngineReqStorageSpecCodeEnum() CreateInstanceByEngineReqStorageSpecCodeEnum {
	return CreateInstanceByEngineReqStorageSpecCodeEnum{
		DMS_PHYSICAL_STORAGE_HIGH_V2: CreateInstanceByEngineReqStorageSpecCode{
			value: "dms.physical.storage.high.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_V2: CreateInstanceByEngineReqStorageSpecCode{
			value: "dms.physical.storage.ultra.v2",
		},
		DMS_PHYSICAL_STORAGE_GENERAL_TAGHWSHWS_HKDTAX: CreateInstanceByEngineReqStorageSpecCode{
			value: "[dms.physical.storage.general](tag:hws,hws_hk,dt,ax)",
		},
		DMS_PHYSICAL_STORAGE_EXTREME_TAGHWSHWS_HKDTAX: CreateInstanceByEngineReqStorageSpecCode{
			value: "[dms.physical.storage.extreme](tag:hws,hws_hk,dt,ax)",
		},
	}
}

func (c CreateInstanceByEngineReqStorageSpecCode) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqStorageSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqStorageSpecCode) UnmarshalJSON(b []byte) error {
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
