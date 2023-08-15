package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ComponentStorage struct {
	Type ComponentStorageType `json:"type"`

	Parameters *StorageParameter `json:"parameters"`

	Mounts []ComponentMount `json:"mounts"`
}

func (o ComponentStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentStorage struct{}"
	}

	return strings.Join([]string{"ComponentStorage", string(data)}, " ")
}

type ComponentStorageType struct {
	value string
}

type ComponentStorageTypeEnum struct {
	HOST_PATH               ComponentStorageType
	EMPTY_DIR               ComponentStorageType
	CONFIG_MAP              ComponentStorageType
	SECRET                  ComponentStorageType
	PERSISTENT_VOLUME_CLAIM ComponentStorageType
}

func GetComponentStorageTypeEnum() ComponentStorageTypeEnum {
	return ComponentStorageTypeEnum{
		HOST_PATH: ComponentStorageType{
			value: "HostPath",
		},
		EMPTY_DIR: ComponentStorageType{
			value: "EmptyDir",
		},
		CONFIG_MAP: ComponentStorageType{
			value: "ConfigMap",
		},
		SECRET: ComponentStorageType{
			value: "Secret",
		},
		PERSISTENT_VOLUME_CLAIM: ComponentStorageType{
			value: "PersistentVolumeClaim",
		},
	}
}

func (c ComponentStorageType) Value() string {
	return c.value
}

func (c ComponentStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentStorageType) UnmarshalJSON(b []byte) error {
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
