package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateIndexRequestBody 更新索引请求体
type UpdateIndexRequestBody struct {

	// 索引映射信息
	Mapping map[string]KeyIndex `json:"mapping"`

	// 数据管道ID
	PipeId string `json:"pipe_id"`

	// 索引状态；open 开启，closed 关闭
	Status UpdateIndexRequestBodyStatus `json:"status"`

	// 时间戳字段名称
	TimestampField string `json:"timestamp_field"`
}

func (o UpdateIndexRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIndexRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIndexRequestBody", string(data)}, " ")
}

type UpdateIndexRequestBodyStatus struct {
	value string
}

type UpdateIndexRequestBodyStatusEnum struct {
	OPEN   UpdateIndexRequestBodyStatus
	CLOSED UpdateIndexRequestBodyStatus
}

func GetUpdateIndexRequestBodyStatusEnum() UpdateIndexRequestBodyStatusEnum {
	return UpdateIndexRequestBodyStatusEnum{
		OPEN: UpdateIndexRequestBodyStatus{
			value: "open",
		},
		CLOSED: UpdateIndexRequestBodyStatus{
			value: "closed",
		},
	}
}

func (c UpdateIndexRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateIndexRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIndexRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
