package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAppRequest struct {
	// 单次请求返回APP列表的最大数量。  取值范围：1~100。  默认值：10

	Limit *int32 `json:"limit,omitempty"`
	// 从该app名称开始返回app列表，返回的app列表不包括此app名称。

	StartAppName *string `json:"start_app_name,omitempty"`
	// 返回该通道下的app列表。

	StreamName *string `json:"stream_name,omitempty"`
}

func (o ListAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppRequest struct{}"
	}

	return strings.Join([]string{"ListAppRequest", string(data)}, " ")
}
