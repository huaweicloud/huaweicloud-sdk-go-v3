package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobResponse Response Object
type ShowJobResponse struct {
	Name *string `json:"name,omitempty"`

	Nodes *[]Node `json:"nodes,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	Params *[]JobParam `json:"params,omitempty"`

	Directory *string `json:"directory,omitempty"`

	JobType *ShowJobResponseJobType `json:"jobType,omitempty"`

	BasicConfig    *BasicInfo `json:"basicConfig,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}

type ShowJobResponseJobType struct {
	value string
}

type ShowJobResponseJobTypeEnum struct {
	BATCH     ShowJobResponseJobType
	REAL_TIME ShowJobResponseJobType
}

func GetShowJobResponseJobTypeEnum() ShowJobResponseJobTypeEnum {
	return ShowJobResponseJobTypeEnum{
		BATCH: ShowJobResponseJobType{
			value: "BATCH",
		},
		REAL_TIME: ShowJobResponseJobType{
			value: "REAL_TIME",
		},
	}
}

func (c ShowJobResponseJobType) Value() string {
	return c.value
}

func (c ShowJobResponseJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobResponseJobType) UnmarshalJSON(b []byte) error {
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
