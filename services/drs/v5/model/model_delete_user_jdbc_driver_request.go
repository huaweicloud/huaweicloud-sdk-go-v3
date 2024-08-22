package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteUserJdbcDriverRequest Request Object
type DeleteUserJdbcDriverRequest struct {

	// 请求语言类型。
	XLanguage *DeleteUserJdbcDriverRequestXLanguage `json:"X-Language,omitempty"`

	Body *DeleteUserDriverReq `json:"body,omitempty"`
}

func (o DeleteUserJdbcDriverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserJdbcDriverRequest struct{}"
	}

	return strings.Join([]string{"DeleteUserJdbcDriverRequest", string(data)}, " ")
}

type DeleteUserJdbcDriverRequestXLanguage struct {
	value string
}

type DeleteUserJdbcDriverRequestXLanguageEnum struct {
	EN_US DeleteUserJdbcDriverRequestXLanguage
	ZH_CN DeleteUserJdbcDriverRequestXLanguage
}

func GetDeleteUserJdbcDriverRequestXLanguageEnum() DeleteUserJdbcDriverRequestXLanguageEnum {
	return DeleteUserJdbcDriverRequestXLanguageEnum{
		EN_US: DeleteUserJdbcDriverRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteUserJdbcDriverRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteUserJdbcDriverRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteUserJdbcDriverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteUserJdbcDriverRequestXLanguage) UnmarshalJSON(b []byte) error {
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
