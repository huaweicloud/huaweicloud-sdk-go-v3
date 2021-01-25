package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type JobInfo struct {
	Name        *string         `json:"name,omitempty"`
	Nodes       *[]Node         `json:"nodes,omitempty"`
	Schedule    *Schedule       `json:"schedule,omitempty"`
	Params      *[]JobParam     `json:"params,omitempty"`
	Directory   *string         `json:"directory,omitempty"`
	JobType     *JobInfoJobType `json:"jobType,omitempty"`
	BasicConfig *BasicInfo      `json:"basicConfig,omitempty"`
}

func (o JobInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}

type JobInfoJobType struct {
	value string
}

type JobInfoJobTypeEnum struct {
	BATCH     JobInfoJobType
	REAL_TIME JobInfoJobType
}

func GetJobInfoJobTypeEnum() JobInfoJobTypeEnum {
	return JobInfoJobTypeEnum{
		BATCH: JobInfoJobType{
			value: "BATCH",
		},
		REAL_TIME: JobInfoJobType{
			value: "REAL_TIME",
		},
	}
}

func (c JobInfoJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *JobInfoJobType) UnmarshalJSON(b []byte) error {
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
