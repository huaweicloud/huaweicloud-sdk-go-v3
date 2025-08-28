package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePluginConfigResponse Response Object
type DeletePluginConfigResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePluginConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePluginConfigResponse struct{}"
	}

	return strings.Join([]string{"DeletePluginConfigResponse", string(data)}, " ")
}
