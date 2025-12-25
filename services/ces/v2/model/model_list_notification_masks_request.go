package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNotificationMasksRequest Request Object
type ListNotificationMasksRequest struct {

	// **参数解释**： 分页偏移量 **约束限制**： 不涉及 **取值范围**： 整数，[0,10000] **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小 **约束限制**： 不涉及 **取值范围**： 整数，[1,100] **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 排序关键字，与sort_dir同时使用。 **约束限制**： 目前只支持create_time与update_time **取值范围**： - create_time：按创建时间排序 - update_time：按修改时间排序 **默认取值**： 不涉及。
	SortKey *ListNotificationMasksRequestSortKey `json:"sort_key,omitempty"`

	// **参数解释**： 排序顺序，与sort_key同时使用。 **约束限制**： 不涉及。 **取值范围**： - DESC：降序排序。 - ASC：升序排序。 **默认取值**： 不涉及。
	SortDir *ListNotificationMasksRequestSortDir `json:"sort_dir,omitempty"`

	Body *ListNotificationMaskRequestBody `json:"body,omitempty"`
}

func (o ListNotificationMasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationMasksRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationMasksRequest", string(data)}, " ")
}

type ListNotificationMasksRequestSortKey struct {
	value string
}

type ListNotificationMasksRequestSortKeyEnum struct {
	CREATE_TIME ListNotificationMasksRequestSortKey
	UPDATE_TIME ListNotificationMasksRequestSortKey
}

func GetListNotificationMasksRequestSortKeyEnum() ListNotificationMasksRequestSortKeyEnum {
	return ListNotificationMasksRequestSortKeyEnum{
		CREATE_TIME: ListNotificationMasksRequestSortKey{
			value: "create_time",
		},
		UPDATE_TIME: ListNotificationMasksRequestSortKey{
			value: "update_time",
		},
	}
}

func (c ListNotificationMasksRequestSortKey) Value() string {
	return c.value
}

func (c ListNotificationMasksRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationMasksRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListNotificationMasksRequestSortDir struct {
	value string
}

type ListNotificationMasksRequestSortDirEnum struct {
	ASC  ListNotificationMasksRequestSortDir
	DESC ListNotificationMasksRequestSortDir
}

func GetListNotificationMasksRequestSortDirEnum() ListNotificationMasksRequestSortDirEnum {
	return ListNotificationMasksRequestSortDirEnum{
		ASC: ListNotificationMasksRequestSortDir{
			value: "ASC",
		},
		DESC: ListNotificationMasksRequestSortDir{
			value: "DESC",
		},
	}
}

func (c ListNotificationMasksRequestSortDir) Value() string {
	return c.value
}

func (c ListNotificationMasksRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationMasksRequestSortDir) UnmarshalJSON(b []byte) error {
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
