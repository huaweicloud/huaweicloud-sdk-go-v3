package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckJobResp 响应体
type CheckJobResp struct {

	// 任务id。
	Id string `json:"id"`

	// 测试结果。 - true：成功 - false：失败
	Status CheckJobRespStatus `json:"status"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`
}

func (o CheckJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobResp struct{}"
	}

	return strings.Join([]string{"CheckJobResp", string(data)}, " ")
}

type CheckJobRespStatus struct {
	value string
}

type CheckJobRespStatusEnum struct {
	TRUE  CheckJobRespStatus
	FALSE CheckJobRespStatus
}

func GetCheckJobRespStatusEnum() CheckJobRespStatusEnum {
	return CheckJobRespStatusEnum{
		TRUE: CheckJobRespStatus{
			value: "true",
		},
		FALSE: CheckJobRespStatus{
			value: "false",
		},
	}
}

func (c CheckJobRespStatus) Value() string {
	return c.value
}

func (c CheckJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckJobRespStatus) UnmarshalJSON(b []byte) error {
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
