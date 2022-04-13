package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateFollow302SwitchResponse struct {
	FollowStatus   *Follow302StatusBody `json:"follow_status,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateFollow302SwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFollow302SwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateFollow302SwitchResponse", string(data)}, " ")
}
