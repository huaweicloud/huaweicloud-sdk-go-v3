package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAggregationAuthorizationsResponse struct {

	// 授权过的资源聚合器帐号列表。
	AggregationAuthorizations *[]AggregationAuthorizationResp `json:"aggregation_authorizations,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAggregationAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAggregationAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"ListAggregationAuthorizationsResponse", string(data)}, " ")
}
