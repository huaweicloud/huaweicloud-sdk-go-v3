package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookMonitorsResponse Response Object
type ShowPlaybookMonitorsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *PlaybookInstanceMonitorDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPlaybookMonitorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookMonitorsResponse struct{}"
	}

	return strings.Join([]string{"ShowPlaybookMonitorsResponse", string(data)}, " ")
}
