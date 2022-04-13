package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListInstanceSnapshotsRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 组件ID。

	ComponentId string `json:"component_id"`
	// 组件实例ID。

	InstanceId string `json:"instance_id"`
	// 指定个数，明确指定的时候用于分页，取值[0, 100]。不指定的时候表示不分页，最多查询1000条记录。

	Limit *int32 `json:"limit,omitempty"`
	// 指定查询偏移量，默认偏移量为0.

	Offset *int32 `json:"offset,omitempty"`
	// 排序字段，默认按创建时间排序。  排序字段支持枚举值：create_time、version。

	SnapshotOrderBy *string `json:"snapshot_order_by,omitempty"`
	// desc/asc，默认desc。

	Order *ListInstanceSnapshotsRequestOrder `json:"order,omitempty"`
}

func (o ListInstanceSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceSnapshotsRequest", string(data)}, " ")
}

type ListInstanceSnapshotsRequestOrder struct {
	value string
}

type ListInstanceSnapshotsRequestOrderEnum struct {
	DESC ListInstanceSnapshotsRequestOrder
	ASC  ListInstanceSnapshotsRequestOrder
}

func GetListInstanceSnapshotsRequestOrderEnum() ListInstanceSnapshotsRequestOrderEnum {
	return ListInstanceSnapshotsRequestOrderEnum{
		DESC: ListInstanceSnapshotsRequestOrder{
			value: "desc",
		},
		ASC: ListInstanceSnapshotsRequestOrder{
			value: "asc",
		},
	}
}

func (c ListInstanceSnapshotsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceSnapshotsRequestOrder) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
