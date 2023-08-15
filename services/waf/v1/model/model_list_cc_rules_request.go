package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCcRulesRequest Request Object
type ListCcRulesRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 策略id（策略id从查询防护策略列表接口获取）
	PolicyId string `json:"policy_id"`

	// 偏移量，表示查询该偏移量之后的记录。
	Offset int32 `json:"offset"`

	// 查询返回记录的数量限制。
	Limit int32 `json:"limit"`
}

func (o ListCcRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCcRulesRequest struct{}"
	}

	return strings.Join([]string{"ListCcRulesRequest", string(data)}, " ")
}
