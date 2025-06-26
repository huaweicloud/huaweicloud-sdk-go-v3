package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateInstanceRequestBody struct {

	// 企业仓库实例名称，小写字母或数字开头，后面跟小写字母、数字、点、下划线或中划线（其中点、下划线、中划线不能直接连续），小写字母或数字结尾，3-64个字符。
	Name string `json:"name"`

	// 企业仓库实例描述
	Description *string `json:"description,omitempty"`

	// 企业仓库实例规格，目前支持基础版(swr.ee.basic)，企业版(swr.ee.professional)
	Spec CreateInstanceRequestBodySpec `json:"spec"`

	// 用户虚拟私有云ID
	VpcId string `json:"vpc_id"`

	// 用户子网的网络ID
	SubnetId string `json:"subnet_id"`

	// vpc和子网所在项目编号
	ProjectId string `json:"project_id"`

	// 实例计费模式，目前只支持按需(postPaid)
	ChargeMode CreateInstanceRequestBodyChargeMode `json:"charge_mode"`

	// 企业项目编号
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 指定资源tag.
	ResourceTags *[]CreateInstanceRequestBodyResourceTags `json:"resource_tags,omitempty"`

	// obs桶是否开启加密, 如果开启了加密，则可以根据encrypt_type指定加密算法
	ObsEncrypt *bool `json:"obs_encrypt,omitempty"`

	// obs桶加密类型，空值使用AES-256加密算法, gm为国密加密算法
	EncryptType *CreateInstanceRequestBodyEncryptType `json:"encrypt_type,omitempty"`

	// 指定obs桶的名称，当指定自定义obs桶之后，则无需对obs_encrypt、encrypt_type进行传值。
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`
}

func (o CreateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBody", string(data)}, " ")
}

type CreateInstanceRequestBodySpec struct {
	value string
}

type CreateInstanceRequestBodySpecEnum struct {
	SWR_EE_BASIC        CreateInstanceRequestBodySpec
	SWR_EE_PROFESSIONAL CreateInstanceRequestBodySpec
}

func GetCreateInstanceRequestBodySpecEnum() CreateInstanceRequestBodySpecEnum {
	return CreateInstanceRequestBodySpecEnum{
		SWR_EE_BASIC: CreateInstanceRequestBodySpec{
			value: "swr.ee.basic",
		},
		SWR_EE_PROFESSIONAL: CreateInstanceRequestBodySpec{
			value: "swr.ee.professional",
		},
	}
}

func (c CreateInstanceRequestBodySpec) Value() string {
	return c.value
}

func (c CreateInstanceRequestBodySpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceRequestBodySpec) UnmarshalJSON(b []byte) error {
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

type CreateInstanceRequestBodyChargeMode struct {
	value string
}

type CreateInstanceRequestBodyChargeModeEnum struct {
	POST_PAID CreateInstanceRequestBodyChargeMode
}

func GetCreateInstanceRequestBodyChargeModeEnum() CreateInstanceRequestBodyChargeModeEnum {
	return CreateInstanceRequestBodyChargeModeEnum{
		POST_PAID: CreateInstanceRequestBodyChargeMode{
			value: "postPaid",
		},
	}
}

func (c CreateInstanceRequestBodyChargeMode) Value() string {
	return c.value
}

func (c CreateInstanceRequestBodyChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceRequestBodyChargeMode) UnmarshalJSON(b []byte) error {
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

type CreateInstanceRequestBodyEncryptType struct {
	value string
}

type CreateInstanceRequestBodyEncryptTypeEnum struct {
	GM CreateInstanceRequestBodyEncryptType
}

func GetCreateInstanceRequestBodyEncryptTypeEnum() CreateInstanceRequestBodyEncryptTypeEnum {
	return CreateInstanceRequestBodyEncryptTypeEnum{
		GM: CreateInstanceRequestBodyEncryptType{
			value: "gm",
		},
	}
}

func (c CreateInstanceRequestBodyEncryptType) Value() string {
	return c.value
}

func (c CreateInstanceRequestBodyEncryptType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceRequestBodyEncryptType) UnmarshalJSON(b []byte) error {
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
