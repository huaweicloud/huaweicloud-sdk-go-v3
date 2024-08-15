package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWarRoomsRequest Request Object
type ListWarRoomsRequest struct {
	Body *ListTenantWarRoomRequestBody `json:"body,omitempty"`
}

func (o ListWarRoomsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWarRoomsRequest struct{}"
	}

	return strings.Join([]string{"ListWarRoomsRequest", string(data)}, " ")
}
