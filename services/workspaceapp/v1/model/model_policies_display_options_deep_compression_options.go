package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesDisplayOptionsDeepCompressionOptions 深度压缩控制选项。
type PoliciesDisplayOptionsDeepCompressionOptions struct {

	// 深度压缩级别。取值为： 压缩级别0：Compression grade 0 压缩级别1：Compression grade 1 压缩级别2：Compression grade 2 压缩级别3：Compression grade 3 压缩级别4：Compression grade 4 压缩级别5：Compression grade 5 压缩级别6：Compression grade 6 压缩级别7：Compression grade 7 压缩级别8：Compression grade 8 压缩级别9：Compression grade 9
	DeepCompressionLevel *PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel `json:"deep_compression_level,omitempty"`
}

func (o PoliciesDisplayOptionsDeepCompressionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesDisplayOptionsDeepCompressionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesDisplayOptionsDeepCompressionOptions", string(data)}, " ")
}

type PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel struct {
	value string
}

type PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevelEnum struct {
	COMPRESSION_GRADE_0 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_1 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_2 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_3 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_4 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_5 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_6 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_7 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_8 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
	COMPRESSION_GRADE_9 PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel
}

func GetPoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevelEnum() PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevelEnum {
	return PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevelEnum{
		COMPRESSION_GRADE_0: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 0",
		},
		COMPRESSION_GRADE_1: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 1",
		},
		COMPRESSION_GRADE_2: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 2",
		},
		COMPRESSION_GRADE_3: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 3",
		},
		COMPRESSION_GRADE_4: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 4",
		},
		COMPRESSION_GRADE_5: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 5",
		},
		COMPRESSION_GRADE_6: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 6",
		},
		COMPRESSION_GRADE_7: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 7",
		},
		COMPRESSION_GRADE_8: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 8",
		},
		COMPRESSION_GRADE_9: PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel{
			value: "Compression grade 9",
		},
	}
}

func (c PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel) Value() string {
	return c.value
}

func (c PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesDisplayOptionsDeepCompressionOptionsDeepCompressionLevel) UnmarshalJSON(b []byte) error {
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
