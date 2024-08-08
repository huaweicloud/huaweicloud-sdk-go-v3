package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageJobStatus job状态 * `WAITING` - 等待 * `RUNNING` - 运行中 * `SUCCESS - 完成 * `FAILED` - 失败
type ImageJobStatus struct {
	value string
}

type ImageJobStatusEnum struct {
	WAITING ImageJobStatus
	RUNNING ImageJobStatus
	SUCCESS ImageJobStatus
	FAILED  ImageJobStatus
}

func GetImageJobStatusEnum() ImageJobStatusEnum {
	return ImageJobStatusEnum{
		WAITING: ImageJobStatus{
			value: "WAITING",
		},
		RUNNING: ImageJobStatus{
			value: "RUNNING",
		},
		SUCCESS: ImageJobStatus{
			value: "SUCCESS",
		},
		FAILED: ImageJobStatus{
			value: "FAILED",
		},
	}
}

func (c ImageJobStatus) Value() string {
	return c.value
}

func (c ImageJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageJobStatus) UnmarshalJSON(b []byte) error {
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
