package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRouteResponse Response Object
type UpdateRouteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteResponse struct{}"
	}

	return strings.Join([]string{"UpdateRouteResponse", string(data)}, " ")
}
