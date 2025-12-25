package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PipeCategory **参数解释**: 管道目录 - STREAMING_TO_INDEX 流式写入索引 - STREAMING_TO_LAKE 流式写入数据湖 - STREAMING_TO_INDEX_LAKE 流式写入索引和数据湖 - STREAMING  流式传输中  **约束限制** 不涉及  **取值范围**: - STREAMING_TO_INDEX - STREAMING_TO_LAKE - STREAMING_TO_INDEX_LAKE - STREAMING  **默认值** 不涉及
type PipeCategory struct {
	value string
}

type PipeCategoryEnum struct {
	STREAMING_TO_INDEX      PipeCategory
	STREAMING_TO_LAKE       PipeCategory
	STREAMING_TO_INDEX_LAKE PipeCategory
	STREAMING               PipeCategory
}

func GetPipeCategoryEnum() PipeCategoryEnum {
	return PipeCategoryEnum{
		STREAMING_TO_INDEX: PipeCategory{
			value: "STREAMING_TO_INDEX",
		},
		STREAMING_TO_LAKE: PipeCategory{
			value: "STREAMING_TO_LAKE",
		},
		STREAMING_TO_INDEX_LAKE: PipeCategory{
			value: "STREAMING_TO_INDEX_LAKE",
		},
		STREAMING: PipeCategory{
			value: "STREAMING",
		},
	}
}

func (c PipeCategory) Value() string {
	return c.value
}

func (c PipeCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipeCategory) UnmarshalJSON(b []byte) error {
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
