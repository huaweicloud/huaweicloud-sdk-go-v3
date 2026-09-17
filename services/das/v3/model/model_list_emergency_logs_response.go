package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEmergencyLogsResponse Response Object
type ListEmergencyLogsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 数据
	Data *[]interface{} `json:"data,omitempty"`

	// 对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// 采集时间
	CollectDate    *string `json:"collect_date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEmergencyLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEmergencyLogsResponse struct{}"
	}

	return strings.Join([]string{"ListEmergencyLogsResponse", string(data)}, " ")
}
