package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntileakageRulesRequest Request Object
type ListAntileakageRulesRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，通过指定防护策略id来指明查询该防护策略下的防敏感信息泄露规则，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 偏移量，表示查询该偏移量之后的记录。
	Offset int32 `json:"offset"`

	// 查询返回记录的数量限制。
	Limit int32 `json:"limit"`
}

func (o ListAntileakageRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntileakageRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAntileakageRulesRequest", string(data)}, " ")
}
