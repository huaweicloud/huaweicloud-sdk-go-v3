package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateJobResponse struct {

	// 查询结果id
	Id *string `json:"id,omitempty"`

	// 查询状态
	Status         *UpdateJobResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o UpdateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobResponse", string(data)}, " ")
}

type UpdateJobResponseStatus struct {
	value string
}

type UpdateJobResponseStatusEnum struct {
	PENDING UpdateJobResponseStatus
	FAILED  UpdateJobResponseStatus
	SUCCESS UpdateJobResponseStatus
}

func GetUpdateJobResponseStatusEnum() UpdateJobResponseStatusEnum {
	return UpdateJobResponseStatusEnum{
		PENDING: UpdateJobResponseStatus{
			value: "pending",
		},
		FAILED: UpdateJobResponseStatus{
			value: "failed",
		},
		SUCCESS: UpdateJobResponseStatus{
			value: "success",
		},
	}
}

func (c UpdateJobResponseStatus) Value() string {
	return c.value
}

func (c UpdateJobResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateJobResponseStatus) UnmarshalJSON(b []byte) error {
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
