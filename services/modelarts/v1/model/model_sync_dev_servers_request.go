package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SyncDevServersRequest Request Object
type SyncDevServersRequest struct {

	// **参数解释**：实例归属的用户ID。 **约束限制**：可选。 **取值范围**：1 - 64字符，小写字母、数字和中划线。在大账号/有admin权限场景下生效，值通常为当前登录用户ID。 **默认取值**：不涉及。
	Owner *string `json:"owner,omitempty"`

	// **参数解释**：排序方式。 **约束限制**：可选。 **取值范围**： - ASC：升序 - DESC：降序 **默认取值**：ASC。
	SortDir *SyncDevServersRequestSortDir `json:"sort_dir,omitempty"`

	// **参数解释**：排序字段。 **约束限制**：可选。 **取值范围**： - createTime：默认值，创建时间。 - updateTime：更新时间。 **默认取值**：createTime。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：分页记录的起始位置偏移量。 **约束限制**：可选。 **取值范围**：0 - 2147483647 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：每一页的数量。 **约束限制**：可选。 **取值范围**：0 - 1024 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o SyncDevServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncDevServersRequest struct{}"
	}

	return strings.Join([]string{"SyncDevServersRequest", string(data)}, " ")
}

type SyncDevServersRequestSortDir struct {
	value string
}

type SyncDevServersRequestSortDirEnum struct {
	ASC  SyncDevServersRequestSortDir
	DESC SyncDevServersRequestSortDir
}

func GetSyncDevServersRequestSortDirEnum() SyncDevServersRequestSortDirEnum {
	return SyncDevServersRequestSortDirEnum{
		ASC: SyncDevServersRequestSortDir{
			value: "ASC",
		},
		DESC: SyncDevServersRequestSortDir{
			value: "DESC",
		},
	}
}

func (c SyncDevServersRequestSortDir) Value() string {
	return c.value
}

func (c SyncDevServersRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncDevServersRequestSortDir) UnmarshalJSON(b []byte) error {
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
