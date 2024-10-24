package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAccountAssignmentDeletionStatusRequest Request Object
type ListAccountAssignmentDeletionStatusRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM身份中心实例的全局唯一标识符（ID）。
	InstanceId string `json:"instance_id"`

	// 根据传递的属性值过滤操作状态列表
	Status *ListAccountAssignmentDeletionStatusRequestStatus `json:"status,omitempty"`

	// 每个请求返回的最大结果数
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
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
