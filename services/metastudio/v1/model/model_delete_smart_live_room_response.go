package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSmartLiveRoomResponse Response Object
type DeleteSmartLiveRoomResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSmartLiveRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSmartLiveRoomResponse struct{}"
	}

	return strings.Join([]string{"DeleteSmartLiveRoomResponse", string(data)}, " ")
}
