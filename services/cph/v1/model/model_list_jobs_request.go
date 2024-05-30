package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// 任务下发请求时响应的request_id。 request_id和request_ids必须指定其中一个。request_id和request_ids同时指定的时候，以request_ids为准。
	RequestId *string `json:"request_id,omitempty"`

	// 任务下发请求时响应的多个request_id，用逗号分隔，最多不能超过20个。 request_id和request_ids必须指定其中一个。request_id和request_ids同时指定的时候，以request_ids为准。
	RequestIds *string `json:"request_ids,omitempty"`

	// 偏移量为一个大于等于0整数，表示查询该偏移量后面的所有的资源数，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页返回的资源个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
