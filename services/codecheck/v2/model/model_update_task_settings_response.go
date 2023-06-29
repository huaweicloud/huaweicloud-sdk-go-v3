package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskSettingsResponse Response Object
type UpdateTaskSettingsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTaskSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskSettingsResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskSettingsResponse", string(data)}, " ")
}
