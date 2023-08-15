package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHealthMonitorsResponse Response Object
type ListHealthMonitorsResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 健康检查对象。
	Healthmonitors *[]HealthMonitor `json:"healthmonitors,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListHealthMonitorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthMonitorsResponse struct{}"
	}

	return strings.Join([]string{"ListHealthMonitorsResponse", string(data)}, " ")
}
