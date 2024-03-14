package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDataProgressResponse Response Object
type UpdateDataProgressResponse struct {

	// 查询结果id
	Id *string `json:"id,omitempty"`

	// 查询状态
	Status *UpdateDataProgressResponseStatus `json:"status,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDataProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataProgressResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataProgressResponse", string(data)}, " ")
}

type UpdateDataProgressResponseStatus struct {
	value string
}

type UpdateDataProgressResponseStatusEnum struct {
	PENDING UpdateDataProgressResponseStatus
	FAILED  UpdateDataProgressResponseStatus
	SUCCESS UpdateDataProgressResponseStatus
}

func GetUpdateDataProgressResponseStatusEnum() UpdateDataProgressResponseStatusEnum {
	return UpdateDataProgressResponseStatusEnum{
		PENDING: UpdateDataProgressResponseStatus{
			value: "pending",
		},
		FAILED: UpdateDataProgressResponseStatus{
			value: "failed",
		},
		SUCCESS: UpdateDataProgressResponseStatus{
			value: "success",
		},
	}
}

func (c UpdateDataProgressResponseStatus) Value() string {
	return c.value
}

func (c UpdateDataProgressResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDataProgressResponseStatus) UnmarshalJSON(b []byte) error {
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
