package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteJdbcDriverRequest Request Object
type DeleteJdbcDriverRequest struct {

	// 请求语言类型。
	XLanguage *DeleteJdbcDriverRequestXLanguage `json:"X-Language,omitempty"`

	Body *DeleteDriverReq `json:"body,omitempty"`
}

func (o DeleteJdbcDriverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJdbcDriverRequest struct{}"
	}

	return strings.Join([]string{"DeleteJdbcDriverRequest", string(data)}, " ")
}

type DeleteJdbcDriverRequestXLanguage struct {
	value string
}

type DeleteJdbcDriverRequestXLanguageEnum struct {
	EN_US DeleteJdbcDriverRequestXLanguage
	ZH_CN DeleteJdbcDriverRequestXLanguage
}

func GetDeleteJdbcDriverRequestXLanguageEnum() DeleteJdbcDriverRequestXLanguageEnum {
	return DeleteJdbcDriverRequestXLanguageEnum{
		EN_US: DeleteJdbcDriverRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteJdbcDriverRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteJdbcDriverRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteJdbcDriverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteJdbcDriverRequestXLanguage) UnmarshalJSON(b []byte) error {
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
