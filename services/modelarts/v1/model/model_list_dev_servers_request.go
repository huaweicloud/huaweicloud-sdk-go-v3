package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDevServersRequest Request Object
type ListDevServersRequest struct {

	// **参数解释**：实例归属的用户ID。 **约束限制**：可选。 **取值范围**：1 - 64字符，小写字母、数字和中划线。在大账号/有admin权限场景下生效，值通常为当前登录用户ID。 **默认取值**：不涉及。
	Owner *string `json:"owner,omitempty"`

	// **参数解释**：排序方式。 **约束限制**：可选。 **取值范围**： - ASC：升序 - DESC：降序 **默认取值**：ASC。
	SortDir *ListDevServersRequestSortDir `json:"sort_dir,omitempty"`

	// **参数解释**：排序字段。 **约束限制**：可选。 **取值范围**： - createTime：默认值，创建时间。 - updateTime：更新时间。 **默认取值**：createTime。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：每一页的数量。 **约束限制**：可选。 **取值范围**：0 - 1024 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页记录的起始位置偏移量。 **约束限制**：可选。 **取值范围**：0 - 2147483647 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDevServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevServersRequest struct{}"
	}

	return strings.Join([]string{"ListDevServersRequest", string(data)}, " ")
}

type ListDevServersRequestSortDir struct {
	value string
}

type ListDevServersRequestSortDirEnum struct {
	ASC  ListDevServersRequestSortDir
	DESC ListDevServersRequestSortDir
}

func GetListDevServersRequestSortDirEnum() ListDevServersRequestSortDirEnum {
	return ListDevServersRequestSortDirEnum{
		ASC: ListDevServersRequestSortDir{
			value: "ASC",
		},
		DESC: ListDevServersRequestSortDir{
			value: "DESC",
		},
	}
}

func (c ListDevServersRequestSortDir) Value() string {
	return c.value
}

func (c ListDevServersRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDevServersRequestSortDir) UnmarshalJSON(b []byte) error {
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
