package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceRequest Request Object
type ListInstanceRequest struct {

	// 起始索引，默认为0。**注意：offset和limit参数需要配套使用。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为100，最大值为100。**注意：offset和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`

	// 实例状态, Initial, Creating, Running, Unavailable
	Status *ListInstanceRequestStatus `json:"status,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceRequest", string(data)}, " ")
}

type ListInstanceRequestStatus struct {
	value string
}

type ListInstanceRequestStatusEnum struct {
	INITIAL     ListInstanceRequestStatus
	CREATING    ListInstanceRequestStatus
	RUNNING     ListInstanceRequestStatus
	UNAVAILABLE ListInstanceRequestStatus
}

func GetListInstanceRequestStatusEnum() ListInstanceRequestStatusEnum {
	return ListInstanceRequestStatusEnum{
		INITIAL: ListInstanceRequestStatus{
			value: "Initial",
		},
		CREATING: ListInstanceRequestStatus{
			value: "Creating",
		},
		RUNNING: ListInstanceRequestStatus{
			value: "Running",
		},
		UNAVAILABLE: ListInstanceRequestStatus{
			value: "Unavailable",
		},
	}
}

func (c ListInstanceRequestStatus) Value() string {
	return c.value
}

func (c ListInstanceRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceRequestStatus) UnmarshalJSON(b []byte) error {
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
