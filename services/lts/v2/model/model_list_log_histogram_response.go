package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogHistogramResponse Response Object
type ListLogHistogramResponse struct {

	// **参数解释：** 直方图结果。 **取值范围：** 不涉及。
	Histogram *[]LogCounts `json:"histogram,omitempty"`

	// 日志条数
	Count *int64 `json:"count,omitempty"`

	// 是否查询完成。
	IsQueryComplete *bool `json:"isQueryComplete,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ListLogHistogramResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogHistogramResponse struct{}"
	}

	return strings.Join([]string{"ListLogHistogramResponse", string(data)}, " ")
}
