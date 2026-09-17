package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSensitiveOperationSwitchResponse Response Object
type UpdateSensitiveOperationSwitchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSensitiveOperationSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSensitiveOperationSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateSensitiveOperationSwitchResponse", string(data)}, " ")
}
