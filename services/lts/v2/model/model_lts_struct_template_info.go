package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LtsStructTemplateInfo struct {
	// 结构化字段

	DemoFields []StructFieldInfo `json:"demo_fields"`
	// 示例日志

	Content string `json:"content"`
	// 日志组ID

	LogGroupId string `json:"log_group_id"`
	// 结构化方式

	ParseType LtsStructTemplateInfoParseType `json:"parse_type"`
	// 日志流ID

	LogStreamId string `json:"log_stream_id"`
	// 项目ID

	ProjectId string `json:"project_id"`
	// parse_type为custom_regex类型时必填，regex提取规则

	RegexRules *string `json:"regex_rules,omitempty"`
	// parse_type为json类型时必填，解析层数，目前固定是3

	Layers *int32 `json:"layers,omitempty"`
	// parse_type为split类型时必填，分隔符，分词符号

	Tokenizer *string `json:"tokenizer,omitempty"`
	// parse_type为nginx类型时必填，nginx日志格式模板

	LogFormat *string `json:"log_format,omitempty"`
}

func (o LtsStructTemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsStructTemplateInfo struct{}"
	}

	return strings.Join([]string{"LtsStructTemplateInfo", string(data)}, " ")
}

type LtsStructTemplateInfoParseType struct {
	value string
}

type LtsStructTemplateInfoParseTypeEnum struct {
	BUILT_IN     LtsStructTemplateInfoParseType
	JSON         LtsStructTemplateInfoParseType
	CUSTOM_REGEX LtsStructTemplateInfoParseType
	SPLIT        LtsStructTemplateInfoParseType
	NGINX        LtsStructTemplateInfoParseType
}

func GetLtsStructTemplateInfoParseTypeEnum() LtsStructTemplateInfoParseTypeEnum {
	return LtsStructTemplateInfoParseTypeEnum{
		BUILT_IN: LtsStructTemplateInfoParseType{
			value: "built_in",
		},
		JSON: LtsStructTemplateInfoParseType{
			value: "json",
		},
		CUSTOM_REGEX: LtsStructTemplateInfoParseType{
			value: "custom_regex",
		},
		SPLIT: LtsStructTemplateInfoParseType{
			value: "split",
		},
		NGINX: LtsStructTemplateInfoParseType{
			value: "nginx",
		},
	}
}

func (c LtsStructTemplateInfoParseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LtsStructTemplateInfoParseType) UnmarshalJSON(b []byte) error {
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
