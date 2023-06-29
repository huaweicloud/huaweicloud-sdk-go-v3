package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DatastoresResult 数据库引擎列表。
type DatastoresResult struct {

	// 部署形态支持的引擎版本列表
	SupportedVersions []string `json:"supported_versions"`

	// 部署形态
	InstanceMode DatastoresResultInstanceMode `json:"instance_mode"`
}

func (o DatastoresResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoresResult struct{}"
	}

	return strings.Join([]string{"DatastoresResult", string(data)}, " ")
}

type DatastoresResultInstanceMode struct {
	value string
}

type DatastoresResultInstanceModeEnum struct {
	HA          DatastoresResultInstanceMode
	INDEPENDENT DatastoresResultInstanceMode
}

func GetDatastoresResultInstanceModeEnum() DatastoresResultInstanceModeEnum {
	return DatastoresResultInstanceModeEnum{
		HA: DatastoresResultInstanceMode{
			value: "ha",
		},
		INDEPENDENT: DatastoresResultInstanceMode{
			value: "independent",
		},
	}
}

func (c DatastoresResultInstanceMode) Value() string {
	return c.value
}

func (c DatastoresResultInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatastoresResultInstanceMode) UnmarshalJSON(b []byte) error {
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
