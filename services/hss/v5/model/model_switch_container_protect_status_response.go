package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchContainerProtectStatusResponse Response Object
type SwitchContainerProtectStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchContainerProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchContainerProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchContainerProtectStatusResponse", string(data)}, " ")
}
