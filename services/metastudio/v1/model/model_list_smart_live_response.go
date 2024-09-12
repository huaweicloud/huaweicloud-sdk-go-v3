package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartLiveResponse Response Object
type ListSmartLiveResponse struct {

	// **参数解释**： 数字人直播任务总数。
	Count *int32 `json:"count,omitempty"`

	// 数字人直播任务列表。
	SmartLiveJobs *[]SmartLiveJob `json:"smart_live_jobs,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSmartLiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartLiveResponse struct{}"
	}

	return strings.Join([]string{"ListSmartLiveResponse", string(data)}, " ")
}
