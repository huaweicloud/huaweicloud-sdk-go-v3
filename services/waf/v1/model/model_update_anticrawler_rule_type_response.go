package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAnticrawlerRuleTypeResponse Response Object
type UpdateAnticrawlerRuleTypeResponse struct {

	// **参数解释：** JS脚本反爬虫规则类型 **约束限制：** 不涉及 **取值范围：**  - anticrawler_except_url: 防护所有路径模式，在该模式下，查询的JS脚本反爬虫规则为排除的防护路径规则  - anticrawler_specific_url: 防护指定路径模式，在该模式下，查询的JS脚本反爬虫规则为指定要防护的路径规则  **默认取值：** anticrawler_except_url
	AnticrawlerType *UpdateAnticrawlerRuleTypeResponseAnticrawlerType `json:"anticrawler_type,omitempty"`
	HttpStatusCode  int                                               `json:"-"`
}

func (o UpdateAnticrawlerRuleTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnticrawlerRuleTypeResponse struct{}"
	}

	return strings.Join([]string{"UpdateAnticrawlerRuleTypeResponse", string(data)}, " ")
}

type UpdateAnticrawlerRuleTypeResponseAnticrawlerType struct {
	value string
}

type UpdateAnticrawlerRuleTypeResponseAnticrawlerTypeEnum struct {
	ANTICRAWLER_EXCEPT_URL   UpdateAnticrawlerRuleTypeResponseAnticrawlerType
	ANTICRAWLER_SPECIFIC_URL UpdateAnticrawlerRuleTypeResponseAnticrawlerType
}

func GetUpdateAnticrawlerRuleTypeResponseAnticrawlerTypeEnum() UpdateAnticrawlerRuleTypeResponseAnticrawlerTypeEnum {
	return UpdateAnticrawlerRuleTypeResponseAnticrawlerTypeEnum{
		ANTICRAWLER_EXCEPT_URL: UpdateAnticrawlerRuleTypeResponseAnticrawlerType{
			value: "anticrawler_except_url",
		},
		ANTICRAWLER_SPECIFIC_URL: UpdateAnticrawlerRuleTypeResponseAnticrawlerType{
			value: "anticrawler_specific_url",
		},
	}
}

func (c UpdateAnticrawlerRuleTypeResponseAnticrawlerType) Value() string {
	return c.value
}

func (c UpdateAnticrawlerRuleTypeResponseAnticrawlerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAnticrawlerRuleTypeResponseAnticrawlerType) UnmarshalJSON(b []byte) error {
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
