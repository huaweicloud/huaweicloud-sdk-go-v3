package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTrainingExperimentsRequest Request Object
type ListTrainingExperimentsRequest struct {

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 返回的数据条目数。
	Limit *int32 `json:"limit,omitempty"`

	// 数据条目偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：排序依据字段，例如sort_by=update_time，则表示以条目的更新时间进行排序。 **约束限制**：不涉及。 **取值范围**： - update_time：更新时间。 - name：实验名称。 - create_time：创建时间。 **默认取值**：不涉及。
	SortBy *string `json:"sort_by,omitempty"`

	// 排序的方式。该字段必须与sort_by同时使用。 缺省值: desc 枚举值： - asc：表示升序排列， - desc：降序排列。
	Order *ListTrainingExperimentsRequestOrder `json:"order,omitempty"`
}

func (o ListTrainingExperimentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingExperimentsRequest struct{}"
	}

	return strings.Join([]string{"ListTrainingExperimentsRequest", string(data)}, " ")
}

type ListTrainingExperimentsRequestOrder struct {
	value string
}

type ListTrainingExperimentsRequestOrderEnum struct {
	DESC ListTrainingExperimentsRequestOrder
	ASC  ListTrainingExperimentsRequestOrder
}

func GetListTrainingExperimentsRequestOrderEnum() ListTrainingExperimentsRequestOrderEnum {
	return ListTrainingExperimentsRequestOrderEnum{
		DESC: ListTrainingExperimentsRequestOrder{
			value: "desc",
		},
		ASC: ListTrainingExperimentsRequestOrder{
			value: "asc",
		},
	}
}

func (c ListTrainingExperimentsRequestOrder) Value() string {
	return c.value
}

func (c ListTrainingExperimentsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTrainingExperimentsRequestOrder) UnmarshalJSON(b []byte) error {
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
