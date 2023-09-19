package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSwitchResponse Response Object
type UpdateSwitchResponse struct {
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateSwitchResponse", string(data)}, " ")
}
