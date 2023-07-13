package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIncidentResponse Response Object
type ShowIncidentResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data           *ShowIncidentDetail `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowIncidentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIncidentResponse struct{}"
	}

	return strings.Join([]string{"ShowIncidentResponse", string(data)}, " ")
}
