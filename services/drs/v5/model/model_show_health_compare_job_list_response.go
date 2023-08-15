package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHealthCompareJobListResponse Response Object
type ShowHealthCompareJobListResponse struct {

	// 总数。
	Count *int64 `json:"count,omitempty"`

	// 健康对比任务列表。
	CompareJobs    *[]HealthCompareJob `json:"compare_jobs,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowHealthCompareJobListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthCompareJobListResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthCompareJobListResponse", string(data)}, " ")
}
