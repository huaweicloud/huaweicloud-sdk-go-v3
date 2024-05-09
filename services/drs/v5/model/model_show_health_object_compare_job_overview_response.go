package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHealthObjectCompareJobOverviewResponse Response Object
type ShowHealthObjectCompareJobOverviewResponse struct {

	// 健康对比对象级对比结果。
	CompareResult  *[]ObjectsHealthCompareOverviewInfo `json:"compare_result,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowHealthObjectCompareJobOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthObjectCompareJobOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthObjectCompareJobOverviewResponse", string(data)}, " ")
}
