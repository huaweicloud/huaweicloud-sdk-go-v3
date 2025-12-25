package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIncidentResponse Response Object
type CreateIncidentResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *IncidentDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIncidentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIncidentResponse struct{}"
	}

	return strings.Join([]string{"CreateIncidentResponse", string(data)}, " ")
}
