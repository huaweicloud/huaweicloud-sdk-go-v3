package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPositionResultResponse Response Object
type ShowPositionResultResponse struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 位点信息
	Position *string `json:"position,omitempty"`

	// 查询状态
	Status         *ShowPositionResultResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ShowPositionResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPositionResultResponse struct{}"
	}

	return strings.Join([]string{"ShowPositionResultResponse", string(data)}, " ")
}

type ShowPositionResultResponseStatus struct {
	value string
}

type ShowPositionResultResponseStatusEnum struct {
	PENDING ShowPositionResultResponseStatus
	FAILED  ShowPositionResultResponseStatus
	SUCCESS ShowPositionResultResponseStatus
}

func GetShowPositionResultResponseStatusEnum() ShowPositionResultResponseStatusEnum {
	return ShowPositionResultResponseStatusEnum{
		PENDING: ShowPositionResultResponseStatus{
			value: "pending",
		},
		FAILED: ShowPositionResultResponseStatus{
			value: "failed",
		},
		SUCCESS: ShowPositionResultResponseStatus{
			value: "success",
		},
	}
}

func (c ShowPositionResultResponseStatus) Value() string {
	return c.value
}

func (c ShowPositionResultResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPositionResultResponseStatus) UnmarshalJSON(b []byte) error {
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
