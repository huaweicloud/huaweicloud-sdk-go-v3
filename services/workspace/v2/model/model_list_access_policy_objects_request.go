package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPolicyObjectsRequest Request Object
type ListAccessPolicyObjectsRequest struct {

	// 接入策略id。
	AccessPolicyId string `json:"access_policy_id"`

	// 每页数量,范围0-2000,默认10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量,范围0-1999,默认0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAccessPolicyObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPolicyObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListAccessPolicyObjectsRequest", string(data)}, " ")
}
