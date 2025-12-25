package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableHistogramsResponse Response Object
type ListTableHistogramsResponse struct {

	// 总数
	TotalCount *int64 `json:"total_count,omitempty"`

	// 直方图
	Histograms     *[]SearchQueryResultHistogram `json:"histograms,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListTableHistogramsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableHistogramsResponse struct{}"
	}

	return strings.Join([]string{"ListTableHistogramsResponse", string(data)}, " ")
}
