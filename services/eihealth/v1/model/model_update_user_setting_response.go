package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateUserSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateUserSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserSettingResponse", string(data)}, " ")
}
