package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWdrSnapshotAvailableGroupsRequest Request Object
type ListWdrSnapshotAvailableGroupsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn - en-us  **默认取值**: en-us
	XLanguage *ListWdrSnapshotAvailableGroupsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 开始时间。 **约束限制**: 不涉及。 **取值范围**: 格式为“yyyy-mm-ddThh:mm:ssZ”。注意，对于时区的+号，需要进行编码，替换为%2B。 **默认取值**: 不涉及。
	BeginTime string `json:"begin_time"`

	// **参数解释**: 结束时间。 **约束限制**: 不涉及。 **取值范围**: 格式为“yyyy-mm-ddThh:mm:ssZ”。对于时区的+号，需要进行编码，替换为%2B。 **默认取值**: 不涉及。
	EndTime string `json:"end_time"`
}

func (o ListWdrSnapshotAvailableGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWdrSnapshotAvailableGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListWdrSnapshotAvailableGroupsRequest", string(data)}, " ")
}

type ListWdrSnapshotAvailableGroupsRequestXLanguage struct {
	value string
}

type ListWdrSnapshotAvailableGroupsRequestXLanguageEnum struct {
	ZH_CN ListWdrSnapshotAvailableGroupsRequestXLanguage
	EN_US ListWdrSnapshotAvailableGroupsRequestXLanguage
}

func GetListWdrSnapshotAvailableGroupsRequestXLanguageEnum() ListWdrSnapshotAvailableGroupsRequestXLanguageEnum {
	return ListWdrSnapshotAvailableGroupsRequestXLanguageEnum{
		ZH_CN: ListWdrSnapshotAvailableGroupsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListWdrSnapshotAvailableGroupsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListWdrSnapshotAvailableGroupsRequestXLanguage) Value() string {
	return c.value
}

func (c ListWdrSnapshotAvailableGroupsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWdrSnapshotAvailableGroupsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
