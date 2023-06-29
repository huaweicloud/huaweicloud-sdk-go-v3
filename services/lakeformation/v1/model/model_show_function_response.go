package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowFunctionResponse Response Object
type ShowFunctionResponse struct {

	// catalog名字
	CatalogName *string `json:"catalog_name,omitempty"`

	// 数据库名字
	DatabaseName *string `json:"database_name,omitempty"`

	// 函数名
	FunctionName *string `json:"function_name,omitempty"`

	// 函数类型
	FunctionType *ShowFunctionResponseFunctionType `json:"function_type,omitempty"`

	// 函数所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者类型
	OwnerType *ShowFunctionResponseOwnerType `json:"owner_type,omitempty"`

	// 函数类名
	ClassName *string `json:"class_name,omitempty"`

	// 创建时间格式为yyyy-mm-ddThh:mm:sss
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 函数地址信息
	ResourceUris   *[]FunctionResourceUri `json:"resource_uris,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionResponse struct{}"
	}

	return strings.Join([]string{"ShowFunctionResponse", string(data)}, " ")
}

type ShowFunctionResponseFunctionType struct {
	value string
}

type ShowFunctionResponseFunctionTypeEnum struct {
	JAVA ShowFunctionResponseFunctionType
}

func GetShowFunctionResponseFunctionTypeEnum() ShowFunctionResponseFunctionTypeEnum {
	return ShowFunctionResponseFunctionTypeEnum{
		JAVA: ShowFunctionResponseFunctionType{
			value: "JAVA",
		},
	}
}

func (c ShowFunctionResponseFunctionType) Value() string {
	return c.value
}

func (c ShowFunctionResponseFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionResponseFunctionType) UnmarshalJSON(b []byte) error {
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

type ShowFunctionResponseOwnerType struct {
	value string
}

type ShowFunctionResponseOwnerTypeEnum struct {
	USER  ShowFunctionResponseOwnerType
	GROUP ShowFunctionResponseOwnerType
	ROLE  ShowFunctionResponseOwnerType
}

func GetShowFunctionResponseOwnerTypeEnum() ShowFunctionResponseOwnerTypeEnum {
	return ShowFunctionResponseOwnerTypeEnum{
		USER: ShowFunctionResponseOwnerType{
			value: "USER",
		},
		GROUP: ShowFunctionResponseOwnerType{
			value: "GROUP",
		},
		ROLE: ShowFunctionResponseOwnerType{
			value: "ROLE",
		},
	}
}

func (c ShowFunctionResponseOwnerType) Value() string {
	return c.value
}

func (c ShowFunctionResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionResponseOwnerType) UnmarshalJSON(b []byte) error {
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
