package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregationAuthorizationResp 资源聚合器授权。
type AggregationAuthorizationResp struct {

	// 资源聚合器授权标识符。
	AggregationAuthorizationUrn *string `json:"aggregation_authorization_urn,omitempty"`

	// 授权的资源聚合器的帐号ID。
	AuthorizedAccountId *string `json:"authorized_account_id,omitempty"`

	// 资源聚合器授权的创建时间。
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o AggregationAuthorizationResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregationAuthorizationResp struct{}"
	}

	return strings.Join([]string{"AggregationAuthorizationResp", string(data)}, " ")
}
