package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// L7转发规则
type L7Rule struct {
	// 转发规则的管理状，默认为true。  不支持该字段，请勿使用。

	AdminStateUp bool `json:"admin_state_up"`
	// 转发规则的匹配方式。type为HOST_NAME时可以为EQUAL_TO。type为PATH时可以为Perl类型的REGEX， STARTS_WITH，EQUAL_TO。

	CompareType string `json:"compare_type"`
	// 匹配内容的键值。目前匹配内容为HOST_NAME和PATH时，该字段不生效。该字段能更新但不会生效。

	Key string `json:"key"`
	// 转发规则所在的项目ID。

	ProjectId string `json:"project_id"`
	// 一个l7policy下创建的l7rule的HOST_NAME，PATH，METHOD，SOURCE_IP 不能重复。HEADER、QUERY_STRING支持重复的rule配置。  匹配内容：可以为HOST_NAME,PATH,METHOD,HEADER,QUERY_STRING,SOURCE_IP

	Type L7RuleType `json:"type"`
	// 匹配内容的值。仅当conditions空时该字段生效。 当type为HOST_NAME时，字符串只能包含英文字母、数字、“-”、“.”或“*”，必须以字母、数字或“*”开头。 若域名中包含“*”，则“*”只能出现在开头且必须以*.开始。当*开头时表示通配0~任一个字符。 当type为PATH时，当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!&#58;|/()[]{}，且必须以\"/\"开头。 当type为METHOD、SOURCE_IP、HEADER, QUERY_STRING时，该字段无意义，使用condition_pair来指定key，value。

	Value string `json:"value"`
	// provisioning状态，可以为ACTIVE、PENDING_CREATE 或者ERROR。 说明：该字段无实际含义，默认为ACTIVE。

	ProvisioningStatus string `json:"provisioning_status"`
	// 是否反向匹配。 使用说明： - 固定为false。该字段能更新但不会生效。

	Invert bool `json:"invert"`
	// 规则ID。

	Id string `json:"id"`
	// 转发规则的匹配条件。当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效。 配置了conditions后，字段key、字段value的值无意义。 若指定了conditons，该rule的条件数为conditons列表长度。 列表中key必须相同，value不允许重复。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	Conditions []RuleCondition `json:"conditions"`
}

func (o L7Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7Rule struct{}"
	}

	return strings.Join([]string{"L7Rule", string(data)}, " ")
}

type L7RuleType struct {
	value string
}

type L7RuleTypeEnum struct {
	HOST_NAME    L7RuleType
	PATH         L7RuleType
	METHOD       L7RuleType
	HEADER       L7RuleType
	QUERY_STRING L7RuleType
	SOURCE_IP    L7RuleType
}

func GetL7RuleTypeEnum() L7RuleTypeEnum {
	return L7RuleTypeEnum{
		HOST_NAME: L7RuleType{
			value: "HOST_NAME",
		},
		PATH: L7RuleType{
			value: "PATH",
		},
		METHOD: L7RuleType{
			value: "METHOD",
		},
		HEADER: L7RuleType{
			value: "HEADER",
		},
		QUERY_STRING: L7RuleType{
			value: "QUERY_STRING",
		},
		SOURCE_IP: L7RuleType{
			value: "SOURCE_IP",
		},
	}
}

func (c L7RuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7RuleType) UnmarshalJSON(b []byte) error {
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
