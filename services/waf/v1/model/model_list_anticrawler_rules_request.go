package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAnticrawlerRulesRequest Request Object
type ListAnticrawlerRulesRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，通过指定防护策略id来指明查询该防护策略下的防护规则，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 偏移量，表示查询该偏移量之后的记录。
	Offset int32 `json:"offset"`

	// 查询返回记录的数量限制。
	Limit int32 `json:"limit"`

	// **参数解释：** JS脚本反爬虫规则防护模式 **约束限制：** 不涉及 **取值范围：**  - anticrawler_except_url: 防护所有路径模式，在该模式下，查询的JS脚本反爬虫规则为排除的防护路径规则  - anticrawler_specific_url: 防护指定路径模式，在该模式下，查询的JS脚本反爬虫规则为指定要防护的路径规则  **默认取值：** anticrawler_except_url
	Type *ListAnticrawlerRulesRequestType `json:"type,omitempty"`
}

func (o ListAnticrawlerRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnticrawlerRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAnticrawlerRulesRequest", string(data)}, " ")
}

type ListAnticrawlerRulesRequestType struct {
	value string
}

type ListAnticrawlerRulesRequestTypeEnum struct {
	ANTICRAWLER_EXCEPT_URL   ListAnticrawlerRulesRequestType
	ANTICRAWLER_SPECIFIC_URL ListAnticrawlerRulesRequestType
}

func GetListAnticrawlerRulesRequestTypeEnum() ListAnticrawlerRulesRequestTypeEnum {
	return ListAnticrawlerRulesRequestTypeEnum{
		ANTICRAWLER_EXCEPT_URL: ListAnticrawlerRulesRequestType{
			value: "anticrawler_except_url",
		},
		ANTICRAWLER_SPECIFIC_URL: ListAnticrawlerRulesRequestType{
			value: "anticrawler_specific_url",
		},
	}
}

func (c ListAnticrawlerRulesRequestType) Value() string {
	return c.value
}

func (c ListAnticrawlerRulesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAnticrawlerRulesRequestType) UnmarshalJSON(b []byte) error {
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
