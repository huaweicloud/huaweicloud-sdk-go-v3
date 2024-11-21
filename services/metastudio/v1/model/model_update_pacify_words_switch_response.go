package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePacifyWordsSwitchResponse Response Object
type UpdatePacifyWordsSwitchResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePacifyWordsSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePacifyWordsSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdatePacifyWordsSwitchResponse", string(data)}, " ")
}
