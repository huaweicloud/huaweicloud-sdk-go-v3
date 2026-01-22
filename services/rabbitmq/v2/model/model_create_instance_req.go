package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceReq 创建实例请求体。
type CreateInstanceReq struct {

	// **参数解释**： 实例名称。 **约束限制**： 由英文字符开头，只能由英文字母、数字、中划线、下划线组成，长度为4~64的字符。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 实例的描述信息。 **约束限制**： 长度不超过1024的字符串。[且字符串不能包含\">\"与\"<\"，字符串首字符不能为\"=\",\"+\",\"-\",\"@\"的全角和半角字符。](tag:hcs,fcs)  \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\\\\\或者\\\\\"。  **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 消息引擎。 **约束限制**： 不涉及 **取值范围**： rabbitmq：RabbitMQ引擎。 **默认取值**： 不涉及。
	Engine CreateInstanceReqEngine `json:"engine"`

	// **参数解释**： 消息引擎的版本。 **约束限制**： 不涉及 **取值范围**： - 3.8.35 [- 3.12.13](tag:srg) [- AMQP-0-9-1](tag:hws,hws_hk,hws_eu) **默认取值**： 不涉及。
	EngineVersion string `json:"engine_version"`

	// **参数解释**： ACL访问控制 **约束限制**： 仅AMQP版本支持此参数。 **取值范围**： - true：开启ACL访问控制。 - false：不开启ACL访问控制。 **默认取值**： 不涉及。
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// **参数解释**： 消息存储空间，单位GB。 **约束限制**： 磁盘容量仅支持设置为100的整数倍。 **取值范围**： - 单机实例：100GB~30000GB。  - 集群实例：100GB * 节点数 ~ 30000GB * 节点数。 **默认取值**： 不涉及。
	StorageSpace int32 `json:"storage_space"`

	// **参数解释**：  认证用户名。 **约束限制**： 只能由英文字母开头且由英文字母、数字、中划线、下划线组成，长度为4~64的字符。当ssl_enable为true时，该参数必选，ssl_enable为false时，该参数无效。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AccessUser *string `json:"access_user,omitempty"`

	// **参数解释**： 实例的认证密码。 **约束限制**： - 输入长度为8到32位的字符串。 - 必须包含如下四种字符中的三种组合：   - 小写字母   - 大写字母   - 数字   - 特殊字符包括（`~!@#$%^&*()-_=+\\|[{}]:'\",<.>/?）和空格，并且不能以-开头 - 当ssl_enable为true时，该参数必选，ssl_enable为false时，该参数无效。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Password *string `json:"password,omitempty"`

	// **参数解释**：  虚拟私有云ID。获取方法如下：参考[[《虚拟私有云 API参考》](https://support.huaweicloud.com/api-vpc/vpc_apiv3_0003.html)](tag:hws)[[《虚拟私有云 API参考》](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_apiv3_0003.html)](tag:hws_hk)[[《虚拟私有云 API参考》](https://support.huaweicloud.com/eu/api-vpc/vpc_apiv3_0003.html)](tag:hws_eu)[《虚拟私有云 API参考》](tag:ax,cmcc,ctc,sbc,hk_sbc,g42,hk_g42,tm,hk_tm,srg)，调用“查询VPC列表”接口，从响应体中获取VPC ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 指定实例所属的安全组。获取方法如下：参考[[《虚拟私有云 API参考》](https://support.huaweicloud.com/api-vpc/vpc_apiv3_0011.html)](tag:hws)[[《虚拟私有云 API参考》](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_apiv3_0011.html)](tag:hws_hk)[[《虚拟私有云 API参考》](https://support.huaweicloud.com/eu/api-vpc/vpc_apiv3_0011.html)](tag:hws_eu)[《虚拟私有云 API参考》](tag:ax,cmcc,ctc,sbc,hk_sbc,g42,hk_g42,tm,hk_tm,srg)，调用“查询安全组列表”接口，从响应体中获取安全组ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityGroupId string `json:"security_group_id"`

	// **参数解释**： 子网信息。获取方法如下：参考[[《虚拟私有云 API参考》](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)](tag:hws)[[《虚拟私有云 API参考》](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_subnet01_0003.html)](tag:hws_hk)[[《虚拟私有云 API参考》](https://support.huaweicloud.com/eu/api-vpc/vpc_subnet01_0003.html)](tag:hws_eu)[《虚拟私有云 API参考》](tag:ax,cmcc,ctc,sbc,hk_sbc,g42,hk_g42,tm,hk_tm,srg)，调用“查询子网列表”接口，从响应体中获取子网ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubnetId string `json:"subnet_id"`

	// **参数解释**： 创建节点到指定且有资源的可用区ID。请参考[查询可用区信息](ListAvailableZones.xml)获取可用区ID。 **约束限制**： 该参数不能为空数组或者数组的值为空。  创建RabbitMQ实例，节点需要部署在1个或3个及以上可用区中。如果部署在多个可用区中，以英文逗号隔开多个可用区ID。
	AvailableZones []string `json:"available_zones"`

	// **参数解释**： 产品ID。产品ID可以从[查询产品规格列表](ListEngineProducts.xml)获取。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProductId string `json:"product_id"`

	// **参数解释**： 代理个数。 **约束限制**： 当产品为单机类型，代理个数只能为1；当产品为集群类型，可选3、5、7个代理个数。 **取值范围**： - 1 - 3 - 5 - 7 **默认取值**： 不涉及。
	BrokerNum *CreateInstanceReqBrokerNum `json:"broker_num,omitempty"`

	// **参数解释**： 维护时间窗开始时间。 **约束限制**： 格式为HH:mm。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// **参数解释**： 维护时间窗结束时间。 **约束限制**： 格式为HH:mm。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// **参数解释**： 是否开启公网访问功能。 **约束限制**： 不涉及。 **取值范围**： - true：开启 - false：不开启 **默认取值**： false。
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// **参数解释**： 实例绑定的弹性IP地址的ID。获取方法：参考[[《弹性公网IP API参考》](https://support.huaweicloud.com/api-eip/ListPublicipsV3.html)](tag:hws)[[《弹性公网IP API参考》](https://support.huaweicloud.com/intl/zh-cn/api-eip/ListPublicipsV3.html)](tag:hws_hk)[[《弹性公网IP API参考》](https://support.huaweicloud.com/eu/api-eip/ListPublicipsV3.html)](tag:hws_eu)[《弹性公网IP API参考》](tag:ax,cmcc,ctc,sbc,hk_sbc,g42,hk_g42,tm,hk_tm,srg)，调用“查询弹性公网IP列表”接口，从响应体中获取弹性公网IP的ID。 **约束限制**： 以英文逗号隔开多个弹性IP地址的ID。  如果开启了公网访问功能（即enable_publicip为true），该字段为必选。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublicipId *string `json:"publicip_id,omitempty"`

	// **参数解释**： 是否开启SSL加密访问。 **约束限制**： 不涉及。 **取值范围**： - true：开启SSL加密访问。 - false：关闭SSL加密访问。 **默认取值**： 不涉及。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// **参数解释**： 存储IO规格。  [如何选择磁盘类型请参考[磁盘类型及性能介绍](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0014580744.html)。](tag:hws) [如何选择磁盘类型请参考[磁盘类型及性能介绍](https://support.huaweicloud.com/intl/zh-cn/productdesc-evs/zh-cn_topic_0014580744.html)。](tag:hws_hk) [如何选择磁盘类型请参考[磁盘类型及性能介绍](https://support.huaweicloud.com/eu/productdesc-evs/en-us_topic_0014580744.html)。](tag:hws_eu) **约束限制**： 不涉及。 **取值范围**： - dms.physical.storage.high.v2：高IO云硬盘。 - dms.physical.storage.ultra.v2：超高IO云硬盘。 [- dms.physical.storage.general：通用型SSD云硬盘。](tag:hws,hws_hk,dt,ax) [- dms.physical.storage.extreme：极速型SSD云硬盘。](tag:hws,hws_hk,dt,ax) **默认取值**： 不涉及。
	StorageSpecCode CreateInstanceReqStorageSpecCode `json:"storage_spec_code"`

	// **参数解释**： 企业项目ID。 **约束限制**： 若为企业项目账号，该参数必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 标签列表。 **约束限制**： 一个RabbitMQ实例最多添加20个标签。
	Tags *[]TagEntity `json:"tags,omitempty"`

	BssParam *BssParam `json:"bss_param,omitempty"`
}

func (o CreateInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceReq", string(data)}, " ")
}

type CreateInstanceReqEngine struct {
	value string
}

type CreateInstanceReqEngineEnum struct {
	RABBITMQ CreateInstanceReqEngine
}

func GetCreateInstanceReqEngineEnum() CreateInstanceReqEngineEnum {
	return CreateInstanceReqEngineEnum{
		RABBITMQ: CreateInstanceReqEngine{
			value: "rabbitmq",
		},
	}
}

func (c CreateInstanceReqEngine) Value() string {
	return c.value
}

func (c CreateInstanceReqEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceReqEngine) UnmarshalJSON(b []byte) error {
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

type CreateInstanceReqBrokerNum struct {
	value int32
}

type CreateInstanceReqBrokerNumEnum struct {
	E_1 CreateInstanceReqBrokerNum
	E_3 CreateInstanceReqBrokerNum
	E_5 CreateInstanceReqBrokerNum
	E_7 CreateInstanceReqBrokerNum
}

func GetCreateInstanceReqBrokerNumEnum() CreateInstanceReqBrokerNumEnum {
	return CreateInstanceReqBrokerNumEnum{
		E_1: CreateInstanceReqBrokerNum{
			value: 1,
		}, E_3: CreateInstanceReqBrokerNum{
			value: 3,
		}, E_5: CreateInstanceReqBrokerNum{
			value: 5,
		}, E_7: CreateInstanceReqBrokerNum{
			value: 7,
		},
	}
}

func (c CreateInstanceReqBrokerNum) Value() int32 {
	return c.value
}

func (c CreateInstanceReqBrokerNum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceReqBrokerNum) UnmarshalJSON(b []byte) error {
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

type CreateInstanceReqStorageSpecCode struct {
	value string
}

type CreateInstanceReqStorageSpecCodeEnum struct {
	DMS_PHYSICAL_STORAGE_HIGH_V2                  CreateInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA_V2                 CreateInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_GENERAL_TAGHWSHWS_HKDTAX CreateInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_EXTREME_TAGHWSHWS_HKDTAX CreateInstanceReqStorageSpecCode
}

func GetCreateInstanceReqStorageSpecCodeEnum() CreateInstanceReqStorageSpecCodeEnum {
	return CreateInstanceReqStorageSpecCodeEnum{
		DMS_PHYSICAL_STORAGE_HIGH_V2: CreateInstanceReqStorageSpecCode{
			value: "dms.physical.storage.high.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_V2: CreateInstanceReqStorageSpecCode{
			value: "dms.physical.storage.ultra.v2",
		},
		DMS_PHYSICAL_STORAGE_GENERAL_TAGHWSHWS_HKDTAX: CreateInstanceReqStorageSpecCode{
			value: "[dms.physical.storage.general](tag:hws,hws_hk,dt,ax)",
		},
		DMS_PHYSICAL_STORAGE_EXTREME_TAGHWSHWS_HKDTAX: CreateInstanceReqStorageSpecCode{
			value: "[dms.physical.storage.extreme](tag:hws,hws_hk,dt,ax)",
		},
	}
}

func (c CreateInstanceReqStorageSpecCode) Value() string {
	return c.value
}

func (c CreateInstanceReqStorageSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceReqStorageSpecCode) UnmarshalJSON(b []byte) error {
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
