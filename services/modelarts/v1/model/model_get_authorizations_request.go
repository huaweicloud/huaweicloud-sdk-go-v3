package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetAuthorizationsRequest Request Object
type GetAuthorizationsRequest struct {

	// **参数解释**：指定排序字段。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - user_name：IAM用户名称。 - create_time：创建时间。 **默认取值**：user_name。
	SortBy *GetAuthorizationsRequestSortBy `json:"sort_by,omitempty"`

	// **参数解释**：排序方式。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - ASC：递增排序。 - DESC：递减排序。 **默认取值**：ASC。
	Order *GetAuthorizationsRequestOrder `json:"order,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：[1,1000]。 **默认取值**：1000。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页列表的起始页。 **约束限制**：不涉及。 **取值范围**：非负整数。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o GetAuthorizationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAuthorizationsRequest struct{}"
	}

	return strings.Join([]string{"GetAuthorizationsRequest", string(data)}, " ")
}

type GetAuthorizationsRequestSortBy struct {
	value string
}

type GetAuthorizationsRequestSortByEnum struct {
	USER_NAME   GetAuthorizationsRequestSortBy
	CREATE_TIME GetAuthorizationsRequestSortBy
}

func GetGetAuthorizationsRequestSortByEnum() GetAuthorizationsRequestSortByEnum {
	return GetAuthorizationsRequestSortByEnum{
		USER_NAME: GetAuthorizationsRequestSortBy{
			value: "user_name",
		},
		CREATE_TIME: GetAuthorizationsRequestSortBy{
			value: "create_time",
		},
	}
}

func (c GetAuthorizationsRequestSortBy) Value() string {
	return c.value
}

func (c GetAuthorizationsRequestSortBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetAuthorizationsRequestSortBy) UnmarshalJSON(b []byte) error {
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

type GetAuthorizationsRequestOrder struct {
	value string
}

type GetAuthorizationsRequestOrderEnum struct {
	ASC  GetAuthorizationsRequestOrder
	DESC GetAuthorizationsRequestOrder
}

func GetGetAuthorizationsRequestOrderEnum() GetAuthorizationsRequestOrderEnum {
	return GetAuthorizationsRequestOrderEnum{
		ASC: GetAuthorizationsRequestOrder{
			value: "asc",
		},
		DESC: GetAuthorizationsRequestOrder{
			value: "desc",
		},
	}
}

func (c GetAuthorizationsRequestOrder) Value() string {
	return c.value
}

func (c GetAuthorizationsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetAuthorizationsRequestOrder) UnmarshalJSON(b []byte) error {
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
