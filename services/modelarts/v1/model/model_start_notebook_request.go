package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartNotebookRequest Request Object
type StartNotebookRequest struct {

	// **参数解释**：启动后运行时长（单位:毫秒）。 **约束限制**：不涉及。 **取值范围**：3600000-259200000。 **默认取值**：3600000。
	Duration *int64 `json:"duration,omitempty"`

	// **参数解释**：Notebook实例ID。ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID），可通过调用[[查询Notebook实例列表接口](https://support.huaweicloud.com/api-modelarts/ListNotebooks.html#section0)](tag:hc)[[查询Notebook实例列表接口](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListNotebooks.html#section0)](tag:hk)获取。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`

	// **参数解释**：自动停止类别。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - TIMING：自动停止。 - IDLE：空闲停止。  **默认取值**：TIMING。
	Type *StartNotebookRequestType `json:"type,omitempty"`
}

func (o StartNotebookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartNotebookRequest struct{}"
	}

	return strings.Join([]string{"StartNotebookRequest", string(data)}, " ")
}

type StartNotebookRequestType struct {
	value string
}

type StartNotebookRequestTypeEnum struct {
	TIMING StartNotebookRequestType
	IDLE   StartNotebookRequestType
}

func GetStartNotebookRequestTypeEnum() StartNotebookRequestTypeEnum {
	return StartNotebookRequestTypeEnum{
		TIMING: StartNotebookRequestType{
			value: "timing",
		},
		IDLE: StartNotebookRequestType{
			value: "idle",
		},
	}
}

func (c StartNotebookRequestType) Value() string {
	return c.value
}

func (c StartNotebookRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartNotebookRequestType) UnmarshalJSON(b []byte) error {
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
