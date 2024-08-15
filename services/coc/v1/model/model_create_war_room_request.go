package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWarRoomRequest Request Object
type CreateWarRoomRequest struct {
	Body *CreateWarRoomRequestBody `json:"body,omitempty"`
}

func (o CreateWarRoomRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWarRoomRequest struct{}"
	}

	return strings.Join([]string{"CreateWarRoomRequest", string(data)}, " ")
}
