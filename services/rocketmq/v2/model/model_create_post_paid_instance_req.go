package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePostPaidInstanceReq 创建实例请求体。
type CreatePostPaidInstanceReq struct {

	// 实例名称。  由英文字符开头，只能由英文字母、数字、中划线、下划线组成，长度为4~64的字符。
	Name string `json:"name"`

	// 实例的描述信息。  长度不超过1024的字符串。[且字符串不能包含\">\"与\"<\"，字符串首字符不能为\"=\",\"+\",\"-\",\"@\"的全角和半角字符。](tag:hcs)  > \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\或者\\\"。
	Description *string `json:"description,omitempty"`

	// 消息引擎。取值填写为：reliability。
	Engine CreatePostPaidInstanceReqEngine `json:"engine"`

	// 消息引擎的版本。取值填写为：[4.8.0](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[5.x](tag:hcs,fcs)。
	EngineVersion CreatePostPaidInstanceReqEngineVersion `json:"engine_version"`

	// 存储空间。
	StorageSpace int32 `json:"storage_space"`

	// 虚拟私有云ID。  获取方法如下：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。
	VpcId string `json:"vpc_id"`

	// 子网信息。  获取方法如下：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。
	SubnetId string `json:"subnet_id"`

	// 指定实例所属的安全组。  获取方法如下：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	// 创建节点到指定且有资源的可用区ID。请参考[查询可用区信息](ListAvailableZones.xml)获取可用区ID。 该参数不能为空数组或者数组的值为空， 请注意查看该可用区是否有资源。  创建RocketMQ实例，支持节点部署在1个或[3个及3个以上的可用区。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[2个可用区。](tag:fcs)在为节点指定可用区时，用逗号分隔开。
	AvailableZones []string `json:"available_zones"`

	// RocketMQ实例规格。[x86环境后缀为.x86，arm环境为.arm。](tag:hcs,fcs)   - [c6.4u8g.cluster：单个代理最大Topic数4000，单个代理最大消费组数4000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[c6.4u16g.cluster.x86或c6.4u16g.cluster.arm：单个代理最大分区数100，单个代理最大消费组数200](tag:hcs)   - [c6.8u16g.cluster：单个代理最大Topic数8000，单个代理最大消费组数8000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[c6.8u32g.cluster.x86或c6.8u32g.cluster.arm：单个代理最大Topic数200，单个代理最大消费组数400](tag:hcs,fcs)   - [c6.12u24g.cluster：单个代理最大Topic数12000，单个代理最大消费组数12000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[c6.16u64g.cluster.x86或c6.16u64g.cluster.arm：单个代理最大Topic数300，单个代理最大消费组数600](tag:hcs,fcs)   - [c6.16u32g.cluster：单个代理最大Topic数16000，单个代理最大消费组数16000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[c6.32u128g.cluster.x86或c6.32u128g.cluster.arm：单个代理最大Topic数400，单个代理最大消费组数800](tag:hcs,fcs)
	ProductId CreatePostPaidInstanceReqProductId `json:"product_id"`

	// 是否打开SSL加密访问。 - true：打开SSL加密访问。 - false：不打开SSL加密访问。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// 存储IO规格。   - dms.physical.storage.high.v2: 高IO类型磁盘   - dms.physical.storage.ultra.v2: 超高IO类型磁盘
	StorageSpecCode CreatePostPaidInstanceReqStorageSpecCode `json:"storage_spec_code"`

	// 企业项目ID。若为企业项目账号，该参数必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否开启访问控制列表。
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// 是否支持IPv6。   - true：支持   - false：不支持
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 是否开启公网访问功能。默认不开启公网。 - true：开启 - false：不开启
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// 实例绑定的弹性IP地址的ID。  以英文逗号隔开多个弹性IP地址的ID。  如果开启了公网访问功能（即enable_publicip为true），该字段为必选。
	PublicipId *string `json:"publicip_id,omitempty"`

	// 代理个数。
	BrokerNum int32 `json:"broker_num"`
}

func (o CreatePostPaidInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceReq struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceReq", string(data)}, " ")
}

type CreatePostPaidInstanceReqEngine struct {
	value string
}

type CreatePostPaidInstanceReqEngineEnum struct {
	RELIABILITY CreatePostPaidInstanceReqEngine
}

func GetCreatePostPaidInstanceReqEngineEnum() CreatePostPaidInstanceReqEngineEnum {
	return CreatePostPaidInstanceReqEngineEnum{
		RELIABILITY: CreatePostPaidInstanceReqEngine{
			value: "reliability",
		},
	}
}

func (c CreatePostPaidInstanceReqEngine) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqEngine) UnmarshalJSON(b []byte) error {
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

type CreatePostPaidInstanceReqEngineVersion struct {
	value string
}

type CreatePostPaidInstanceReqEngineVersionEnum struct {
	E_4_8_0_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_5_X_TAGHCS CreatePostPaidInstanceReqEngineVersion
}

func GetCreatePostPaidInstanceReqEngineVersionEnum() CreatePostPaidInstanceReqEngineVersionEnum {
	return CreatePostPaidInstanceReqEngineVersionEnum{
		E_4_8_0_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_5_X_TAGHCS: CreatePostPaidInstanceReqEngineVersion{
			value: "[4.8.0](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[5.x](tag:hcs)",
		},
	}
}

func (c CreatePostPaidInstanceReqEngineVersion) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqEngineVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqEngineVersion) UnmarshalJSON(b []byte) error {
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

type CreatePostPaidInstanceReqProductId struct {
	value string
}

type CreatePostPaidInstanceReqProductIdEnum struct {
	C6_4U8G_CLUSTER_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_C6_4U16G_CLUSTER_X86_TAGHCSFCS     CreatePostPaidInstanceReqProductId
	C6_8U16G_CLUSTER_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_C6_8U32G_CLUSTER_X86_TAGHCSFCS    CreatePostPaidInstanceReqProductId
	C6_12U24G_CLUSTER_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_C6_16U64G_CLUSTER_X86_TAGHCSFCS  CreatePostPaidInstanceReqProductId
	C6_16U32G_CLUSTER_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_C6_32U128G_CLUSTER_X86_TAGHCSFCS CreatePostPaidInstanceReqProductId
	C6_4U16G_CLUSTER_ARM_TAGHCSFCS                                                                              CreatePostPaidInstanceReqProductId
	C6_8U32G_CLUSTER_ARM_TAGHCSFCS                                                                              CreatePostPaidInstanceReqProductId
	C6_16U64G_CLUSTER_ARM_TAGHCSFCS                                                                             CreatePostPaidInstanceReqProductId
	C6_32U128G_CLUSTER_ARM_TAGHCSFCS                                                                            CreatePostPaidInstanceReqProductId
}

func GetCreatePostPaidInstanceReqProductIdEnum() CreatePostPaidInstanceReqProductIdEnum {
	return CreatePostPaidInstanceReqProductIdEnum{
		C6_4U8G_CLUSTER_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_C6_4U16G_CLUSTER_X86_TAGHCSFCS: CreatePostPaidInstanceReqProductId{
			value: "[c6.4u8g.cluster](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[c6.4u16g.cluster.x86](tag:hcs,fcs)",
		},
		C6_8U16G_CLUSTER_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_C6_8U32G_CLUSTER_X86_TAGHCSFCS: CreatePostPaidInstanceReqProductId{
			value: "[c6.8u16g.cluster](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[c6.8u32g.cluster.x86](tag:hcs,fcs)",
		},
		C6_12U24G_CLUSTER_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_C6_16U64G_CLUSTER_X86_TAGHCSFCS: CreatePostPaidInstanceReqProductId{
			value: "[c6.12u24g.cluster](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[c6.16u64g.cluster.x86](tag:hcs,fcs)",
		},
		C6_16U32G_CLUSTER_TAGHWSHWS_EUHWS_HKOCBHWS_OCBCTCG42HK_G42TMSBCHK_SBCHK_TM_C6_32U128G_CLUSTER_X86_TAGHCSFCS: CreatePostPaidInstanceReqProductId{
			value: "[c6.16u32g.cluster](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[c6.32u128g.cluster.x86](tag:hcs,fcs)",
		},
		C6_4U16G_CLUSTER_ARM_TAGHCSFCS: CreatePostPaidInstanceReqProductId{
			value: "[c6.4u16g.cluster.arm](tag:hcs,fcs)",
		},
		C6_8U32G_CLUSTER_ARM_TAGHCSFCS: CreatePostPaidInstanceReqProductId{
			value: "[c6.8u32g.cluster.arm](tag:hcs,fcs)",
		},
		C6_16U64G_CLUSTER_ARM_TAGHCSFCS: CreatePostPaidInstanceReqProductId{
			value: "[c6.16u64g.cluster.arm](tag:hcs,fcs)",
		},
		C6_32U128G_CLUSTER_ARM_TAGHCSFCS: CreatePostPaidInstanceReqProductId{
			value: "[c6.32u128g.cluster.arm](tag:hcs,fcs)",
		},
	}
}

func (c CreatePostPaidInstanceReqProductId) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqProductId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqProductId) UnmarshalJSON(b []byte) error {
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

type CreatePostPaidInstanceReqStorageSpecCode struct {
	value string
}

type CreatePostPaidInstanceReqStorageSpecCodeEnum struct {
	DMS_PHYSICAL_STORAGE_HIGH_V2  CreatePostPaidInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA_V2 CreatePostPaidInstanceReqStorageSpecCode
}

func GetCreatePostPaidInstanceReqStorageSpecCodeEnum() CreatePostPaidInstanceReqStorageSpecCodeEnum {
	return CreatePostPaidInstanceReqStorageSpecCodeEnum{
		DMS_PHYSICAL_STORAGE_HIGH_V2: CreatePostPaidInstanceReqStorageSpecCode{
			value: "dms.physical.storage.high.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_V2: CreatePostPaidInstanceReqStorageSpecCode{
			value: "dms.physical.storage.ultra.v2",
		},
	}
}

func (c CreatePostPaidInstanceReqStorageSpecCode) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqStorageSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqStorageSpecCode) UnmarshalJSON(b []byte) error {
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
