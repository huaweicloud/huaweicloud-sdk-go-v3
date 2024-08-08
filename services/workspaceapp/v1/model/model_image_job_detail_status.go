package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageJobDetailStatus job详情的状态 * `WAITING` - 等待 * `RUNNING` - 运行中 * `SUCCESS` - 成功 * `FAILED` - 失败 * `ABNORMAL` - 异常 * `ROLLBACK` - 回滚中 * `ABORTING` - 终止中
type ImageJobDetailStatus struct {
	value string
}

type ImageJobDetailStatusEnum struct {
	WAITING  ImageJobDetailStatus
	RUNNING  ImageJobDetailStatus
	SUCCESS  ImageJobDetailStatus
	FAILED   ImageJobDetailStatus
	ABNORMAL ImageJobDetailStatus
	ROLLBACK ImageJobDetailStatus
	ABORTING ImageJobDetailStatus
}

func GetImageJobDetailStatusEnum() ImageJobDetailStatusEnum {
	return ImageJobDetailStatusEnum{
		WAITING: ImageJobDetailStatus{
			value: "WAITING",
		},
		RUNNING: ImageJobDetailStatus{
			value: "RUNNING",
		},
		SUCCESS: ImageJobDetailStatus{
			value: "SUCCESS",
		},
		FAILED: ImageJobDetailStatus{
			value: "FAILED",
		},
		ABNORMAL: ImageJobDetailStatus{
			value: "ABNORMAL",
		},
		ROLLBACK: ImageJobDetailStatus{
			value: "ROLLBACK",
		},
		ABORTING: ImageJobDetailStatus{
			value: "ABORTING",
		},
	}
}

func (c ImageJobDetailStatus) Value() string {
	return c.value
}

func (c ImageJobDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageJobDetailStatus) UnmarshalJSON(b []byte) error {
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
