package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDataProcessingRulesResultResponse Response Object
type ShowDataProcessingRulesResultResponse struct {

	// 查询结果id
	Id *string `json:"id,omitempty"`

	// 查询状态 pending：处理中 failed：失败 success：成功
	Status *ShowDataProcessingRulesResultResponseStatus `json:"status,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDataProcessingRulesResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataProcessingRulesResultResponse struct{}"
	}

	return strings.Join([]string{"ShowDataProcessingRulesResultResponse", string(data)}, " ")
}

type ShowDataProcessingRulesResultResponseStatus struct {
	value string
}

type ShowDataProcessingRulesResultResponseStatusEnum struct {
	PENDING ShowDataProcessingRulesResultResponseStatus
	FAILED  ShowDataProcessingRulesResultResponseStatus
	SUCCESS ShowDataProcessingRulesResultResponseStatus
}

func GetShowDataProcessingRulesResultResponseStatusEnum() ShowDataProcessingRulesResultResponseStatusEnum {
	return ShowDataProcessingRulesResultResponseStatusEnum{
		PENDING: ShowDataProcessingRulesResultResponseStatus{
			value: "pending",
		},
		FAILED: ShowDataProcessingRulesResultResponseStatus{
			value: "failed",
		},
		SUCCESS: ShowDataProcessingRulesResultResponseStatus{
			value: "success",
		},
	}
}

func (c ShowDataProcessingRulesResultResponseStatus) Value() string {
	return c.value
}

func (c ShowDataProcessingRulesResultResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataProcessingRulesResultResponseStatus) UnmarshalJSON(b []byte) error {
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
