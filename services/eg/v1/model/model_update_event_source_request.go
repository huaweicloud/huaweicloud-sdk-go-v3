package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEventSourceRequest struct {

	// 指定查询的事件源ID
	SourceId string `json:"source_id"`

	Body *CustomizeSourceUpdateReq `json:"body,omitempty"`
}

func (o UpdateEventSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEventSourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateEventSourceRequest", string(data)}, " ")
}
