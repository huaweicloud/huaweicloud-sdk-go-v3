package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHealthMonitorsResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty" xml:"page_info"`

	// 健康检查对象。
	Healthmonitors *[]HealthMonitor `json:"healthmonitors,omitempty" xml:"healthmonitors"`
	HttpStatusCode int              `json:"-"`
}

func (o ListHealthMonitorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthMonitorsResponse struct{}"
	}

	return strings.Join([]string{"ListHealthMonitorsResponse", string(data)}, " ")
}
