package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteDatabaseRequest Request Object
type DeleteDatabaseRequest struct {

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *DeleteDatabaseRequestXLanguage `json:"X-Language,omitempty"`

	// 数据库库名称。
	DatabaseName string `json:"database_name"`
}

func (o DeleteDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRequest", string(data)}, " ")
}

type DeleteDatabaseRequestXLanguage struct {
	value string
}

type DeleteDatabaseRequestXLanguageEnum struct {
	ZH_CN DeleteDatabaseRequestXLanguage
	EN_US DeleteDatabaseRequestXLanguage
}

func GetDeleteDatabaseRequestXLanguageEnum() DeleteDatabaseRequestXLanguageEnum {
	return DeleteDatabaseRequestXLanguageEnum{
		ZH_CN: DeleteDatabaseRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteDatabaseRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteDatabaseRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteDatabaseRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDatabaseRequestXLanguage) UnmarshalJSON(b []byte) error {
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
