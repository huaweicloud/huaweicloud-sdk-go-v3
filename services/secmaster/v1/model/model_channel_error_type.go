package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChannelErrorType **参数解释**: 采集通道失败类型 - SUCCESS 成功 - FILE_UPLOAD_ERROR 文件上传失败 - FILE_COPY_ERROR 文件复制失败 - FILE_ZIP_ERROR 文件压缩失败 - SALT_EXECUTE_ERROR Salt执行出错  **约束限制** 不涉及 **取值范围**: - SUCCESS - FILE_UPLOAD_ERROR - FILE_COPY_ERROR - FILE_ZIP_ERROR - SALT_EXECUTE_ERROR  **默认值** 不涉及
type ChannelErrorType struct {
	value string
}

type ChannelErrorTypeEnum struct {
	SUCCESS            ChannelErrorType
	FILE_UPLOAD_ERROR  ChannelErrorType
	FILE_COPY_ERROR    ChannelErrorType
	FILE_ZIP_ERROR     ChannelErrorType
	SALT_EXECUTE_ERROR ChannelErrorType
}

func GetChannelErrorTypeEnum() ChannelErrorTypeEnum {
	return ChannelErrorTypeEnum{
		SUCCESS: ChannelErrorType{
			value: "SUCCESS",
		},
		FILE_UPLOAD_ERROR: ChannelErrorType{
			value: "FILE_UPLOAD_ERROR",
		},
		FILE_COPY_ERROR: ChannelErrorType{
			value: "FILE_COPY_ERROR",
		},
		FILE_ZIP_ERROR: ChannelErrorType{
			value: "FILE_ZIP_ERROR",
		},
		SALT_EXECUTE_ERROR: ChannelErrorType{
			value: "SALT_EXECUTE_ERROR",
		},
	}
}

func (c ChannelErrorType) Value() string {
	return c.value
}

func (c ChannelErrorType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChannelErrorType) UnmarshalJSON(b []byte) error {
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
