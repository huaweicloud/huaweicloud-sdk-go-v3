package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowRecommendSqlLimitRuleRequestBody struct {

	// 引擎类型，目前只支持mysql或者taurus
	EngineType *string `json:"engine_type,omitempty"`

	// 如果是rds类型， 那么可以选择'count', 'average_time', 'max_time', 'all'这几种类型
	RdsRecommendationType *string `json:"rds_recommendation_type,omitempty"`

	// 如果选择了taurus类型，那么可以选择'count', 'avg_time'
	TaurusRecommendationType *string `json:"taurus_recommendation_type,omitempty"`

	// 推荐数量
	RecommendCount *int64 `json:"recommend_count,omitempty"`

	// 如果选择了taurus， 那么需要制定node id
	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowRecommendSqlLimitRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecommendSqlLimitRuleRequestBody struct{}"
	}

	return strings.Join([]string{"ShowRecommendSqlLimitRuleRequestBody", string(data)}, " ")
}
