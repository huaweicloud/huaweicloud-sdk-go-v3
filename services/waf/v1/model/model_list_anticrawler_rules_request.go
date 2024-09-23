package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAnticrawlerRulesRequest Request Object
type ListAnticrawlerRulesRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，通过指定防护策略id来指明查询该防护策略下的防护规则，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 偏移量，表示查询该偏移量之后的记录。
	Offset int32 `json:"offset"`

	// 查询返回记录的数量限制。
	Limit int32 `json:"limit"`

	// JS脚本反爬虫规则防护模式   - anticrawler_except_url: 防护所有路径模式，在该模式下，查询的JS脚本反爬虫规则为排除的防护路径规则   - anticrawler_specific_url: 防护指定路径模式，在该模式下，查询的JS脚本反爬虫规则为指定要防护的路径规则   - 默认值：anticrawler_except_url
	Type *string `json:"type,omitempty"`
}

func (o ListAnticrawlerRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnticrawlerRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAnticrawlerRulesRequest", string(data)}, " ")
}
