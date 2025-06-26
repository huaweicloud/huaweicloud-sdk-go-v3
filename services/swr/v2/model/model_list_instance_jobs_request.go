package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceJobsRequest Request Object
type ListInstanceJobsRequest struct {

	// 任务状态, 支持的状态为：Creating,Initializing,Running,Failed,Success
	Status *ListInstanceJobsRequestStatus `json:"status,omitempty"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为100，最大值为100。**注意：offset和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceJobsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceJobsRequest", string(data)}, " ")
}

type ListInstanceJobsRequestStatus struct {
	value string
}

type ListInstanceJobsRequestStatusEnum struct {
	CREATING     ListInstanceJobsRequestStatus
	INITIALIZING ListInstanceJobsRequestStatus
	RUNNING      ListInstanceJobsRequestStatus
	FAILED       ListInstanceJobsRequestStatus
	SUCCESS      ListInstanceJobsRequestStatus
}

func GetListInstanceJobsRequestStatusEnum() ListInstanceJobsRequestStatusEnum {
	return ListInstanceJobsRequestStatusEnum{
		CREATING: ListInstanceJobsRequestStatus{
			value: "Creating",
		},
		INITIALIZING: ListInstanceJobsRequestStatus{
			value: "Initializing",
		},
		RUNNING: ListInstanceJobsRequestStatus{
			value: "Running",
		},
		FAILED: ListInstanceJobsRequestStatus{
			value: "Failed",
		},
		SUCCESS: ListInstanceJobsRequestStatus{
			value: "Success",
		},
	}
}

func (c ListInstanceJobsRequestStatus) Value() string {
	return c.value
}

func (c ListInstanceJobsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceJobsRequestStatus) UnmarshalJSON(b []byte) error {
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
