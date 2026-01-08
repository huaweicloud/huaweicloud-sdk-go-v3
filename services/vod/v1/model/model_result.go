package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Result struct {

	// 任务类型
	Type *string `json:"type,omitempty"`

	// 媒资ID
	AssetId *string `json:"asset_id,omitempty"`

	// 转码状态 - WAITING 等待中 - PROCESSING 处理中 - SUCCEED 成功 - FAILED 失败
	Status *ResultStatus `json:"status,omitempty"`

	// 转码下发时间，格式按照RFC3339，UTC时间
	CreateTime *string `json:"create_time,omitempty"`

	// 转码结束时间
	EndTime *string `json:"end_time,omitempty"`

	TranscodeResult *TranscodeInfoResult `json:"transcode_result,omitempty"`
}

func (o Result) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Result struct{}"
	}

	return strings.Join([]string{"Result", string(data)}, " ")
}

type ResultStatus struct {
	value string
}

type ResultStatusEnum struct {
	WAITING    ResultStatus
	PROCESSING ResultStatus
	SUCCEED    ResultStatus
	FAILED     ResultStatus
}

func GetResultStatusEnum() ResultStatusEnum {
	return ResultStatusEnum{
		WAITING: ResultStatus{
			value: "WAITING",
		},
		PROCESSING: ResultStatus{
			value: "PROCESSING",
		},
		SUCCEED: ResultStatus{
			value: "SUCCEED",
		},
		FAILED: ResultStatus{
			value: "FAILED",
		},
	}
}

func (c ResultStatus) Value() string {
	return c.value
}

func (c ResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResultStatus) UnmarshalJSON(b []byte) error {
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
