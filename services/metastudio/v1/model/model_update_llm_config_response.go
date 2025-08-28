package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLlmConfigResponse Response Object
type UpdateLlmConfigResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLlmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLlmConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateLlmConfigResponse", string(data)}, " ")
}
