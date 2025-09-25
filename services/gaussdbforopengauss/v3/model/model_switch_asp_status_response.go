package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAspStatusResponse Response Object
type SwitchAspStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchAspStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAspStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchAspStatusResponse", string(data)}, " ")
}
