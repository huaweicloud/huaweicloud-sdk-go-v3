package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrendResponse Response Object
type ShowTrendResponse struct {

	// 趋势图数据列表。
	LineList *[]FrontLine `json:"line_list,omitempty"`

	// 最后日期时间。
	LatestDataTime *int64 `json:"latest_data_Time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowTrendResponse", string(data)}, " ")
}
