package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHealthMonitorsResponse Response Object
type ListHealthMonitorsResponse struct {

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// **参数解释**：健康检查对象。
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
