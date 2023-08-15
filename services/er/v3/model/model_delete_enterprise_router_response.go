package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnterpriseRouterResponse Response Object
type DeleteEnterpriseRouterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEnterpriseRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnterpriseRouterResponse struct{}"
	}

	return strings.Join([]string{"DeleteEnterpriseRouterResponse", string(data)}, " ")
}
