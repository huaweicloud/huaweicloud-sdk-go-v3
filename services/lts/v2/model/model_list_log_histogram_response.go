package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLogHistogramResponse struct {
	// 直方图结果

	Histogram *string `json:"histogram,omitempty"`
	// 日志条数

	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLogHistogramResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogHistogramResponse struct{}"
	}

	return strings.Join([]string{"ListLogHistogramResponse", string(data)}, " ")
}
