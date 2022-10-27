package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 请求数据。
type ExecuteActionRequestBodyMetadata struct {

	// 资源信息。
	Annotations map[string]string `json:"annotations,omitempty"`

	// action名称。
	Name ExecuteActionRequestBodyMetadataName `json:"name"`
}

func (o ExecuteActionRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteActionRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"ExecuteActionRequestBodyMetadata", string(data)}, " ")
}

type ExecuteActionRequestBodyMetadataName struct {
	value string
}

type ExecuteActionRequestBodyMetadataNameEnum struct {
	DEPLOY    ExecuteActionRequestBodyMetadataName
	UPGRADE   ExecuteActionRequestBodyMetadataName
	CONFIGURE ExecuteActionRequestBodyMetadataName
	SCALE     ExecuteActionRequestBodyMetadataName
	RESTART   ExecuteActionRequestBodyMetadataName
	START     ExecuteActionRequestBodyMetadataName
	STOP      ExecuteActionRequestBodyMetadataName
	ROLLBACK  ExecuteActionRequestBodyMetadataName
}

func GetExecuteActionRequestBodyMetadataNameEnum() ExecuteActionRequestBodyMetadataNameEnum {
	return ExecuteActionRequestBodyMetadataNameEnum{
		DEPLOY: ExecuteActionRequestBodyMetadataName{
			value: "deploy",
		},
		UPGRADE: ExecuteActionRequestBodyMetadataName{
			value: "upgrade",
		},
		CONFIGURE: ExecuteActionRequestBodyMetadataName{
			value: "configure",
		},
		SCALE: ExecuteActionRequestBodyMetadataName{
			value: "scale",
		},
		RESTART: ExecuteActionRequestBodyMetadataName{
			value: "restart",
		},
		START: ExecuteActionRequestBodyMetadataName{
			value: "start",
		},
		STOP: ExecuteActionRequestBodyMetadataName{
			value: "stop",
		},
		ROLLBACK: ExecuteActionRequestBodyMetadataName{
			value: "rollback",
		},
	}
}

func (c ExecuteActionRequestBodyMetadataName) Value() string {
	return c.value
}

func (c ExecuteActionRequestBodyMetadataName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteActionRequestBodyMetadataName) UnmarshalJSON(b []byte) error {
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
