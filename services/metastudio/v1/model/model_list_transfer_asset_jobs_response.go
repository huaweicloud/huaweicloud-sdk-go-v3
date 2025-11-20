package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransferAssetJobsResponse Response Object
type ListTransferAssetJobsResponse struct {

	// 任务总数
	Count *int32 `json:"count,omitempty"`

	// 任务信息
	Jobs *[]TransferAssetJobInfo `json:"jobs,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTransferAssetJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransferAssetJobsResponse struct{}"
	}

	return strings.Join([]string{"ListTransferAssetJobsResponse", string(data)}, " ")
}
