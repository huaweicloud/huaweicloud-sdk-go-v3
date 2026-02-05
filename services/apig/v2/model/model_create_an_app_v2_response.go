package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateAnAppV2Response Response Object
type CreateAnAppV2Response struct {

	// 编号
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Remark *string `json:"remark,omitempty"`

	// APP的创建者 - USER：用户自行创建 - MARKET：云商店分配  暂不支持MARKET
	Creator *CreateAnAppV2ResponseCreator `json:"creator,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// APP凭据的key。
	AppKey *string `json:"app_key,omitempty"`

	// 凭据关联的账号ID。
	RelatedDomainId *string `json:"related_domain_id,omitempty"`

	// 凭据关联的项目ID。
	RelatedProjectId *string `json:"related_project_id,omitempty"`

	// 密钥
	AppSecret *string `json:"app_secret,omitempty"`

	// 注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// 状态   - 1： 有效
	Status *CreateAnAppV2ResponseStatus `json:"status,omitempty"`

	// APP的类型。 - apig：APIG凭据应用  默认apig，暂不支持其他类型
	AppType *CreateAnAppV2ResponseAppType `json:"app_type,omitempty"`

	// ROMA_APP的类型： - subscription：订阅应用 - integration：集成应用  暂不支持
	RomaAppType    *string `json:"roma_app_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAnAppV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnAppV2Response struct{}"
	}

	return strings.Join([]string{"CreateAnAppV2Response", string(data)}, " ")
}

type CreateAnAppV2ResponseCreator struct {
	value string
}

type CreateAnAppV2ResponseCreatorEnum struct {
	USER   CreateAnAppV2ResponseCreator
	MARKET CreateAnAppV2ResponseCreator
}

func GetCreateAnAppV2ResponseCreatorEnum() CreateAnAppV2ResponseCreatorEnum {
	return CreateAnAppV2ResponseCreatorEnum{
		USER: CreateAnAppV2ResponseCreator{
			value: "USER",
		},
		MARKET: CreateAnAppV2ResponseCreator{
			value: "MARKET",
		},
	}
}

func (c CreateAnAppV2ResponseCreator) Value() string {
	return c.value
}

func (c CreateAnAppV2ResponseCreator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAnAppV2ResponseCreator) UnmarshalJSON(b []byte) error {
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

type CreateAnAppV2ResponseStatus struct {
	value int32
}

type CreateAnAppV2ResponseStatusEnum struct {
	E_1 CreateAnAppV2ResponseStatus
}

func GetCreateAnAppV2ResponseStatusEnum() CreateAnAppV2ResponseStatusEnum {
	return CreateAnAppV2ResponseStatusEnum{
		E_1: CreateAnAppV2ResponseStatus{
			value: 1,
		},
	}
}

func (c CreateAnAppV2ResponseStatus) Value() int32 {
	return c.value
}

func (c CreateAnAppV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAnAppV2ResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateAnAppV2ResponseAppType struct {
	value string
}

type CreateAnAppV2ResponseAppTypeEnum struct {
	APIG CreateAnAppV2ResponseAppType
}

func GetCreateAnAppV2ResponseAppTypeEnum() CreateAnAppV2ResponseAppTypeEnum {
	return CreateAnAppV2ResponseAppTypeEnum{
		APIG: CreateAnAppV2ResponseAppType{
			value: "apig",
		},
	}
}

func (c CreateAnAppV2ResponseAppType) Value() string {
	return c.value
}

func (c CreateAnAppV2ResponseAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAnAppV2ResponseAppType) UnmarshalJSON(b []byte) error {
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
