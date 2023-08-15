package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LdApiCheckInfo struct {

	// 校验类型：   - path：路径类型   - name：名称类型
	Type LdApiCheckInfoType `json:"type"`

	// 自定义后端API名称。  type = name时必填
	LdApiName *string `json:"ld_api_name,omitempty"`

	// 自定义后端API请求方式。  type = path时必填
	LdApiMethod *LdApiCheckInfoLdApiMethod `json:"ld_api_method,omitempty"`

	// 自定义后端API的访问地址。  type = path时必填
	LdApiPath *string `json:"ld_api_path,omitempty"`

	// 集成应用ID。  校验应用下后端API定义是否重复时必填
	RomaAppId *string `json:"roma_app_id,omitempty"`
}

func (o LdApiCheckInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiCheckInfo struct{}"
	}

	return strings.Join([]string{"LdApiCheckInfo", string(data)}, " ")
}

type LdApiCheckInfoType struct {
	value string
}

type LdApiCheckInfoTypeEnum struct {
	PATH LdApiCheckInfoType
	NAME LdApiCheckInfoType
}

func GetLdApiCheckInfoTypeEnum() LdApiCheckInfoTypeEnum {
	return LdApiCheckInfoTypeEnum{
		PATH: LdApiCheckInfoType{
			value: "path",
		},
		NAME: LdApiCheckInfoType{
			value: "name",
		},
	}
}

func (c LdApiCheckInfoType) Value() string {
	return c.value
}

func (c LdApiCheckInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiCheckInfoType) UnmarshalJSON(b []byte) error {
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

type LdApiCheckInfoLdApiMethod struct {
	value string
}

type LdApiCheckInfoLdApiMethodEnum struct {
	GET    LdApiCheckInfoLdApiMethod
	POST   LdApiCheckInfoLdApiMethod
	PUT    LdApiCheckInfoLdApiMethod
	DELETE LdApiCheckInfoLdApiMethod
}

func GetLdApiCheckInfoLdApiMethodEnum() LdApiCheckInfoLdApiMethodEnum {
	return LdApiCheckInfoLdApiMethodEnum{
		GET: LdApiCheckInfoLdApiMethod{
			value: "GET",
		},
		POST: LdApiCheckInfoLdApiMethod{
			value: "POST",
		},
		PUT: LdApiCheckInfoLdApiMethod{
			value: "PUT",
		},
		DELETE: LdApiCheckInfoLdApiMethod{
			value: "DELETE",
		},
	}
}

func (c LdApiCheckInfoLdApiMethod) Value() string {
	return c.value
}

func (c LdApiCheckInfoLdApiMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiCheckInfoLdApiMethod) UnmarshalJSON(b []byte) error {
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
