package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Index struct {

	// 索引映射信息
	Mapping map[string]KeyIndex `json:"mapping"`

	// 数据管道ID
	PipeId string `json:"pipe_id"`

	// 索引状态；open 开启，closed 关闭
	Status IndexStatus `json:"status"`

	// 时间戳字段名称
	TimestampField string `json:"timestamp_field"`
}

func (o Index) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Index struct{}"
	}

	return strings.Join([]string{"Index", string(data)}, " ")
}

type IndexStatus struct {
	value string
}

type IndexStatusEnum struct {
	OPEN   IndexStatus
	CLOSED IndexStatus
}

func GetIndexStatusEnum() IndexStatusEnum {
	return IndexStatusEnum{
		OPEN: IndexStatus{
			value: "open",
		},
		CLOSED: IndexStatus{
			value: "closed",
		},
	}
}

func (c IndexStatus) Value() string {
	return c.value
}

func (c IndexStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IndexStatus) UnmarshalJSON(b []byte) error {
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
