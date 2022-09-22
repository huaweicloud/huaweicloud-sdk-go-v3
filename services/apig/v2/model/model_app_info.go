package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type AppInfo struct {

	// 编号
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Remark *string `json:"remark,omitempty"`

	// APP的创建者 - USER：用户自行创建 - MARKET：云商店分配  暂不支持MARKET
	Creator *AppInfoCreator `json:"creator,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// APP的key
	AppKey *string `json:"app_key,omitempty"`

	// 密钥
	AppSecret *string `json:"app_secret,omitempty"`

	// 注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// 状态   - 1： 有效
	Status *AppInfoStatus `json:"status,omitempty"`

	// APP的类型： - apig：存量apig应用，不推荐使用 - roma：roma集成应用  默认apig，暂不支持其他类型
	AppType *AppInfoAppType `json:"app_type,omitempty"`

	// ROMA_APP的类型： - subscription：订阅应用 - integration：集成应用  暂不支持
	RomaAppType *string `json:"roma_app_type,omitempty"`
}

func (o AppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppInfo struct{}"
	}

	return strings.Join([]string{"AppInfo", string(data)}, " ")
}

type AppInfoCreator struct {
	value string
}

type AppInfoCreatorEnum struct {
	USER   AppInfoCreator
	MARKET AppInfoCreator
}

func GetAppInfoCreatorEnum() AppInfoCreatorEnum {
	return AppInfoCreatorEnum{
		USER: AppInfoCreator{
			value: "USER",
		},
		MARKET: AppInfoCreator{
			value: "MARKET",
		},
	}
}

func (c AppInfoCreator) Value() string {
	return c.value
}

func (c AppInfoCreator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppInfoCreator) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type AppInfoStatus struct {
	value int32
}

type AppInfoStatusEnum struct {
	E_1 AppInfoStatus
}

func GetAppInfoStatusEnum() AppInfoStatusEnum {
	return AppInfoStatusEnum{
		E_1: AppInfoStatus{
			value: 1,
		},
	}
}

func (c AppInfoStatus) Value() int32 {
	return c.value
}

func (c AppInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppInfoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AppInfoAppType struct {
	value string
}

type AppInfoAppTypeEnum struct {
	APIG AppInfoAppType
	ROMA AppInfoAppType
}

func GetAppInfoAppTypeEnum() AppInfoAppTypeEnum {
	return AppInfoAppTypeEnum{
		APIG: AppInfoAppType{
			value: "apig",
		},
		ROMA: AppInfoAppType{
			value: "roma",
		},
	}
}

func (c AppInfoAppType) Value() string {
	return c.value
}

func (c AppInfoAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppInfoAppType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
