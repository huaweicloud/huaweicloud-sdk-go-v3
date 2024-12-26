package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmSmartLiveRoomResponse Response Object
type ConfirmSmartLiveRoomResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ConfirmSmartLiveRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmSmartLiveRoomResponse struct{}"
	}

	return strings.Join([]string{"ConfirmSmartLiveRoomResponse", string(data)}, " ")
}
