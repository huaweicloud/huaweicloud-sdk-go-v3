package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateIncidentResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
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
