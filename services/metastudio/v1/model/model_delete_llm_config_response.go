package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLlmConfigResponse Response Object
type DeleteLlmConfigResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLlmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLlmConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteLlmConfigResponse", string(data)}, " ")
}
