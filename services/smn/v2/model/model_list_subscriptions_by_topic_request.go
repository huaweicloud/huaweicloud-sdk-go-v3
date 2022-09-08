package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSubscriptionsByTopicRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](https://support.huaweicloud.com/api-smn/smn_api_51004.html)获取该标识。
	TopicUrn string `json:"topic_urn"`

	// 偏移量。  偏移量为一个大于0小于资源总个数的整数，表示查询该偏移量后面的所有的资源，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量限制。  取值范围：1~100，取值一般为10，20，50。功能说明：每页返回的资源个数。默认值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSubscriptionsByTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsByTopicRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsByTopicRequest", string(data)}, " ")
}
