package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAccountAssignmentCreationStatusRequest Request Object
type ListAccountAssignmentCreationStatusRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// 根据传递的属性值过滤操作状态列表
	Status *ListAccountAssignmentCreationStatusRequestStatus `json:"status,omitempty"`

	// 每个请求返回的最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`
}

func (o ListAccountAssignmentCreationStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAssignmentCreationStatusRequest struct{}"
	}

	return strings.Join([]string{"ListAccountAssignmentCreationStatusRequest", string(data)}, " ")
}

type ListAccountAssignmentCreationStatusRequestStatus struct {
	value string
}

type ListAccountAssignmentCreationStatusRequestStatusEnum struct {
	IN_PROGRESS ListAccountAssignmentCreationStatusRequestStatus
	SUCCEEDED   ListAccountAssignmentCreationStatusRequestStatus
	FAILED      ListAccountAssignmentCreationStatusRequestStatus
}

func GetListAccountAssignmentCreationStatusRequestStatusEnum() ListAccountAssignmentCreationStatusRequestStatusEnum {
	return ListAccountAssignmentCreationStatusRequestStatusEnum{
		IN_PROGRESS: ListAccountAssignmentCreationStatusRequestStatus{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: ListAccountAssignmentCreationStatusRequestStatus{
			value: "SUCCEEDED",
		},
		FAILED: ListAccountAssignmentCreationStatusRequestStatus{
			value: "FAILED",
		},
	}
}

func (c ListAccountAssignmentCreationStatusRequestStatus) Value() string {
	return c.value
}

func (c ListAccountAssignmentCreationStatusRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAccountAssignmentCreationStatusRequestStatus) UnmarshalJSON(b []byte) error {
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
