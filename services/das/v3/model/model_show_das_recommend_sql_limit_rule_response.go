package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDasRecommendSqlLimitRuleResponse Response Object
type ShowDasRecommendSqlLimitRuleResponse struct {

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// sql的限流信息
	SqlLimitInfos  *[]RecommendSqlLimitRuleRespSqlLimitInfos `json:"sql_limit_infos,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ShowDasRecommendSqlLimitRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDasRecommendSqlLimitRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowDasRecommendSqlLimitRuleResponse", string(data)}, " ")
}
