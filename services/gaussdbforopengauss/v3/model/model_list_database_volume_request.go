package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDatabaseVolumeRequest Request Object
type ListDatabaseVolumeRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListDatabaseVolumeRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 数据库名称。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**: 数据库的缺省表空间名称。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	TableSpaceName *string `json:"table_space_name,omitempty"`

	// **参数解释**: 表所属用户名称。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 不涉及 **取值范围**: 0 - 10000 **默认取值**: 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数 **约束限制**: 不能为负数 **取值范围**: 最小值为1，最大值为200。 **默认取值**: 100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabaseVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseVolumeRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseVolumeRequest", string(data)}, " ")
}

type ListDatabaseVolumeRequestXLanguage struct {
	value string
}

type ListDatabaseVolumeRequestXLanguageEnum struct {
	ZH_CN ListDatabaseVolumeRequestXLanguage
	EN_US ListDatabaseVolumeRequestXLanguage
}

func GetListDatabaseVolumeRequestXLanguageEnum() ListDatabaseVolumeRequestXLanguageEnum {
	return ListDatabaseVolumeRequestXLanguageEnum{
		ZH_CN: ListDatabaseVolumeRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListDatabaseVolumeRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListDatabaseVolumeRequestXLanguage) Value() string {
	return c.value
}

func (c ListDatabaseVolumeRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatabaseVolumeRequestXLanguage) UnmarshalJSON(b []byte) error {
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
