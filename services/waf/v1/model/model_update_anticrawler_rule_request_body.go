package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAnticrawlerRuleRequestBody struct {

	// 匹配条件列表
	Conditions []AnticrawlerCondition `json:"conditions"`

	// 规则名称
	Name string `json:"name"`

	// **参数解释：** JS脚本反爬虫规则类型 **约束限制：** 不涉及 **取值范围：**  - anticrawler_except_url: 防护所有路径模式，在该模式下，查询的JS脚本反爬虫规则为排除的防护路径规则  - anticrawler_specific_url: 防护指定路径模式，在该模式下，查询的JS脚本反爬虫规则为指定要防护的路径规则  **默认取值：** anticrawler_except_url
	Type UpdateAnticrawlerRuleRequestBodyType `json:"type"`

	// 执行该规则的优先级，值越小，优先级越高，值相同时，规则创建时间早，优先级越高。取值范围：65535。
	Priority int32 `json:"priority"`
}

func (o UpdateAnticrawlerRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnticrawlerRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAnticrawlerRuleRequestBody", string(data)}, " ")
}

type UpdateAnticrawlerRuleRequestBodyType struct {
	value string
}

type UpdateAnticrawlerRuleRequestBodyTypeEnum struct {
	ANTICRAWLER_EXCEPT_URL   UpdateAnticrawlerRuleRequestBodyType
	ANTICRAWLER_SPECIFIC_URL UpdateAnticrawlerRuleRequestBodyType
}

func GetUpdateAnticrawlerRuleRequestBodyTypeEnum() UpdateAnticrawlerRuleRequestBodyTypeEnum {
	return UpdateAnticrawlerRuleRequestBodyTypeEnum{
		ANTICRAWLER_EXCEPT_URL: UpdateAnticrawlerRuleRequestBodyType{
			value: "anticrawler_except_url",
		},
		ANTICRAWLER_SPECIFIC_URL: UpdateAnticrawlerRuleRequestBodyType{
			value: "anticrawler_specific_url",
		},
	}
}

func (c UpdateAnticrawlerRuleRequestBodyType) Value() string {
	return c.value
}

func (c UpdateAnticrawlerRuleRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAnticrawlerRuleRequestBodyType) UnmarshalJSON(b []byte) error {
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
