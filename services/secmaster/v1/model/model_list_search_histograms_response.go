package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchHistogramsResponse Response Object
type ListSearchHistogramsResponse struct {

	// 直方图
	Histograms *[]Histogram `json:"histograms,omitempty"`

	// 数据总条数
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSearchHistogramsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchHistogramsResponse struct{}"
	}

	return strings.Join([]string{"ListSearchHistogramsResponse", string(data)}, " ")
}
