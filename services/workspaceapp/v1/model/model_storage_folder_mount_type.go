package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StorageFolderMountType 用户访问APS的应用时，对应NAS存储目录在APS上的挂载策略。 * `USER` - 仅挂载个人目录。 * `SHARE` - 仅挂载共享目录。 * `ANY` - 挂载目录不做限制(个人和共享NAS存储目录都会自动挂载)。
type StorageFolderMountType struct {
	value string
}

type StorageFolderMountTypeEnum struct {
	USER  StorageFolderMountType
	SHARE StorageFolderMountType
	ANY   StorageFolderMountType
}

func GetStorageFolderMountTypeEnum() StorageFolderMountTypeEnum {
	return StorageFolderMountTypeEnum{
		USER: StorageFolderMountType{
			value: "USER",
		},
		SHARE: StorageFolderMountType{
			value: "SHARE",
		},
		ANY: StorageFolderMountType{
			value: "ANY",
		},
	}
}

func (c StorageFolderMountType) Value() string {
	return c.value
}

func (c StorageFolderMountType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StorageFolderMountType) UnmarshalJSON(b []byte) error {
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
