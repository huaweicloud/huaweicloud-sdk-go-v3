package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ImportMqsInstanceTopicRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 导入topic的模式。 - AddNew：全量新增导入。 - Merge：合并导入。  默认为AddNew模式。

	Mode *ImportMqsInstanceTopicRequestMode `json:"mode,omitempty"`
	// App应用的前缀。  若加上前缀，导入Topic时会拼接前缀和已有的App应用，形成新的App应用名称，再根据新的App应用名称导入Topic。

	Prefix *string `json:"prefix,omitempty"`

	Body *ImportMqsInstanceTopicRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportMqsInstanceTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportMqsInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"ImportMqsInstanceTopicRequest", string(data)}, " ")
}

type ImportMqsInstanceTopicRequestMode struct {
	value string
}

type ImportMqsInstanceTopicRequestModeEnum struct {
	ADD_NEW ImportMqsInstanceTopicRequestMode
	MERGE   ImportMqsInstanceTopicRequestMode
}

func GetImportMqsInstanceTopicRequestModeEnum() ImportMqsInstanceTopicRequestModeEnum {
	return ImportMqsInstanceTopicRequestModeEnum{
		ADD_NEW: ImportMqsInstanceTopicRequestMode{
			value: "AddNew",
		},
		MERGE: ImportMqsInstanceTopicRequestMode{
			value: "Merge",
		},
	}
}

func (c ImportMqsInstanceTopicRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportMqsInstanceTopicRequestMode) UnmarshalJSON(b []byte) error {
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
