package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectTrackerResponse Response Object
type ShowProjectTrackerResponse struct {

	// 追踪器配置转储生命周期
	BucketLifecycle *int32 `json:"bucket_lifecycle,omitempty"`

	// 追踪器配置监听事件
	DataEvent      *[]string `json:"data_event,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowProjectTrackerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTrackerResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectTrackerResponse", string(data)}, " ")
}
