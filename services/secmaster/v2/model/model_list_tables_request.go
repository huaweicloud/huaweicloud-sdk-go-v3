package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTablesRequest Request Object
type ListTablesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 目录类型 - STREAMING 实时流 - INDEX 索引 - APPLICATION 应用 - TENANT_BUCKET 租户桶 - LAKE 数据湖  **约束限制** 不涉及 **取值范围**: - STREAMING - INDEX - APPLICATION - TENANT_BUCKET - LAKE  **默认值** 不涉及
	Category *ListTablesRequestCategory `json:"category,omitempty"`

	// 表id
	TableId *string `json:"table_id,omitempty"`

	// 表别名
	TableAlias *string `json:"table_alias,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// **参数解释：** 偏移量 **取值范围：** 0-1000000 **默认取值：** 0
	Offset *int64 `json:"offset,omitempty"`

	// **参数解释：** 查询数据限制 **取值范围：** 0-1000 **默认取值：** 不涉及
	Limit *int64 `json:"limit,omitempty"`

	// 按照属性排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序，支持 `ASC` 或 `DESC`。
	SortDir *string `json:"sort_dir,omitempty"`

	// 是否存在
	Exists *bool `json:"exists,omitempty"`
}

func (o ListTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesRequest struct{}"
	}

	return strings.Join([]string{"ListTablesRequest", string(data)}, " ")
}

type ListTablesRequestCategory struct {
	value string
}

type ListTablesRequestCategoryEnum struct {
	STREAMING     ListTablesRequestCategory
	INDEX         ListTablesRequestCategory
	APPLICATION   ListTablesRequestCategory
	TENANT_BUCKET ListTablesRequestCategory
	LAKE          ListTablesRequestCategory
}

func GetListTablesRequestCategoryEnum() ListTablesRequestCategoryEnum {
	return ListTablesRequestCategoryEnum{
		STREAMING: ListTablesRequestCategory{
			value: "STREAMING",
		},
		INDEX: ListTablesRequestCategory{
			value: "INDEX",
		},
		APPLICATION: ListTablesRequestCategory{
			value: "APPLICATION",
		},
		TENANT_BUCKET: ListTablesRequestCategory{
			value: "TENANT_BUCKET",
		},
		LAKE: ListTablesRequestCategory{
			value: "LAKE",
		},
	}
}

func (c ListTablesRequestCategory) Value() string {
	return c.value
}

func (c ListTablesRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTablesRequestCategory) UnmarshalJSON(b []byte) error {
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
