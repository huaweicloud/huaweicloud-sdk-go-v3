package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIncidentResponse Response Object
type DeleteIncidentResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *DeleteIncidentResponseBodyData `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteIncidentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIncidentResponse struct{}"
	}

	return strings.Join([]string{"DeleteIncidentResponse", string(data)}, " ")
}
