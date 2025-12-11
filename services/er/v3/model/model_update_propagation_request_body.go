package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePropagationRequestBody struct {
	RoutePolicy *ImportRoutePolicy `json:"route_policy"`
}

func (o UpdatePropagationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePropagationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePropagationRequestBody", string(data)}, " ")
}
