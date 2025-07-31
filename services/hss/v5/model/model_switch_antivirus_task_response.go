package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAntivirusTaskResponse Response Object
type SwitchAntivirusTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchAntivirusTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAntivirusTaskResponse struct{}"
	}

	return strings.Join([]string{"SwitchAntivirusTaskResponse", string(data)}, " ")
}
