package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
