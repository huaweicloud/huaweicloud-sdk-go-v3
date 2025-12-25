package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPipeIndexResponse Response Object
type ShowPipeIndexResponse struct {

	// 索引映射信息
	Mapping map[string]KeyIndex `json:"mapping,omitempty"`

	// 数据管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 索引状态；open 开启，closed 关闭
	Status *ShowPipeIndexResponseStatus `json:"status,omitempty"`

	// 时间戳字段名称
	TimestampField *string `json:"timestamp_field,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPipeIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipeIndexResponse struct{}"
	}

	return strings.Join([]string{"ShowPipeIndexResponse", string(data)}, " ")
}

type ShowPipeIndexResponseStatus struct {
	value string
}

type ShowPipeIndexResponseStatusEnum struct {
	OPEN   ShowPipeIndexResponseStatus
	CLOSED ShowPipeIndexResponseStatus
}

func GetShowPipeIndexResponseStatusEnum() ShowPipeIndexResponseStatusEnum {
	return ShowPipeIndexResponseStatusEnum{
		OPEN: ShowPipeIndexResponseStatus{
			value: "open",
		},
		CLOSED: ShowPipeIndexResponseStatus{
			value: "closed",
		},
	}
}

func (c ShowPipeIndexResponseStatus) Value() string {
	return c.value
}

func (c ShowPipeIndexResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPipeIndexResponseStatus) UnmarshalJSON(b []byte) error {
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
