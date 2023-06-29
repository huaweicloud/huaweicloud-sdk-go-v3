package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobInfoResponse Response Object
type ShowJobInfoResponse struct {

	// Job任务ID
	JobId *string `json:"job_id,omitempty"`

	// Job类型，取值范围：   masterStandbySwapJob：主备倒换任务   modifyClientIpTransJob：修改源ip透传
	JobType *ShowJobInfoResponseJobType `json:"job_type,omitempty"`

	// Job状态
	Status *ShowJobInfoResponseStatus `json:"status,omitempty"`

	// Job开始时间戳，单位为ms，格式为:1684191545379
	BeginTime *string `json:"begin_time,omitempty"`

	// Job开始时间戳，单位为ms，格式为:1684191548248
	EndTime *string `json:"end_time,omitempty"`

	// 失败原因
	FailReason     *string `json:"fail_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowJobInfoResponse", string(data)}, " ")
}

type ShowJobInfoResponseJobType struct {
	value string
}

type ShowJobInfoResponseJobTypeEnum struct {
	MASTER_STANDBY_SWAP_JOB    ShowJobInfoResponseJobType
	MODIFY_CLIENT_IP_TRANS_JOB ShowJobInfoResponseJobType
}

func GetShowJobInfoResponseJobTypeEnum() ShowJobInfoResponseJobTypeEnum {
	return ShowJobInfoResponseJobTypeEnum{
		MASTER_STANDBY_SWAP_JOB: ShowJobInfoResponseJobType{
			value: "masterStandbySwapJob",
		},
		MODIFY_CLIENT_IP_TRANS_JOB: ShowJobInfoResponseJobType{
			value: "modifyClientIpTransJob",
		},
	}
}

func (c ShowJobInfoResponseJobType) Value() string {
	return c.value
}

func (c ShowJobInfoResponseJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobInfoResponseJobType) UnmarshalJSON(b []byte) error {
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

type ShowJobInfoResponseStatus struct {
	value string
}

type ShowJobInfoResponseStatusEnum struct {
	INIT    ShowJobInfoResponseStatus
	RUNNING ShowJobInfoResponseStatus
	SUCCESS ShowJobInfoResponseStatus
	FAIL    ShowJobInfoResponseStatus
}

func GetShowJobInfoResponseStatusEnum() ShowJobInfoResponseStatusEnum {
	return ShowJobInfoResponseStatusEnum{
		INIT: ShowJobInfoResponseStatus{
			value: "INIT",
		},
		RUNNING: ShowJobInfoResponseStatus{
			value: "RUNNING",
		},
		SUCCESS: ShowJobInfoResponseStatus{
			value: "SUCCESS",
		},
		FAIL: ShowJobInfoResponseStatus{
			value: "FAIL",
		},
	}
}

func (c ShowJobInfoResponseStatus) Value() string {
	return c.value
}

func (c ShowJobInfoResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobInfoResponseStatus) UnmarshalJSON(b []byte) error {
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
