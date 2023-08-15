package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAppResponse Response Object
type ShowAppResponse struct {

	// aomId
	AomId *string `json:"aom_id,omitempty"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 应用名称
	DisplayName *string `json:"display_name,omitempty"`

	// 企业项目id
	EpsId *string `json:"eps_id,omitempty"`

	// 修改时间
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 修改人
	Modifier *string `json:"modifier,omitempty"`

	// 唯一标识
	Name *string `json:"name,omitempty"`

	// 注册方式
	RegisterType   *ShowAppResponseRegisterType `json:"register_type,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppResponse struct{}"
	}

	return strings.Join([]string{"ShowAppResponse", string(data)}, " ")
}

type ShowAppResponseRegisterType struct {
	value string
}

type ShowAppResponseRegisterTypeEnum struct {
	API               ShowAppResponseRegisterType
	CONSOLE           ShowAppResponseRegisterType
	SERVICE_DISCOVERY ShowAppResponseRegisterType
}

func GetShowAppResponseRegisterTypeEnum() ShowAppResponseRegisterTypeEnum {
	return ShowAppResponseRegisterTypeEnum{
		API: ShowAppResponseRegisterType{
			value: "API",
		},
		CONSOLE: ShowAppResponseRegisterType{
			value: "CONSOLE",
		},
		SERVICE_DISCOVERY: ShowAppResponseRegisterType{
			value: "SERVICE_DISCOVERY",
		},
	}
}

func (c ShowAppResponseRegisterType) Value() string {
	return c.value
}

func (c ShowAppResponseRegisterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppResponseRegisterType) UnmarshalJSON(b []byte) error {
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
