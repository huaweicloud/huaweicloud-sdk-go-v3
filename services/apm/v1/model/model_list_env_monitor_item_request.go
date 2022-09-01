package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEnvMonitorItemRequest struct {
	XBusinessId *int64 `json:"x-business-id,omitempty" xml:"x-business-id"`

	Body *GetEnvMonitorItemListParam `json:"body,omitempty" xml:"body"`
}

func (o ListEnvMonitorItemRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvMonitorItemRequest struct{}"
	}

	return strings.Join([]string{"ListEnvMonitorItemRequest", string(data)}, " ")
}
