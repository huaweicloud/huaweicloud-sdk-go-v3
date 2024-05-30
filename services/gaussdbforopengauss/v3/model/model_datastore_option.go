package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DatastoreOption struct {

	// 数据库版本。支持2.3版本，取值为“2.3”。
	EngineVersion string `json:"engine_version"`

	// 部署形态。
	InstanceMode DatastoreOptionInstanceMode `json:"instance_mode"`
}

func (o DatastoreOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreOption struct{}"
	}

	return strings.Join([]string{"DatastoreOption", string(data)}, " ")
}

type DatastoreOptionInstanceMode struct {
	value string
}

type DatastoreOptionInstanceModeEnum struct {
	HA          DatastoreOptionInstanceMode
	HAREADONLY  DatastoreOptionInstanceMode
	INDEPENDENT DatastoreOptionInstanceMode
}

func GetDatastoreOptionInstanceModeEnum() DatastoreOptionInstanceModeEnum {
	return DatastoreOptionInstanceModeEnum{
		HA: DatastoreOptionInstanceMode{
			value: "ha",
		},
		HAREADONLY: DatastoreOptionInstanceMode{
			value: "ha:readonly",
		},
		INDEPENDENT: DatastoreOptionInstanceMode{
			value: "independent",
		},
	}
}

func (c DatastoreOptionInstanceMode) Value() string {
	return c.value
}

func (c DatastoreOptionInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatastoreOptionInstanceMode) UnmarshalJSON(b []byte) error {
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
