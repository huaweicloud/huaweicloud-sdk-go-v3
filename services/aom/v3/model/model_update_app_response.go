package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateAppResponse struct {

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
	RegisterType   *UpdateAppResponseRegisterType `json:"register_type,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o UpdateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppResponse", string(data)}, " ")
}

type UpdateAppResponseRegisterType struct {
	value string
}

type UpdateAppResponseRegisterTypeEnum struct {
	API               UpdateAppResponseRegisterType
	CONSOLE           UpdateAppResponseRegisterType
	SERVICE_DISCOVERY UpdateAppResponseRegisterType
}

func GetUpdateAppResponseRegisterTypeEnum() UpdateAppResponseRegisterTypeEnum {
	return UpdateAppResponseRegisterTypeEnum{
		API: UpdateAppResponseRegisterType{
			value: "API",
		},
		CONSOLE: UpdateAppResponseRegisterType{
			value: "CONSOLE",
		},
		SERVICE_DISCOVERY: UpdateAppResponseRegisterType{
			value: "SERVICE_DISCOVERY",
		},
	}
}

func (c UpdateAppResponseRegisterType) Value() string {
	return c.value
}

func (c UpdateAppResponseRegisterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAppResponseRegisterType) UnmarshalJSON(b []byte) error {
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
