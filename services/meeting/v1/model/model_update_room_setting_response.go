package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoomSettingResponse Response Object
type UpdateRoomSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRoomSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoomSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateRoomSettingResponse", string(data)}, " ")
}
