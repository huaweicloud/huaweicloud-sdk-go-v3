package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApplicationsRequest struct {
	// 偏移量。  偏移量为一个大于0小于资源总个数的整数，表示查询该偏移量后面的所有的资源，默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 查询的数量限制。  取值范围：1~100，取值一般为10，20，50。功能说明：每页返回的资源个数。默认值为100。

	Limit *int32 `json:"limit,omitempty"`
	// 检索应用名称。  支持后向模糊匹配，长度不得超过64个字符，只能包含英文字母、下划线和数字。

	Name *string `json:"name,omitempty"`
	// 应用平台。  目前仅支持HMS、APNS、APNS_SANDBOX。  HMS是为开发者提供的消息推送平台。  APNS和APNS_SANDBOX是用于推送iOS消息的服务平台。

	Platform *string `json:"platform,omitempty"`
}

func (o ListApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationsRequest", string(data)}, " ")
}
