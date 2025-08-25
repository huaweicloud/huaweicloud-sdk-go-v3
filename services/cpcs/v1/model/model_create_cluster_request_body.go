package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateClusterRequestBody struct {

	// 计费模式  - 0：包年/包月。
	ChargingMode CreateClusterRequestBodyChargingMode `json:"charging_mode"`

	// 局点ID
	RegionId string `json:"region_id"`

	// 周期类型  - 2：月。  - 3：年。
	PeriodType CreateClusterRequestBodyPeriodType `json:"period_type"`

	// 周期数量
	PeriodNum int32 `json:"period_num"`

	// 实例数量
	SubscriptionNum int32 `json:"subscription_num"`

	// 产品id
	PeriodProductId string `json:"period_product_id"`

	// 产品规格code
	PeriodResourceSpecCode string `json:"period_resource_spec_code"`

	// 集群名称
	ClusterName string `json:"cluster_name"`

	// 服务类型  - ENCRYPT_DECRYPT：加解密服务。  - SIGN_VERIFY：签名验签服务。  - KMS：密钥管理服务。  - TIMESTAMP：时间戳服务。  - COLLA_SIGN：协同签名服务。  - OTP：动态令牌服务。  - DB_ENCRYPT：数据库加密服务。  - FILE_ENCRYPT：文件加密服务。  - DIGIT_SEAL：电子签章服务。  - SSL_VPN：SSL_VPN服务。  - IAS：身份认证服务。
	ServiceType string `json:"service_type"`

	// 共享类型  - EXCLUSIVE：应用独享。  - SHARED：应用共享。
	ShareType CreateClusterRequestBodyShareType `json:"share_type"`

	// 镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 是否自动续费
	IsAutoRenew int32 `json:"is_auto_renew"`

	// 折扣
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例类型  - SINGLE：单机。  - STANDBY：主备。
	Type *CreateClusterRequestBodyType `json:"type,omitempty"`

	// az
	Az string `json:"az"`
}

func (o CreateClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterRequestBody struct{}"
	}

	return strings.Join([]string{"CreateClusterRequestBody", string(data)}, " ")
}

type CreateClusterRequestBodyChargingMode struct {
	value int32
}

type CreateClusterRequestBodyChargingModeEnum struct {
	E_0 CreateClusterRequestBodyChargingMode
}

func GetCreateClusterRequestBodyChargingModeEnum() CreateClusterRequestBodyChargingModeEnum {
	return CreateClusterRequestBodyChargingModeEnum{
		E_0: CreateClusterRequestBodyChargingMode{
			value: 0,
		},
	}
}

func (c CreateClusterRequestBodyChargingMode) Value() int32 {
	return c.value
}

func (c CreateClusterRequestBodyChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterRequestBodyChargingMode) UnmarshalJSON(b []byte) error {
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

type CreateClusterRequestBodyPeriodType struct {
	value int32
}

type CreateClusterRequestBodyPeriodTypeEnum struct {
	E_2 CreateClusterRequestBodyPeriodType
	E_3 CreateClusterRequestBodyPeriodType
}

func GetCreateClusterRequestBodyPeriodTypeEnum() CreateClusterRequestBodyPeriodTypeEnum {
	return CreateClusterRequestBodyPeriodTypeEnum{
		E_2: CreateClusterRequestBodyPeriodType{
			value: 2,
		}, E_3: CreateClusterRequestBodyPeriodType{
			value: 3,
		},
	}
}

func (c CreateClusterRequestBodyPeriodType) Value() int32 {
	return c.value
}

func (c CreateClusterRequestBodyPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterRequestBodyPeriodType) UnmarshalJSON(b []byte) error {
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

type CreateClusterRequestBodyShareType struct {
	value string
}

type CreateClusterRequestBodyShareTypeEnum struct {
	EXCLUSIVE CreateClusterRequestBodyShareType
	SHARED    CreateClusterRequestBodyShareType
}

func GetCreateClusterRequestBodyShareTypeEnum() CreateClusterRequestBodyShareTypeEnum {
	return CreateClusterRequestBodyShareTypeEnum{
		EXCLUSIVE: CreateClusterRequestBodyShareType{
			value: "EXCLUSIVE",
		},
		SHARED: CreateClusterRequestBodyShareType{
			value: "SHARED",
		},
	}
}

func (c CreateClusterRequestBodyShareType) Value() string {
	return c.value
}

func (c CreateClusterRequestBodyShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterRequestBodyShareType) UnmarshalJSON(b []byte) error {
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

type CreateClusterRequestBodyType struct {
	value string
}

type CreateClusterRequestBodyTypeEnum struct {
	SINGLE  CreateClusterRequestBodyType
	STANDBY CreateClusterRequestBodyType
}

func GetCreateClusterRequestBodyTypeEnum() CreateClusterRequestBodyTypeEnum {
	return CreateClusterRequestBodyTypeEnum{
		SINGLE: CreateClusterRequestBodyType{
			value: "SINGLE",
		},
		STANDBY: CreateClusterRequestBodyType{
			value: "STANDBY",
		},
	}
}

func (c CreateClusterRequestBodyType) Value() string {
	return c.value
}

func (c CreateClusterRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterRequestBodyType) UnmarshalJSON(b []byte) error {
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
