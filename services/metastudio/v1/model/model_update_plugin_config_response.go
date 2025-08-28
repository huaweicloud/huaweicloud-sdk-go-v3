package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePluginConfigResponse Response Object
type UpdatePluginConfigResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePluginConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdatePluginConfigResponse", string(data)}, " ")
}
