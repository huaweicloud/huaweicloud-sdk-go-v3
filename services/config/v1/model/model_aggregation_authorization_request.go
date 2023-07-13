package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregationAuthorizationRequest 资源聚合器授权请求体。
type AggregationAuthorizationRequest struct {

	// 要授权的资源聚合器的帐号ID。
	AuthorizedAccountId string `json:"authorized_account_id"`
}

func (o AggregationAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregationAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"AggregationAuthorizationRequest", string(data)}, " ")
}
