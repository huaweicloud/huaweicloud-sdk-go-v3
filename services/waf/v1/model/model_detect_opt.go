package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DetectOpt 检测选项配置
type DetectOpt struct {

	// 检测类型，可选：lexicon(词库检测), compliance(合规检测), inject(注入检测), immunity_lexicon(豁免词库),supported_language(语言检测，默认开), pii(敏感词检测), abuse(算力滥用检测)
	Type *DetectOptType `json:"type,omitempty"`

	// 是否启用该检测类型
	Enable *bool `json:"enable,omitempty"`
}

func (o DetectOpt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectOpt struct{}"
	}

	return strings.Join([]string{"DetectOpt", string(data)}, " ")
}

type DetectOptType struct {
	value string
}

type DetectOptTypeEnum struct {
	LEXICON            DetectOptType
	COMPLIANCE         DetectOptType
	INJECT             DetectOptType
	IMMUNITY_LEXICON   DetectOptType
	SUPPORTED_LANGUAGE DetectOptType
	PII                DetectOptType
	ABUSE              DetectOptType
	TOPIC_DETECT       DetectOptType
}

func GetDetectOptTypeEnum() DetectOptTypeEnum {
	return DetectOptTypeEnum{
		LEXICON: DetectOptType{
			value: "lexicon",
		},
		COMPLIANCE: DetectOptType{
			value: "compliance",
		},
		INJECT: DetectOptType{
			value: "inject",
		},
		IMMUNITY_LEXICON: DetectOptType{
			value: "immunity_lexicon",
		},
		SUPPORTED_LANGUAGE: DetectOptType{
			value: "supported_language",
		},
		PII: DetectOptType{
			value: "pii",
		},
		ABUSE: DetectOptType{
			value: "abuse",
		},
		TOPIC_DETECT: DetectOptType{
			value: "topic_detect",
		},
	}
}

func (c DetectOptType) Value() string {
	return c.value
}

func (c DetectOptType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DetectOptType) UnmarshalJSON(b []byte) error {
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
