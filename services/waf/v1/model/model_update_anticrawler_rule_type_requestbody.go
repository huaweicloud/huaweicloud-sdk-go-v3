package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAnticrawlerRuleTypeRequestbody struct {

	// **参数解释：** JS脚本反爬虫规则类型 **约束限制：** 不涉及 **取值范围：**  - anticrawler_except_url: 防护所有路径模式，在该模式下，查询的JS脚本反爬虫规则为排除的防护路径规则  - anticrawler_specific_url: 防护指定路径模式，在该模式下，查询的JS脚本反爬虫规则为指定要防护的路径规则  **默认取值：** anticrawler_except_url
	AnticrawlerType UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerType `json:"anticrawler_type"`
}

func (o UpdateAnticrawlerRuleTypeRequestbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnticrawlerRuleTypeRequestbody struct{}"
	}

	return strings.Join([]string{"UpdateAnticrawlerRuleTypeRequestbody", string(data)}, " ")
}

type UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerType struct {
	value string
}

type UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerTypeEnum struct {
	ANTICRAWLER_EXCEPT_URL   UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerType
	ANTICRAWLER_SPECIFIC_URL UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerType
}

func GetUpdateAnticrawlerRuleTypeRequestbodyAnticrawlerTypeEnum() UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerTypeEnum {
	return UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerTypeEnum{
		ANTICRAWLER_EXCEPT_URL: UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerType{
			value: "anticrawler_except_url",
		},
		ANTICRAWLER_SPECIFIC_URL: UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerType{
			value: "anticrawler_specific_url",
		},
	}
}

func (c UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerType) Value() string {
	return c.value
}

func (c UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAnticrawlerRuleTypeRequestbodyAnticrawlerType) UnmarshalJSON(b []byte) error {
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
