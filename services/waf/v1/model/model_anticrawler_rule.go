package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AnticrawlerRule struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 匹配条件列表
	Conditions *[]AnticrawlerCondition `json:"conditions,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// **参数解释：** JS脚本反爬虫规则防护模式 **约束限制：** 不涉及 **取值范围：**  - anticrawler_except_url: 防护所有路径模式，在该模式下，查询的JS脚本反爬虫规则为排除的防护路径规则  - anticrawler_specific_url: 防护指定路径模式，在该模式下，查询的JS脚本反爬虫规则为指定要防护的路径规则  **默认取值：** anticrawler_except_url
	Type *AnticrawlerRuleType `json:"type,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// 执行该规则的优先级，值越小，优先级越高，值相同时，规则创建时间早，优先级越高。取值范围：0到65535。
	Priority *int32 `json:"priority,omitempty"`
}

func (o AnticrawlerRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnticrawlerRule struct{}"
	}

	return strings.Join([]string{"AnticrawlerRule", string(data)}, " ")
}

type AnticrawlerRuleType struct {
	value string
}

type AnticrawlerRuleTypeEnum struct {
	ANTICRAWLER_EXCEPT_URL   AnticrawlerRuleType
	ANTICRAWLER_SPECIFIC_URL AnticrawlerRuleType
}

func GetAnticrawlerRuleTypeEnum() AnticrawlerRuleTypeEnum {
	return AnticrawlerRuleTypeEnum{
		ANTICRAWLER_EXCEPT_URL: AnticrawlerRuleType{
			value: "anticrawler_except_url",
		},
		ANTICRAWLER_SPECIFIC_URL: AnticrawlerRuleType{
			value: "anticrawler_specific_url",
		},
	}
}

func (c AnticrawlerRuleType) Value() string {
	return c.value
}

func (c AnticrawlerRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnticrawlerRuleType) UnmarshalJSON(b []byte) error {
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
