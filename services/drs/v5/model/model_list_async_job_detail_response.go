package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAsyncJobDetailResponse struct {

	// 列表中的项目总数，与分页无关。
	TotalCount int32 `json:"total_count"`

	// 查询租户指定ID批量异步任务详情响应体。
	Jobs           []JobDetailResp `json:"jobs"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAsyncJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ListAsyncJobDetailResponse", string(data)}, " ")
}
