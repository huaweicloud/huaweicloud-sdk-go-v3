package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowJobResponse struct {
	Name *string `json:"name,omitempty" xml:"name"`

	Nodes *[]Node `json:"nodes,omitempty" xml:"nodes"`

	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule"`

	Params *[]JobParam `json:"params,omitempty" xml:"params"`

	Directory *string `json:"directory,omitempty" xml:"directory"`

	JobType *ShowJobResponseJobType `json:"jobType,omitempty" xml:"jobType"`

	BasicConfig    *BasicInfo `json:"basicConfig,omitempty" xml:"basicConfig"`
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
