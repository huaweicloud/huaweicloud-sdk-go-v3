package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAggregationAuthorizationRequest struct {

	// 授权的资源聚合器的帐号ID。
	AuthorizedAccountId string `json:"authorized_account_id"`
}

func (o DeleteAggregationAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAggregationAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"DeleteAggregationAuthorizationRequest", string(data)}, " ")
}
