package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSubscriptionsRequest struct {

	// 偏移量。  偏移量为一个大于0小于资源总个数的整数，表示查询该偏移量后面的所有的资源，默认值为0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	//  查询数量限制。  取值范围：1~100，取值一般为10，20，50。功能说明：每页返回的资源个数。默认值为100。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 协议名称， 枚举值：http、https、sms、email、functionstage、dms、application。
	Protocol *string `json:"protocol,omitempty" xml:"protocol"`

	// 状态。 0：未确认 1：已确认 2：不需要确认 3：已取消确认 4：已经删除。
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 订阅终端。
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint"`
}

func (o ListSubscriptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsRequest", string(data)}, " ")
}
