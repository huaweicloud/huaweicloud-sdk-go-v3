package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchCpcsTokenRequest Request Object
type SwitchCpcsTokenRequest struct {
	Body *AuthObject `json:"body,omitempty"`
}

func (o SwitchCpcsTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchCpcsTokenRequest struct{}"
	}

	return strings.Join([]string{"SwitchCpcsTokenRequest", string(data)}, " ")
}
