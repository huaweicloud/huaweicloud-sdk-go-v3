package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSubscriptionsRequest struct {
	// 偏移量。  偏移量为一个大于0小于资源总个数的整数，表示查询该偏移量后面的所有的资源，默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	//  查询数量限制。  取值范围：1~100，取值一般为10，20，50。功能说明：每页返回的资源个数。默认值为100。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSubscriptionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubscriptionsRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsRequest", string(data)}, " ")
}
