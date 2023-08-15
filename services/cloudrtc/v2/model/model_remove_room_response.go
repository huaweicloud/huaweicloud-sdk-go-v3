package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveRoomResponse Response Object
type RemoveRoomResponse struct {
	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RemoveRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveRoomResponse struct{}"
	}

	return strings.Join([]string{"RemoveRoomResponse", string(data)}, " ")
}
