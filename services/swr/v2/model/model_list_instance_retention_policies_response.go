package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceRetentionPoliciesResponse Response Object
type ListInstanceRetentionPoliciesResponse struct {

	// 策略总数
	Total *int32 `json:"total,omitempty"`

	// 老化策略详情
	Retentions     *[]ImageRetention `json:"retentions,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListInstanceRetentionPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRetentionPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceRetentionPoliciesResponse", string(data)}, " ")
}
