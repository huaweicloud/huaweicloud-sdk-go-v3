package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlLimitingRecordResponse Response Object
type ShowSqlLimitingRecordResponse struct {

	// SQL限流规则列表
	SqlLimitingRecordList *[]SqlLimitingRecordInfo `json:"sql_limiting_record_list,omitempty"`

	// SQL限流规则总数
	Total *int32 `json:"total,omitempty"`

	// 实例是否支持展示显示限流触发次数
	CanShowLimitCount *bool `json:"can_show_limit_count,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o ShowSqlLimitingRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitingRecordResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitingRecordResponse", string(data)}, " ")
}
