package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAutoProtectStatusResponse Response Object
type SwitchAutoProtectStatusResponse struct {
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchAutoProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAutoProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchAutoProtectStatusResponse", string(data)}, " ")
}
