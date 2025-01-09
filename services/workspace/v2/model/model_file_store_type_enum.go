package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FileStoreTypeEnum 存储方式： * `OBS` -  OBS桶存储。 * `LINK` - 外部存储,使用的可访问下载链接。
type FileStoreTypeEnum struct {
	value string
}

type FileStoreTypeEnumEnum struct {
	OBS  FileStoreTypeEnum
	LINK FileStoreTypeEnum
}

func GetFileStoreTypeEnumEnum() FileStoreTypeEnumEnum {
	return FileStoreTypeEnumEnum{
		OBS: FileStoreTypeEnum{
			value: "OBS",
		},
		LINK: FileStoreTypeEnum{
			value: "LINK",
		},
	}
}

func (c FileStoreTypeEnum) Value() string {
	return c.value
}

func (c FileStoreTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileStoreTypeEnum) UnmarshalJSON(b []byte) error {
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
