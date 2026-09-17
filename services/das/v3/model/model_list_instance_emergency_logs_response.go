package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceEmergencyLogsResponse Response Object
type ListInstanceEmergencyLogsResponse struct {

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 数据列表
	Data *[]DdsEmergencyLogInfo `json:"data,omitempty"`

	// 对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// 采集时间
	CollectDate    *string `json:"collect_date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInstanceEmergencyLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceEmergencyLogsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceEmergencyLogsResponse", string(data)}, " ")
}
