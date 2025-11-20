package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceJobsResponse Response Object
type ListResourceJobsResponse struct {

	// - 参数解释：查询任务列表的响应体。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Jobs *[]Job `json:"jobs,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// - 参数解释：请求的唯一标识。 - 约束限制：UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResourceJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceJobsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceJobsResponse", string(data)}, " ")
}
