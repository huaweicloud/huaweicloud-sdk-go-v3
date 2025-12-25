package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePipeIndexResponse Response Object
type UpdatePipeIndexResponse struct {

	// 索引映射信息
	Mapping map[string]KeyIndex `json:"mapping,omitempty"`

	// 数据管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 索引状态；open 开启，closed 关闭
	Status *UpdatePipeIndexResponseStatus `json:"status,omitempty"`

	// 时间戳字段名称
	TimestampField *string `json:"timestamp_field,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePipeIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipeIndexResponse struct{}"
	}

	return strings.Join([]string{"UpdatePipeIndexResponse", string(data)}, " ")
}

type UpdatePipeIndexResponseStatus struct {
	value string
}

type UpdatePipeIndexResponseStatusEnum struct {
	OPEN   UpdatePipeIndexResponseStatus
	CLOSED UpdatePipeIndexResponseStatus
}

func GetUpdatePipeIndexResponseStatusEnum() UpdatePipeIndexResponseStatusEnum {
	return UpdatePipeIndexResponseStatusEnum{
		OPEN: UpdatePipeIndexResponseStatus{
			value: "open",
		},
		CLOSED: UpdatePipeIndexResponseStatus{
			value: "closed",
		},
	}
}

func (c UpdatePipeIndexResponseStatus) Value() string {
	return c.value
}

func (c UpdatePipeIndexResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePipeIndexResponseStatus) UnmarshalJSON(b []byte) error {
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
