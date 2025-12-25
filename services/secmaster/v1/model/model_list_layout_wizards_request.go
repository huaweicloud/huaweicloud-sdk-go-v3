package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLayoutWizardsRequest Request Object
type ListLayoutWizardsRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 布局ID
	LayoutId string `json:"layout_id"`

	// 分页
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 数据量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序字段， - create_time：创建时间 - update_time：更新时间  **约束限制：** 不涉及 **取值范围：** - create_time - update_time  **默认取值：** create_time
	SortKey *ListLayoutWizardsRequestSortKey `json:"sort_key,omitempty"`

	// **参数解释：** 排序顺序 - ASC：升序 - DESC：降序  **约束限制：** 不涉及 **取值范围：** - ASC：升序 - DESC：降序  **默认取值：** DESC
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListLayoutWizardsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLayoutWizardsRequest struct{}"
	}

	return strings.Join([]string{"ListLayoutWizardsRequest", string(data)}, " ")
}

type ListLayoutWizardsRequestSortKey struct {
	value string
}

type ListLayoutWizardsRequestSortKeyEnum struct {
	UPDATE_TIME ListLayoutWizardsRequestSortKey
	CREATE_TIME ListLayoutWizardsRequestSortKey
}

func GetListLayoutWizardsRequestSortKeyEnum() ListLayoutWizardsRequestSortKeyEnum {
	return ListLayoutWizardsRequestSortKeyEnum{
		UPDATE_TIME: ListLayoutWizardsRequestSortKey{
			value: "update_time",
		},
		CREATE_TIME: ListLayoutWizardsRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListLayoutWizardsRequestSortKey) Value() string {
	return c.value
}

func (c ListLayoutWizardsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLayoutWizardsRequestSortKey) UnmarshalJSON(b []byte) error {
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
