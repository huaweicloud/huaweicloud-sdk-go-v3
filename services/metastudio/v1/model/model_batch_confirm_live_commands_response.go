package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchConfirmLiveCommandsResponse Response Object
type BatchConfirmLiveCommandsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchConfirmLiveCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchConfirmLiveCommandsResponse struct{}"
	}

	return strings.Join([]string{"BatchConfirmLiveCommandsResponse", string(data)}, " ")
}
