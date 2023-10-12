package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DisplayOptionsDeepCompressionOptions 深度压缩控制选项。
type DisplayOptionsDeepCompressionOptions struct {

	// 深度压缩级别。取值为： 压缩级别0：Compression grade 0 压缩级别1：Compression grade 1 压缩级别2：Compression grade 2 压缩级别3：Compression grade 3 压缩级别4：Compression grade 4 压缩级别5：Compression grade 5 压缩级别6：Compression grade 6 压缩级别7：Compression grade 7 压缩级别8：Compression grade 8 压缩级别9：Compression grade 9
	DeepCompressionLevel *DisplayOptionsDeepCompressionOptionsDeepCompressionLevel `json:"deep_compression_level,omitempty"`
}

func (o DisplayOptionsDeepCompressionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisplayOptionsDeepCompressionOptions struct{}"
	}

	return strings.Join([]string{"DisplayOptionsDeepCompressionOptions", string(data)}, " ")
}

type DisplayOptionsDeepCompressionOptionsDeepCompressionLevel struct {
	value string
}

type DisplayOptionsDeepCompressionOptionsDeepCompressionLevelEnum struct {
	COMPRESSION_GRADE_0 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_1 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_2 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_3 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_4 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_5 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_6 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_7 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_8 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_9 DisplayOptionsDeepCompressionOptionsDeepCompressionLevel
}

func GetDisplayOptionsDeepCompressionOptionsDeepCompressionLevelEnum() DisplayOptionsDeepCompressionOptionsDeepCompressionLevelEnum {
	return DisplayOptionsDeepCompressionOptionsDeepCompressionLevelEnum{
		COMPRESSION_GRADE_0: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 0",
		},
		COMPRESSION_GRADE_1: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 1",
		},
		COMPRESSION_GRADE_2: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 2",
		},
		COMPRESSION_GRADE_3: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 3",
		},
		COMPRESSION_GRADE_4: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 4",
		},
		COMPRESSION_GRADE_5: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 5",
		},
		COMPRESSION_GRADE_6: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 6",
		},
		COMPRESSION_GRADE_7: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 7",
		},
		COMPRESSION_GRADE_8: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 8",
		},
		COMPRESSION_GRADE_9: DisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 9",
		},
	}
}

func (c DisplayOptionsDeepCompressionOptionsDeepCompressionLevel) Value() string {
	return c.value
}

func (c DisplayOptionsDeepCompressionOptionsDeepCompressionLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisplayOptionsDeepCompressionOptionsDeepCompressionLevel) UnmarshalJSON(b []byte) error {
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
