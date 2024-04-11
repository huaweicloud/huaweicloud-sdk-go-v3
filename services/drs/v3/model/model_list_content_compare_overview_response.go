package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContentCompareOverviewResponse Response Object
type ListContentCompareOverviewResponse struct {

	// 对比数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 信息列表
	ContentCompareResultInfos *[]NodeContentCompareOverviewResult `json:"content_compare_result_infos,omitempty"`
	HttpStatusCode            int                                 `json:"-"`
}

func (o ListContentCompareOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContentCompareOverviewResponse struct{}"
	}

	return strings.Join([]string{"ListContentCompareOverviewResponse", string(data)}, " ")
}
