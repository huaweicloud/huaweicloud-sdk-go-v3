package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSnapshotsResponse Response Object
type BatchDeleteSnapshotsResponse struct {

	// 响应信息。
	Message *string `json:"message,omitempty"`

	// 响应码。
	Code *string `json:"code,omitempty"`

	// 总数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 成功数量。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 失败数量。
	FailureCount *int32 `json:"failure_count,omitempty"`

	// 失败详情列表。
	FailureDetails *[]FailureDetail `json:"failure_details,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchDeleteSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSnapshotsResponse", string(data)}, " ")
}
