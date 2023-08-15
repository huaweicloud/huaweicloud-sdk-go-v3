package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCheckTaskJobsResponse Response Object
type RunCheckTaskJobsResponse struct {

	// 调用成功时表示调用结果。  调用失败时无此字段。
	Result *[]CheckTaskJobsItemsBody `json:"result,omitempty"`

	// 符合查询条件的总任务数量。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o RunCheckTaskJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCheckTaskJobsResponse struct{}"
	}

	return strings.Join([]string{"RunCheckTaskJobsResponse", string(data)}, " ")
}
