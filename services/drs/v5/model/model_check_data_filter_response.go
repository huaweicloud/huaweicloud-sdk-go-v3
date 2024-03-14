package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckDataFilterResponse Response Object
type CheckDataFilterResponse struct {

	// 查询结果id
	Id *string `json:"id,omitempty"`

	// 查询状态
	Status *CheckDataFilterResponseStatus `json:"status,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckDataFilterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataFilterResponse struct{}"
	}

	return strings.Join([]string{"CheckDataFilterResponse", string(data)}, " ")
}

type CheckDataFilterResponseStatus struct {
	value string
}

type CheckDataFilterResponseStatusEnum struct {
	PENDING CheckDataFilterResponseStatus
	FAILED  CheckDataFilterResponseStatus
	SUCCESS CheckDataFilterResponseStatus
}

func GetCheckDataFilterResponseStatusEnum() CheckDataFilterResponseStatusEnum {
	return CheckDataFilterResponseStatusEnum{
		PENDING: CheckDataFilterResponseStatus{
			value: "pending",
		},
		FAILED: CheckDataFilterResponseStatus{
			value: "failed",
		},
		SUCCESS: CheckDataFilterResponseStatus{
			value: "success",
		},
	}
}

func (c CheckDataFilterResponseStatus) Value() string {
	return c.value
}

func (c CheckDataFilterResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckDataFilterResponseStatus) UnmarshalJSON(b []byte) error {
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
