package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteTrackerRequest Request Object
type DeleteTrackerRequest struct {

	// 标识追踪器名称。 在不传入该字段的情况下，将删除当前租户所有的数据类追踪器。
	TrackerName *string `json:"tracker_name,omitempty"`

	// 标识追踪器类型。 默认值为\"data\"。传入\"system\"时，配合tracker_name参数可删除管理类追踪器。
	TrackerType *DeleteTrackerRequestTrackerType `json:"tracker_type,omitempty"`
}

func (o DeleteTrackerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrackerRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrackerRequest", string(data)}, " ")
}

type DeleteTrackerRequestTrackerType struct {
	value string
}

type DeleteTrackerRequestTrackerTypeEnum struct {
	DATA   DeleteTrackerRequestTrackerType
	SYSTEM DeleteTrackerRequestTrackerType
}

func GetDeleteTrackerRequestTrackerTypeEnum() DeleteTrackerRequestTrackerTypeEnum {
	return DeleteTrackerRequestTrackerTypeEnum{
		DATA: DeleteTrackerRequestTrackerType{
			value: "data",
		},
		SYSTEM: DeleteTrackerRequestTrackerType{
			value: "system",
		},
	}
}

func (c DeleteTrackerRequestTrackerType) Value() string {
	return c.value
}

func (c DeleteTrackerRequestTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTrackerRequestTrackerType) UnmarshalJSON(b []byte) error {
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
