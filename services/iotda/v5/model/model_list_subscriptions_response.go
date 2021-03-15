package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSubscriptionsResponse struct {
	// 订阅配置信息列表。
	Subscriptions *[]SubscriptionItem `json:"subscriptions,omitempty"`
	// 满足查询条件的记录总数。
	Count *int32 `json:"count,omitempty"`
	// 本次分页查询结果中最后一条记录的ID，可在下一次分页查询时使用。
	Marker         *string `json:"marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSubscriptionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubscriptionsResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsResponse", string(data)}, " ")
}
