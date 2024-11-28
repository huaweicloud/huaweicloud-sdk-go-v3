package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecretTaskRequest Request Object
type ListSecretTaskRequest struct {

	// 凭据的名称。
	SecretName *string `json:"secret_name,omitempty"`

	// 任务状态。取值 ：  - SUCCESS ：任务轮转成功。 - FAILED ：任务轮转失败。
	Status *ListSecretTaskRequestStatus `json:"status,omitempty"`

	// 任务ID。 该参数与其他参数不能同时存在。
	TaskId *string `json:"task_id,omitempty"`

	// 每页返回的个数。 默认值：50。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的任务ID，为空时为查询第一页。
	Marker *string `json:"marker,omitempty"`
}

func (o ListSecretTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretTaskRequest struct{}"
	}

	return strings.Join([]string{"ListSecretTaskRequest", string(data)}, " ")
}

type ListSecretTaskRequestStatus struct {
	value string
}

type ListSecretTaskRequestStatusEnum struct {
	SUCCESS ListSecretTaskRequestStatus
	FAILED  ListSecretTaskRequestStatus
}

func GetListSecretTaskRequestStatusEnum() ListSecretTaskRequestStatusEnum {
	return ListSecretTaskRequestStatusEnum{
		SUCCESS: ListSecretTaskRequestStatus{
			value: "SUCCESS",
		},
		FAILED: ListSecretTaskRequestStatus{
			value: "FAILED",
		},
	}
}

func (c ListSecretTaskRequestStatus) Value() string {
	return c.value
}

func (c ListSecretTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecretTaskRequestStatus) UnmarshalJSON(b []byte) error {
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
