package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// Response Object
type ShowDetailsOfAppV2Response struct {
	// 编号

	Id *string `json:"id,omitempty"`
	// 名称

	Name *string `json:"name,omitempty"`
	// 描述

	Remark *string `json:"remark,omitempty"`
	// APP的创建者 - USER：用户自行创建 - MARKET：云市场分配  暂不支持MARKET

	Creator *ShowDetailsOfAppV2ResponseCreator `json:"creator,omitempty"`
	// 更新时间

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
	// APP的key

	AppKey *string `json:"app_key,omitempty"`
	// 密钥

	AppSecret *string `json:"app_secret,omitempty"`
	// 注册时间

	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`
	// 状态   - 1： 有效

	Status *ShowDetailsOfAppV2ResponseStatus `json:"status,omitempty"`
	// APP的类型： - apig：存量apig应用，不推荐使用 - roma：roma集成应用  默认apig，暂不支持其他类型

	AppType *ShowDetailsOfAppV2ResponseAppType `json:"app_type,omitempty"`
	// ROMA_APP的类型： - subscription：订阅应用 - integration：集成应用  暂不支持

	RomaAppType    *string `json:"roma_app_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailsOfAppV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAppV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAppV2Response", string(data)}, " ")
}

type ShowDetailsOfAppV2ResponseCreator struct {
	value string
}

type ShowDetailsOfAppV2ResponseCreatorEnum struct {
	USER   ShowDetailsOfAppV2ResponseCreator
	MARKET ShowDetailsOfAppV2ResponseCreator
}

func GetShowDetailsOfAppV2ResponseCreatorEnum() ShowDetailsOfAppV2ResponseCreatorEnum {
	return ShowDetailsOfAppV2ResponseCreatorEnum{
		USER: ShowDetailsOfAppV2ResponseCreator{
			value: "USER",
		},
		MARKET: ShowDetailsOfAppV2ResponseCreator{
			value: "MARKET",
		},
	}
}

func (c ShowDetailsOfAppV2ResponseCreator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfAppV2ResponseCreator) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfAppV2ResponseStatus struct {
	value int32
}

type ShowDetailsOfAppV2ResponseStatusEnum struct {
	E_1 ShowDetailsOfAppV2ResponseStatus
}

func GetShowDetailsOfAppV2ResponseStatusEnum() ShowDetailsOfAppV2ResponseStatusEnum {
	return ShowDetailsOfAppV2ResponseStatusEnum{
		E_1: ShowDetailsOfAppV2ResponseStatus{
			value: 1,
		},
	}
}

func (c ShowDetailsOfAppV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfAppV2ResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfAppV2ResponseAppType struct {
	value string
}

type ShowDetailsOfAppV2ResponseAppTypeEnum struct {
	APIG ShowDetailsOfAppV2ResponseAppType
	ROMA ShowDetailsOfAppV2ResponseAppType
}

func GetShowDetailsOfAppV2ResponseAppTypeEnum() ShowDetailsOfAppV2ResponseAppTypeEnum {
	return ShowDetailsOfAppV2ResponseAppTypeEnum{
		APIG: ShowDetailsOfAppV2ResponseAppType{
			value: "apig",
		},
		ROMA: ShowDetailsOfAppV2ResponseAppType{
			value: "roma",
		},
	}
}

func (c ShowDetailsOfAppV2ResponseAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfAppV2ResponseAppType) UnmarshalJSON(b []byte) error {
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
