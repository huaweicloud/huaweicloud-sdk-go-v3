package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ComponentStorage struct {
	Type *ComponentStorageType `json:"type,omitempty"`

	// 存储盘的名字
	Name *string `json:"name,omitempty"`

	Parameters *ComponentStorage `json:"parameters,omitempty"`

	Mounts *[]ComponentMount `json:"mounts,omitempty"`
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
