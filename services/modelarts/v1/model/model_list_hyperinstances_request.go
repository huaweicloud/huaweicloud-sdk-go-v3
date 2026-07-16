package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHyperinstancesRequest Request Object
type ListHyperinstancesRequest struct {

	// **参数解释**：排序方式。 **约束限制**：不涉及。 **取值范围**：枚举值如下：  - ASC升序。  - DESC降序。 **默认取值**：不涉及。
	SortDir *ListHyperinstancesRequestSortDir `json:"sort_dir,omitempty"`

	// **参数解释**：排序字段。 **约束限制**：不涉及。 **取值范围**：枚举值如下：  - createTime：默认值，创建时间。  - updateTime：更新时间。 **默认取值**：不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：每一页的数量。 **约束限制**：不涉及。 **取值范围**：[1,1024]。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页记录的起始位置偏移量。 **约束限制**：不涉及。 **取值范围**：[0,2147483647]。 **默认取值**：不涉及。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListHyperinstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperinstancesRequest struct{}"
	}

	return strings.Join([]string{"ListHyperinstancesRequest", string(data)}, " ")
}

type ListHyperinstancesRequestSortDir struct {
	value string
}

type ListHyperinstancesRequestSortDirEnum struct {
	ASC  ListHyperinstancesRequestSortDir
	DESC ListHyperinstancesRequestSortDir
}

func GetListHyperinstancesRequestSortDirEnum() ListHyperinstancesRequestSortDirEnum {
	return ListHyperinstancesRequestSortDirEnum{
		ASC: ListHyperinstancesRequestSortDir{
			value: "ASC",
		},
		DESC: ListHyperinstancesRequestSortDir{
			value: "DESC",
		},
	}
}

func (c ListHyperinstancesRequestSortDir) Value() string {
	return c.value
}

func (c ListHyperinstancesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHyperinstancesRequestSortDir) UnmarshalJSON(b []byte) error {
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
