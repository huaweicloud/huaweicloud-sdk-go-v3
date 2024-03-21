package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobsRequest Request Object
type ShowJobsRequest struct {

	// job_id
	JobId string `json:"job_id"`

	Fields *[]ShowJobsRequestFields `json:"fields,omitempty"`
}

func (o ShowJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobsRequest struct{}"
	}

	return strings.Join([]string{"ShowJobsRequest", string(data)}, " ")
}

type ShowJobsRequestFields struct {
	value string
}

type ShowJobsRequestFieldsEnum struct {
	ID            ShowJobsRequestFields
	ACTION        ShowJobsRequestFields
	STATUS        ShowJobsRequestFields
	ERROR_TASK    ShowJobsRequestFields
	ERROR_CODE    ShowJobsRequestFields
	ERROR_MESSAGE ShowJobsRequestFields
	START_TIME    ShowJobsRequestFields
	END_TIME      ShowJobsRequestFields
}

func GetShowJobsRequestFieldsEnum() ShowJobsRequestFieldsEnum {
	return ShowJobsRequestFieldsEnum{
		ID: ShowJobsRequestFields{
			value: "id",
		},
		ACTION: ShowJobsRequestFields{
			value: "action",
		},
		STATUS: ShowJobsRequestFields{
			value: "status",
		},
		ERROR_TASK: ShowJobsRequestFields{
			value: "error_task",
		},
		ERROR_CODE: ShowJobsRequestFields{
			value: "error_code",
		},
		ERROR_MESSAGE: ShowJobsRequestFields{
			value: "error_message",
		},
		START_TIME: ShowJobsRequestFields{
			value: "start_time",
		},
		END_TIME: ShowJobsRequestFields{
			value: "end_time",
		},
	}
}

func (c ShowJobsRequestFields) Value() string {
	return c.value
}

func (c ShowJobsRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobsRequestFields) UnmarshalJSON(b []byte) error {
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
