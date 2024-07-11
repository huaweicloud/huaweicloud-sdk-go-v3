package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceReq 创建实例请求体。
type CreateInstanceReq struct {

	// 实例名称。  由英文字符开头，只能由英文字母、数字、中划线、下划线组成，长度为4~64的字符。
	Name string `json:"name"`

	// 实例的描述信息。  长度不超过1024的字符串。[且字符串不能包含\">\"与\"<\"，字符串首字符不能为\"=\",\"+\",\"-\",\"@\"的全角和半角字符。](tag:hcs)  > \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\或者\\\"。
	Description *string `json:"description,omitempty"`

	// 消息引擎：rabbitmq。
	Engine CreateInstanceReqEngine `json:"engine"`

	// 消息引擎的版本。   - RabbitMQ版本有：3.8.35[、AMQP-0-9-1](tag:hws,hws_hk)[和3.7.17](tag:tm,hk_tm,hk_sbc,sbc)。
	EngineVersion string `json:"engine_version"`

	// 消息存储空间，单位GB。   [- 单机RabbitMQ实例的存储空间的取值范围100GB~90000GB。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm)    [- 单机RabbitMQ实例的存储空间的取值范围100GB~30000GB。](tag:hcs)    [- 集群RabbitMQ实例的存储空间的取值范围为100GB*节点数~90000GB、200GB*节点数~90000GB、300GB*节点数~90000GB。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm)    [- 集群RabbitMQ实例的存储空间的取值范围为100GB乘以代理数~30000GB乘以代理数。](tag:hcs)
	StorageSpace int32 `json:"storage_space"`

	// 认证用户名，只能由英文字母开头且由英文字母、数字、中划线、下划线组成，长度为4~64的字符。
	AccessUser *string `json:"access_user,omitempty"`

	// 实例的认证密码。  复杂度要求： - 输入长度为8到32位的字符串。 - 必须包含如下四种字符中的两种组合：   - 小写字母   - 大写字母   - 数字   - 特殊字符包括（`~!@#$%^&*()-_=+\\|[{}]:'\",<.>/?）
	Password *string `json:"password,omitempty"`

	// 租户VPC ID。  获取方法如下：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。
	VpcId string `json:"vpc_id"`

	// 租户安全组ID。  获取方法如下：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	// 子网ID。  获取方法如下：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。
	SubnetId string `json:"subnet_id"`

	// 创建节点到指定且有资源的可用区ID。请参考[查询可用区信息](ListAvailableZones.xml)获取可用区ID。  该参数不能为空数组或者数组的值为空。
	AvailableZones []string `json:"available_zones"`

	// 产品标识。  [产品ID可以从[查询产品规格列表](ListEngineProducts.xml)获取。](tag:hws,hws_hk,ctc,cmcc,hws_eu,g42,hk_g42,tm,hk_tm,ocb,hws_ocb) [产品ID可以从[查询产品规格列表](ListProducts.xml)获取。](tag:hk_sbc,sbc) [产品ID可以从“其他接口 > 查询产品规格列表”获取。](tag:hcs)  如果产品ID为集群类型（即对应的type为cluster），broker_num字段为必选。
	ProductId string `json:"product_id"`

	// 代理个数。  当产品为单机类型，代理个数只能为1；当产品为集群类型，可选3、5、7个代理个数。  产品类型为single时:   - 1  产品类型为cluster时:   - 3   - 5   - 7
	BrokerNum *CreateInstanceReqBrokerNum `json:"broker_num,omitempty"`

	// 维护时间窗开始时间，格式为HH:mm。 - 维护时间窗开始和结束时间必须为指定的时间段。 - 开始时间必须为22:00、02:00、06:00、10:00、14:00和18:00。 - 该参数不能单独为空，若该值为空，则结束时间也为空。系统分配一个默认开始时间02:00。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// 维护时间窗结束时间，格式为HH:mm。 - 维护时间窗开始和结束时间必须为指定的时间段。 - 结束时间在开始时间基础上加四个小时，即当开始时间为22:00时，结束时间为02:00。 - 该参数不能单独为空，若该值为空，则开始时间也为空，系统分配一个默认结束时间06:00。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// RabbitMQ实例是否开启公网访问功能。 - true：开启 - false：不开启
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// RabbitMQ实例绑定的弹性IP地址的ID。 如果开启了公网访问功能（即enable_publicip为true），该字段为必选。
	PublicipId *string `json:"publicip_id,omitempty"`

	// 是否打开SSL加密访问。 - true：打开SSL加密访问。 - false：不打开SSL加密访问。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// 存储IO规格。  [如何选择磁盘类型请参考[磁盘类型及性能介绍](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。](tag:hws) [如何选择磁盘类型请参考[磁盘类型及性能介绍](https://support.huaweicloud.com/intl/zh-cn/productdesc-evs/zh-cn_topic_0014580744.html)。](tag:hws_hk) [如何选择磁盘类型请参考[磁盘类型及性能介绍](https://support.huaweicloud.com/eu/productdesc-evs/en-us_topic_0014580744.html)。](tag:hws_eu)  取值范围：   - dms.physical.storage.high.v2   - dms.physical.storage.ultra.v2   [- dms.physical.storage.high.dss.v2(专属云)](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm)   [- dms.physical.storage.ultra.dss.v2(专属云)](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm)
	StorageSpecCode CreateInstanceReqStorageSpecCode `json:"storage_spec_code"`

	// 企业项目ID。若为企业项目账号，该参数必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签列表。
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
	DMS_PHYSICAL_STORAGE_HIGH_V2      CreateInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA_V2     CreateInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_HIGH_DSS_V2  CreateInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA_DSS_V2 CreateInstanceReqStorageSpecCode
}

func GetCreateInstanceReqStorageSpecCodeEnum() CreateInstanceReqStorageSpecCodeEnum {
	return CreateInstanceReqStorageSpecCodeEnum{
		DMS_PHYSICAL_STORAGE_HIGH_V2: CreateInstanceReqStorageSpecCode{
			value: "dms.physical.storage.high.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_V2: CreateInstanceReqStorageSpecCode{
			value: "dms.physical.storage.ultra.v2",
		},
		DMS_PHYSICAL_STORAGE_HIGH_DSS_V2: CreateInstanceReqStorageSpecCode{
			value: "dms.physical.storage.high.dss.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_DSS_V2: CreateInstanceReqStorageSpecCode{
			value: "dms.physical.storage.ultra.dss.v2",
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
