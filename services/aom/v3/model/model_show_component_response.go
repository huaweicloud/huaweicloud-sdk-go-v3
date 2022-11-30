package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowComponentResponse struct {

	// aomId
	AomId *string `json:"aom_id,omitempty"`

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者
	Creator *string `json:"creator,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 组件Id
	Id *string `json:"id,omitempty"`

	// 修改时间
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 修改者
	Modifier *string `json:"modifier,omitempty"`

	// 组件名称
	Name *string `json:"name,omitempty"`

	// 注册方式
	RegisterType *ShowComponentResponseRegisterType `json:"register_type,omitempty"`

	// 子应用id
	SubAppId       *string `json:"sub_app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComponentResponse struct{}"
	}

	return strings.Join([]string{"ShowComponentResponse", string(data)}, " ")
}

type ShowComponentResponseRegisterType struct {
	value string
}

type ShowComponentResponseRegisterTypeEnum struct {
	API               ShowComponentResponseRegisterType
	CONSOLE           ShowComponentResponseRegisterType
	SERVICE_DISCOVERY ShowComponentResponseRegisterType
}

func GetShowComponentResponseRegisterTypeEnum() ShowComponentResponseRegisterTypeEnum {
	return ShowComponentResponseRegisterTypeEnum{
		API: ShowComponentResponseRegisterType{
			value: "API",
		},
		CONSOLE: ShowComponentResponseRegisterType{
			value: "CONSOLE",
		},
		SERVICE_DISCOVERY: ShowComponentResponseRegisterType{
			value: "SERVICE_DISCOVERY",
		},
	}
}

func (c ShowComponentResponseRegisterType) Value() string {
	return c.value
}

func (c ShowComponentResponseRegisterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowComponentResponseRegisterType) UnmarshalJSON(b []byte) error {
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
