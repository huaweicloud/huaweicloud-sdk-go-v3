package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListModulesRequest Request Object
type ListModulesRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 分页
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 数据量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序字段， - create_time：创建时间 - update_time：更新时间  **约束限制：** 不涉及 **取值范围：** - create_time - update_time  **默认取值：** create_time
	SortKey *ListModulesRequestSortKey `json:"sort_key,omitempty"`

	// **参数解释：** 排序顺序 - ASC：升序 - DESC：降序  **约束限制：** 不涉及 **取值范围：** - ASC：升序 - DESC：降序  **默认取值：** DESC
	SortDir *string `json:"sort_dir,omitempty"`

	// 模块类型，tab/section
	ModuleType *string `json:"module_type,omitempty"`
}

func (o ListModulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModulesRequest struct{}"
	}

	return strings.Join([]string{"ListModulesRequest", string(data)}, " ")
}

type ListModulesRequestSortKey struct {
	value string
}

type ListModulesRequestSortKeyEnum struct {
	UPDATE_TIME ListModulesRequestSortKey
	CREATE_TIME ListModulesRequestSortKey
}

func GetListModulesRequestSortKeyEnum() ListModulesRequestSortKeyEnum {
	return ListModulesRequestSortKeyEnum{
		UPDATE_TIME: ListModulesRequestSortKey{
			value: "update_time",
		},
		CREATE_TIME: ListModulesRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListModulesRequestSortKey) Value() string {
	return c.value
}

func (c ListModulesRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListModulesRequestSortKey) UnmarshalJSON(b []byte) error {
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
