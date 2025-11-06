package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIntelligentKillSessionHistoryRequest Request Object
type ShowIntelligentKillSessionHistoryRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 查询开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 查询页编号
	PageNum *int32 `json:"page_num,omitempty"`

	// 查询分页大小
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ShowIntelligentKillSessionHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIntelligentKillSessionHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowIntelligentKillSessionHistoryRequest", string(data)}, " ")
}
