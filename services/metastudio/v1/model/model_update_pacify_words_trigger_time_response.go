package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePacifyWordsTriggerTimeResponse Response Object
type UpdatePacifyWordsTriggerTimeResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePacifyWordsTriggerTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePacifyWordsTriggerTimeResponse struct{}"
	}

	return strings.Join([]string{"UpdatePacifyWordsTriggerTimeResponse", string(data)}, " ")
}
