package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchRouteResponse Response Object
type SwitchRouteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchRouteResponse struct{}"
	}

	return strings.Join([]string{"SwitchRouteResponse", string(data)}, " ")
}
