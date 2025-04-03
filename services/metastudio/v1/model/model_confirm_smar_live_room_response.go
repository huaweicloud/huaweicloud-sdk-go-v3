package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmSmarLiveRoomResponse Response Object
type ConfirmSmarLiveRoomResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ConfirmSmarLiveRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmSmarLiveRoomResponse struct{}"
	}

	return strings.Join([]string{"ConfirmSmarLiveRoomResponse", string(data)}, " ")
}
