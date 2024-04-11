package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataCompareDetailResponse Response Object
type ListDataCompareDetailResponse struct {

	// 对比数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 对比信息列表
	TableLineCompareResultInfos *[]TableLineCompareResultInfo `json:"table_line_compare_result_infos,omitempty"`
	HttpStatusCode              int                           `json:"-"`
}

func (o ListDataCompareDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataCompareDetailResponse struct{}"
	}

	return strings.Join([]string{"ListDataCompareDetailResponse", string(data)}, " ")
}
