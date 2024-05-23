package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHotWordsSwitchResponse Response Object
type UpdateHotWordsSwitchResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHotWordsSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotWordsSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateHotWordsSwitchResponse", string(data)}, " ")
}
