package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublications4SubscriptionResponse Response Object
type ListPublications4SubscriptionResponse struct {

	// 实例发布列表。
	InstancePublications *[]InstancePublicatiosInfo `json:"instance_publications,omitempty"`

	// 发布总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPublications4SubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublications4SubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ListPublications4SubscriptionResponse", string(data)}, " ")
}
