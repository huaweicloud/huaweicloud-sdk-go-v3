package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContentCompareDetailResponse Response Object
type ListContentCompareDetailResponse struct {

	// 对比数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 对比信息列表
	ContentCompareResultInfos *[]NodeContentCompareDetailResult `json:"content_compare_result_infos,omitempty"`
	HttpStatusCode            int                               `json:"-"`
}

func (o ListContentCompareDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContentCompareDetailResponse struct{}"
	}

	return strings.Join([]string{"ListContentCompareDetailResponse", string(data)}, " ")
}
