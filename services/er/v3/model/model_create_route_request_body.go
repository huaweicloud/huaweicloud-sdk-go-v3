package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRouteRequestBody This is a auto create Body Object
type CreateRouteRequestBody struct {
	Route *CreateRoute `json:"route"`
}

func (o CreateRouteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRouteRequestBody", string(data)}, " ")
}
