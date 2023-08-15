package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAccountAssignmentDeletionStatusRequest Request Object
type ListAccountAssignmentDeletionStatusRequest struct {
	InstanceId string `json:"instance_id"`

	// Filters he operation status list based on the passed attribute value.
	Status *ListAccountAssignmentDeletionStatusRequestStatus `json:"status,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`
}

func (o ListAccountAssignmentDeletionStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAssignmentDeletionStatusRequest struct{}"
	}

	return strings.Join([]string{"ListAccountAssignmentDeletionStatusRequest", string(data)}, " ")
}

type ListAccountAssignmentDeletionStatusRequestStatus struct {
	value string
}

type ListAccountAssignmentDeletionStatusRequestStatusEnum struct {
	IN_PROGRESS ListAccountAssignmentDeletionStatusRequestStatus
	SUCCEEDED   ListAccountAssignmentDeletionStatusRequestStatus
	FAILED      ListAccountAssignmentDeletionStatusRequestStatus
}

func GetListAccountAssignmentDeletionStatusRequestStatusEnum() ListAccountAssignmentDeletionStatusRequestStatusEnum {
	return ListAccountAssignmentDeletionStatusRequestStatusEnum{
		IN_PROGRESS: ListAccountAssignmentDeletionStatusRequestStatus{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: ListAccountAssignmentDeletionStatusRequestStatus{
			value: "SUCCEEDED",
		},
		FAILED: ListAccountAssignmentDeletionStatusRequestStatus{
			value: "FAILED",
		},
	}
}

func (c ListAccountAssignmentDeletionStatusRequestStatus) Value() string {
	return c.value
}

func (c ListAccountAssignmentDeletionStatusRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAccountAssignmentDeletionStatusRequestStatus) UnmarshalJSON(b []byte) error {
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
