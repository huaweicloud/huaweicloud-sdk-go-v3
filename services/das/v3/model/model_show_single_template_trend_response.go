package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleTemplateTrendResponse Response Object
type ShowSingleTemplateTrendResponse struct {

	// 趋势图的时间戳
	Timestamps *[]int64 `json:"timestamps,omitempty"`

	// SQL趋势列表
	TrendDataList  *[]SingleSqlTplCmp `json:"trend_data_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowSingleTemplateTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleTemplateTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowSingleTemplateTrendResponse", string(data)}, " ")
}
