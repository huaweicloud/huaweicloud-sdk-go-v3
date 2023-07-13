package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserSettingRequest Request Object
type UpdateUserSettingRequest struct {

	// 用户id
	UserId string `json:"user_id"`

	Body *UpdateUserSettingReq `json:"body,omitempty"`
}

func (o UpdateUserSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserSettingRequest", string(data)}, " ")
}
