package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceOperationsRequest Request Object
type ListInstanceOperationsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 操作信息。
	Operation ListInstanceOperationsRequestOperation `json:"operation"`
}

func (o ListInstanceOperationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceOperationsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceOperationsRequest", string(data)}, " ")
}

type ListInstanceOperationsRequestOperation struct {
	value string
}

type ListInstanceOperationsRequestOperationEnum struct {
	EXTEND ListInstanceOperationsRequestOperation
}

func GetListInstanceOperationsRequestOperationEnum() ListInstanceOperationsRequestOperationEnum {
	return ListInstanceOperationsRequestOperationEnum{
		EXTEND: ListInstanceOperationsRequestOperation{
			value: "extend",
		},
	}
}

func (c ListInstanceOperationsRequestOperation) Value() string {
	return c.value
}

func (c ListInstanceOperationsRequestOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceOperationsRequestOperation) UnmarshalJSON(b []byte) error {
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
