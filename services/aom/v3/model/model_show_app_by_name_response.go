package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAppByNameResponse Response Object
type ShowAppByNameResponse struct {

	// aomId，如果为空则不显示
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
	RegisterType   *ShowAppByNameResponseRegisterType `json:"register_type,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ShowAppByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppByNameResponse struct{}"
	}

	return strings.Join([]string{"ShowAppByNameResponse", string(data)}, " ")
}

type ShowAppByNameResponseRegisterType struct {
	value string
}

type ShowAppByNameResponseRegisterTypeEnum struct {
	API               ShowAppByNameResponseRegisterType
	CONSOLE           ShowAppByNameResponseRegisterType
	SERVICE_DISCOVERY ShowAppByNameResponseRegisterType
}

func GetShowAppByNameResponseRegisterTypeEnum() ShowAppByNameResponseRegisterTypeEnum {
	return ShowAppByNameResponseRegisterTypeEnum{
		API: ShowAppByNameResponseRegisterType{
			value: "API",
		},
		CONSOLE: ShowAppByNameResponseRegisterType{
			value: "CONSOLE",
		},
		SERVICE_DISCOVERY: ShowAppByNameResponseRegisterType{
			value: "SERVICE_DISCOVERY",
		},
	}
}

func (c ShowAppByNameResponseRegisterType) Value() string {
	return c.value
}

func (c ShowAppByNameResponseRegisterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppByNameResponseRegisterType) UnmarshalJSON(b []byte) error {
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
