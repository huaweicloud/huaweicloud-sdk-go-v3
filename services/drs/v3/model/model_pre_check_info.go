package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PreCheckInfo 预检查信息体
type PreCheckInfo struct {

	// 任务id
	JobId string `json:"job_id"`

	// 预检查模式
	PrecheckMode PreCheckInfoPrecheckMode `json:"precheck_mode"`
}

func (o PreCheckInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreCheckInfo struct{}"
	}

	return strings.Join([]string{"PreCheckInfo", string(data)}, " ")
}

type PreCheckInfoPrecheckMode struct {
	value string
}

type PreCheckInfoPrecheckModeEnum struct {
	FOR_START_JOB PreCheckInfoPrecheckMode
	FOR_RETRY_JOB PreCheckInfoPrecheckMode
}

func GetPreCheckInfoPrecheckModeEnum() PreCheckInfoPrecheckModeEnum {
	return PreCheckInfoPrecheckModeEnum{
		FOR_START_JOB: PreCheckInfoPrecheckMode{
			value: "forStartJob",
		},
		FOR_RETRY_JOB: PreCheckInfoPrecheckMode{
			value: "forRetryJob",
		},
	}
}

func (c PreCheckInfoPrecheckMode) Value() string {
	return c.value
}

func (c PreCheckInfoPrecheckMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PreCheckInfoPrecheckMode) UnmarshalJSON(b []byte) error {
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
