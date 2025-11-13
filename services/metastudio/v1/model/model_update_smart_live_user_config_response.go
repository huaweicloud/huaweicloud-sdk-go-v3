package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSmartLiveUserConfigResponse Response Object
type UpdateSmartLiveUserConfigResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSmartLiveUserConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSmartLiveUserConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateSmartLiveUserConfigResponse", string(data)}, " ")
}
