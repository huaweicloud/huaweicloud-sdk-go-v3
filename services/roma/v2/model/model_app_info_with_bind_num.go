package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type AppInfoWithBindNum struct {

	// 编号
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Remark *string `json:"remark,omitempty"`

	// APP的创建者 - USER：用户自行创建 - MARKET：云市场分配  暂不支持MARKET
	Creator *AppInfoWithBindNumCreator `json:"creator,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// APP的key
	AppKey *string `json:"app_key,omitempty"`

	// 密钥
	AppSecret *string `json:"app_secret,omitempty"`

	// 注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// 状态   - 1： 有效
	Status *AppInfoWithBindNumStatus `json:"status,omitempty"`

	// APP的类型： - apig：存量apig应用，不推荐使用 - roma：roma集成应用
	AppType *AppInfoWithBindNumAppType `json:"app_type,omitempty"`

	// ROMA_APP的类型： - subscription：订阅应用 - integration：集成应用
	RomaAppType *string `json:"roma_app_type,omitempty"`

	// 绑定的API数量
	BindNum *int32 `json:"bind_num,omitempty"`
}

func (o AppInfoWithBindNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppInfoWithBindNum struct{}"
	}

	return strings.Join([]string{"AppInfoWithBindNum", string(data)}, " ")
}

type AppInfoWithBindNumCreator struct {
	value string
}

type AppInfoWithBindNumCreatorEnum struct {
	USER   AppInfoWithBindNumCreator
	MARKET AppInfoWithBindNumCreator
}

func GetAppInfoWithBindNumCreatorEnum() AppInfoWithBindNumCreatorEnum {
	return AppInfoWithBindNumCreatorEnum{
		USER: AppInfoWithBindNumCreator{
			value: "USER",
		},
		MARKET: AppInfoWithBindNumCreator{
			value: "MARKET",
		},
	}
}

func (c AppInfoWithBindNumCreator) Value() string {
	return c.value
}

func (c AppInfoWithBindNumCreator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppInfoWithBindNumCreator) UnmarshalJSON(b []byte) error {
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

type AppInfoWithBindNumStatus struct {
	value int32
}

type AppInfoWithBindNumStatusEnum struct {
	E_1 AppInfoWithBindNumStatus
}

func GetAppInfoWithBindNumStatusEnum() AppInfoWithBindNumStatusEnum {
	return AppInfoWithBindNumStatusEnum{
		E_1: AppInfoWithBindNumStatus{
			value: 1,
		},
	}
}

func (c AppInfoWithBindNumStatus) Value() int32 {
	return c.value
}

func (c AppInfoWithBindNumStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppInfoWithBindNumStatus) UnmarshalJSON(b []byte) error {
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

type AppInfoWithBindNumAppType struct {
	value string
}

type AppInfoWithBindNumAppTypeEnum struct {
	APIG AppInfoWithBindNumAppType
	ROMA AppInfoWithBindNumAppType
}

func GetAppInfoWithBindNumAppTypeEnum() AppInfoWithBindNumAppTypeEnum {
	return AppInfoWithBindNumAppTypeEnum{
		APIG: AppInfoWithBindNumAppType{
			value: "apig",
		},
		ROMA: AppInfoWithBindNumAppType{
			value: "roma",
		},
	}
}

func (c AppInfoWithBindNumAppType) Value() string {
	return c.value
}

func (c AppInfoWithBindNumAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppInfoWithBindNumAppType) UnmarshalJSON(b []byte) error {
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
