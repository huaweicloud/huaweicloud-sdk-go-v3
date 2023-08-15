package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRouteRequestBody This is a auto create Body Object
type UpdateRouteRequestBody struct {
	Route *UpdateRoute `json:"route"`
}

func (o UpdateRouteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRouteRequestBody", string(data)}, " ")
}
