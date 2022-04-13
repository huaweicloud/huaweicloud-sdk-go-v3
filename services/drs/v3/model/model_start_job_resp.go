package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 启动任务返回体。
type StartJobResp struct {
	// 任务id

	Id string `json:"id"`
	// 状态。

	Status *StartJobRespStatus `json:"status,omitempty"`
	// 错误码。

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息。

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o StartJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobResp struct{}"
	}

	return strings.Join([]string{"StartJobResp", string(data)}, " ")
}

type StartJobRespStatus struct {
	value string
}

type StartJobRespStatusEnum struct {
	SUCCESS StartJobRespStatus
	FAILED  StartJobRespStatus
}

func GetStartJobRespStatusEnum() StartJobRespStatusEnum {
	return StartJobRespStatusEnum{
		SUCCESS: StartJobRespStatus{
			value: "success",
		},
		FAILED: StartJobRespStatus{
			value: "failed",
		},
	}
}

func (c StartJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartJobRespStatus) UnmarshalJSON(b []byte) error {
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
