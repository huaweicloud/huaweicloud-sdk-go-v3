package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckImageSyncResponse Response Object
type BatchCheckImageSyncResponse struct {

	// 本次请求的唯⼀标识，⽤于问题排查，建议保存。
	RequestId *string `json:"request_id,omitempty"`

	// 调用结果。
	Results        *[]ImageBatchSyncDetectionResult `json:"results,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o BatchCheckImageSyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckImageSyncResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckImageSyncResponse", string(data)}, " ")
}
