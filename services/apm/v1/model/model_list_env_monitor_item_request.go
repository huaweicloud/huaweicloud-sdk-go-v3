package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvMonitorItemRequest Request Object
type ListEnvMonitorItemRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *GetEnvMonitorItemListParam `json:"body,omitempty"`
}

func (o ListEnvMonitorItemRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvMonitorItemRequest struct{}"
	}

	return strings.Join([]string{"ListEnvMonitorItemRequest", string(data)}, " ")
}
