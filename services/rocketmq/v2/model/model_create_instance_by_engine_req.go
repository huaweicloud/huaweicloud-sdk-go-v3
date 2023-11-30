package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceByEngineReq 创建实例请求体。
type CreateInstanceByEngineReq struct {

	// 实例名称。  由英文字符开头，只能由英文字母、数字、中划线、下划线组成，长度为4~64的字符。
	Name string `json:"name"`

	// 实例的描述信息。  长度不超过1024的字符串。  > \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\或者\\\"。
	Description *string `json:"description,omitempty"`

	// 消息引擎。取值填写为：reliability。
	Engine CreateInstanceByEngineReqEngine `json:"engine"`

	// 消息引擎的版本。取值填写为：4.8.0。
	EngineVersion CreateInstanceByEngineReqEngineVersion `json:"engine_version"`

	// 存储空间。
	StorageSpace int32 `json:"storage_space"`

	// 虚拟私有云ID。  获取方法如下：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。
	VpcId string `json:"vpc_id"`

	// 子网信息。  获取方法如下：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。
	SubnetId string `json:"subnet_id"`

	// 指定实例所属的安全组。  获取方法如下：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	// 创建节点到指定且有资源的可用区ID。请参考[查询可用区信息](ListAvailableZones.xml)获取可用区ID。  该参数不能为空数组或者数组的值为空， 请注意查看该可用区是否有资源。  创建RocketMQ实例，支持节点部署在1个或3个及3个以上的可用区。在为节点指定可用区时，用逗号分隔开。
	AvailableZones []string `json:"available_zones"`

	// RocketMQ实例规格。   - c6.4u8g.cluster.small：单个代理最大Topic数2000，单个代理最大消费组数2000   - c6.4u8g.cluster：单个代理最大Topic数4000，单个代理最大消费组数4000   - c6.8u16g.cluster：单个代理最大Topic数8000，单个代理最大消费组数8000   - c6.12u24g.cluster：单个代理最大Topic数12000，单个代理最大消费组数12000   - c6.16u32g.cluster：单个代理最大Topic数16000，单个代理最大消费组数16000
	ProductId CreateInstanceByEngineReqProductId `json:"product_id"`

	// 是否打开SSL加密访问。 - true：打开SSL加密访问。 - false：不打开SSL加密访问。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// 存储IO规格。   - dms.physical.storage.high.v2: 高IO类型磁盘   - dms.physical.storage.ultra.v2: 超高IO类型磁盘
	StorageSpecCode CreateInstanceByEngineReqStorageSpecCode `json:"storage_spec_code"`

	// 企业项目ID。若为企业项目账号，该参数必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否开启访问控制列表。
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// 是否支持IPv6。   - true: 支持   - false：不支持
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 是否开启公网访问功能。默认不开启公网。 - true：开启 - false：不开启
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// 实例绑定的弹性IP地址的ID。  以英文逗号隔开多个弹性IP地址的ID。  如果开启了公网访问功能（即enable_publicip为true），该字段为必选。
	PublicipId *string `json:"publicip_id,omitempty"`

	// 代理个数
	BrokerNum int32 `json:"broker_num"`

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
	RELIABILITY CreateInstanceByEngineReqEngine
}

func GetCreateInstanceByEngineReqEngineEnum() CreateInstanceByEngineReqEngineEnum {
	return CreateInstanceByEngineReqEngineEnum{
		RELIABILITY: CreateInstanceByEngineReqEngine{
			value: "reliability",
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

type CreateInstanceByEngineReqEngineVersion struct {
	value string
}

type CreateInstanceByEngineReqEngineVersionEnum struct {
	E_4_8_0 CreateInstanceByEngineReqEngineVersion
}

func GetCreateInstanceByEngineReqEngineVersionEnum() CreateInstanceByEngineReqEngineVersionEnum {
	return CreateInstanceByEngineReqEngineVersionEnum{
		E_4_8_0: CreateInstanceByEngineReqEngineVersion{
			value: "4.8.0",
		},
	}
}

func (c CreateInstanceByEngineReqEngineVersion) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqEngineVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqEngineVersion) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqProductId struct {
	value string
}

type CreateInstanceByEngineReqProductIdEnum struct {
	C6_4U8G_CLUSTER_SMALL CreateInstanceByEngineReqProductId
	C6_4U8G_CLUSTER       CreateInstanceByEngineReqProductId
	C6_8U16G_CLUSTER      CreateInstanceByEngineReqProductId
	C6_12U24G_CLUSTER     CreateInstanceByEngineReqProductId
	C6_16U32G_CLUSTER     CreateInstanceByEngineReqProductId
}

func GetCreateInstanceByEngineReqProductIdEnum() CreateInstanceByEngineReqProductIdEnum {
	return CreateInstanceByEngineReqProductIdEnum{
		C6_4U8G_CLUSTER_SMALL: CreateInstanceByEngineReqProductId{
			value: "c6.4u8g.cluster.small",
		},
		C6_4U8G_CLUSTER: CreateInstanceByEngineReqProductId{
			value: "c6.4u8g.cluster",
		},
		C6_8U16G_CLUSTER: CreateInstanceByEngineReqProductId{
			value: "c6.8u16g.cluster",
		},
		C6_12U24G_CLUSTER: CreateInstanceByEngineReqProductId{
			value: "c6.12u24g.cluster",
		},
		C6_16U32G_CLUSTER: CreateInstanceByEngineReqProductId{
			value: "c6.16u32g.cluster",
		},
	}
}

func (c CreateInstanceByEngineReqProductId) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqProductId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqProductId) UnmarshalJSON(b []byte) error {
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
	DMS_PHYSICAL_STORAGE_HIGH_V2  CreateInstanceByEngineReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA_V2 CreateInstanceByEngineReqStorageSpecCode
}

func GetCreateInstanceByEngineReqStorageSpecCodeEnum() CreateInstanceByEngineReqStorageSpecCodeEnum {
	return CreateInstanceByEngineReqStorageSpecCodeEnum{
		DMS_PHYSICAL_STORAGE_HIGH_V2: CreateInstanceByEngineReqStorageSpecCode{
			value: "dms.physical.storage.high.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_V2: CreateInstanceByEngineReqStorageSpecCode{
			value: "dms.physical.storage.ultra.v2",
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
