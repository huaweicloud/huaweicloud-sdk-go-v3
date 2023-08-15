package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStaticRouteResponse Response Object
type DeleteStaticRouteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStaticRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStaticRouteResponse struct{}"
	}

	return strings.Join([]string{"DeleteStaticRouteResponse", string(data)}, " ")
}
