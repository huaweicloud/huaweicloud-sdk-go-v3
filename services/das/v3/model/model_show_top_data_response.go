package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopDataResponse Response Object
type ShowTopDataResponse struct {

	// Top库表数据列表
	TopDataList *[]TopDataInfo `json:"top_data_list,omitempty"`

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 采集时间（Unix timestamp），单位：毫秒
	CollectTimestamp *int64 `json:"collect_timestamp,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o ShowTopDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopDataResponse struct{}"
	}

	return strings.Join([]string{"ShowTopDataResponse", string(data)}, " ")
}
