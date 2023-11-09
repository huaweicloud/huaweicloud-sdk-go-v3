package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionMonitorResponse Response Object
type CreateConnectionMonitorResponse struct {
	ConnectionMonitor *CreateConnectionMonitorInfo `json:"connection_monitor,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateConnectionMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionMonitorResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionMonitorResponse", string(data)}, " ")
}
