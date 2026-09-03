package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlTemplateTrendResponse Response Object
type ShowSqlTemplateTrendResponse struct {

	// 聚合毫秒数
	IntervalMillis *int64 `json:"interval_millis,omitempty"`

	// 趋势列表
	ItemList       *[]SqlTplTrendItem `json:"item_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowSqlTemplateTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlTemplateTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlTemplateTrendResponse", string(data)}, " ")
}
