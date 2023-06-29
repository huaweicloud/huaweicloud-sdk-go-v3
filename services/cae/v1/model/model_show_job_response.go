package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobResponse Response Object
type ShowJobResponse struct {

	// API版本。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类。
	Kind *ShowJobResponseKind `json:"kind,omitempty"`

	Spec           *JobSpec `json:"spec,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}

type ShowJobResponseKind struct {
	value string
}

type ShowJobResponseKindEnum struct {
	JOB ShowJobResponseKind
}

func GetShowJobResponseKindEnum() ShowJobResponseKindEnum {
	return ShowJobResponseKindEnum{
		JOB: ShowJobResponseKind{
			value: "Job",
		},
	}
}

func (c ShowJobResponseKind) Value() string {
	return c.value
}

func (c ShowJobResponseKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobResponseKind) UnmarshalJSON(b []byte) error {
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
