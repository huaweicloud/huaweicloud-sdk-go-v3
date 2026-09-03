package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UploadStatusEnum 上传状态枚举。
type UploadStatusEnum struct {
	value string
}

type UploadStatusEnumEnum struct {
	UPLOADING UploadStatusEnum
	UPLOADED  UploadStatusEnum
	FAILED    UploadStatusEnum
}

func GetUploadStatusEnumEnum() UploadStatusEnumEnum {
	return UploadStatusEnumEnum{
		UPLOADING: UploadStatusEnum{
			value: "UPLOADING",
		},
		UPLOADED: UploadStatusEnum{
			value: "UPLOADED",
		},
		FAILED: UploadStatusEnum{
			value: "FAILED",
		},
	}
}

func (c UploadStatusEnum) Value() string {
	return c.value
}

func (c UploadStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadStatusEnum) UnmarshalJSON(b []byte) error {
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
