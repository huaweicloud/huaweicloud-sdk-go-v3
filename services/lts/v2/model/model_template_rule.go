package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 结构化模板规则对象
type TemplateRule struct {

	// 结构化类型，只支持custom_regex,json,split,nginx
	Type TemplateRuleType `json:"type" xml:"type"`

	// 具体结构化规则，每种结构化类型都有自己独有的结构，具体结构如下： 手动正则为json字符串，包含keyObject对象和regex_rules对象，keyObject内为键值对，键为demo_fields数组中元素的index，值为field_name，regex_rules对象为正则表达式字符串，整体例子为{\\\"keyObject\\\":{\\\"1\\\":\\\"date\\\",\\\"2\\\":\\\"num\\\"},\\\"regex_rules\\\":\\\"^(?<date>[^/]+)(?:[^ ]* ){8}(?<num>\\\\\\\\d+)\\\"}； json方式时param为一个json字符串，包含keyObject对象和layers对象，keyObject内为键值对，键为demo_fields数组中元素的field_name，值为user_defined_name，layers为最大解析层数，当前最大值为4，整体例子为{\\\"keyObject\\\":{\\\"metadata.dimention\\\":\\\"dimention\\\",\\\"metadata.value\\\":\\\"\\\",\\\"metadata.unit\\\":\\\"\\\",\\\"collectionTime\\\":\\\"\\\"},\\\"layers\\\":3}； 分隔符方式时为json字符串，包含keyObject对象和tokenizer对象，keyObject内为键值对，键为demo_fields数组中元素的index，值为field_name，tokenizer对象为所用分隔符，整体例子为{\\\"keyObject\\\":{\\\"0\\\":\\\"field1\\\",\\\"1\\\":\\\"field2\\\",\\\"2\\\":\\\"field3\\\",\\\"3\\\":\\\"field4\\\",\\\"4\\\":\\\"field5\\\",\\\"5\\\":\\\"field6\\\",\\\"6\\\":\\\"field7\\\",\\\"7\\\":\\\"field8\\\",\\\"8\\\":\\\"field9\\\"},\\\"tokenizer\\\":\\\" \\\"}； nginx方式时为json字符串，包含keyObject对象，regex对象，field_names对象及log_format对象，keyObject内为键值对，键为demo_fields数组中元素的field_name，值为user_defined_name，regex为正则表达式字符串，field_names对象为demo_fields数组中各元素的field_name的拼接字符串，每个field_name以','分隔，log_format对象为nginx日志格式化方式，具体方式参考https://support.huaweicloud.com/usermanual-lts/lts_0820.html#lts_0820__section1151119552549进行配置，整体例子为\"{\\\"keyObject\\\":{\\\"http_host\\\":\\\"host\\\",\\\"remote_addr\\\":\\\"\\\",\\\"request_method\\\":\\\"\\\",\\\"request_uri\\\":\\\"\\\",\\\"time_local\\\":\\\"\\\"},\\\"regex\\\":\\\"(\\\\\\\\d+/\\\\\\\\S+/\\\\\\\\d+:\\\\\\\\d+:\\\\\\\\d+:\\\\\\\\d+)\\\\\\\\s+\\\\\\\\S+\\\\\\\\s+(\\\\\\\\S*)\\\\\\\\s+(\\\\\\\\S*)\\\\\\\\s+(\\\\\\\\S*)\\\\\\\\s+\\\\\\\"([^\\\\\\\"]*)\\\\\\\".*\\\",\\\"fieldNames\\\":\\\"time_local,remote_addr,request_method,http_host,request_uri\\\",\\\"log_format\\\":\\\"log_format upstreaminfo '$time_local $remote_addr  $request_method $http_host \\\\\\\"$request_uri\\\\\\\"';\\\"}\"
	Param string `json:"param" xml:"param"`
}

func (o TemplateRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateRule struct{}"
	}

	return strings.Join([]string{"TemplateRule", string(data)}, " ")
}

type TemplateRuleType struct {
	value string
}

type TemplateRuleTypeEnum struct {
	CUSTOM_REGEX TemplateRuleType
	JSON         TemplateRuleType
	SPLIT        TemplateRuleType
	NGINX        TemplateRuleType
}

func GetTemplateRuleTypeEnum() TemplateRuleTypeEnum {
	return TemplateRuleTypeEnum{
		CUSTOM_REGEX: TemplateRuleType{
			value: "custom_regex",
		},
		JSON: TemplateRuleType{
			value: "json",
		},
		SPLIT: TemplateRuleType{
			value: "split",
		},
		NGINX: TemplateRuleType{
			value: "nginx",
		},
	}
}

func (c TemplateRuleType) Value() string {
	return c.value
}

func (c TemplateRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateRuleType) UnmarshalJSON(b []byte) error {
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
