package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateStaticRouteResponse struct {
	Route *Route `json:"route,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateStaticRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStaticRouteResponse struct{}"
	}

	return strings.Join([]string{"UpdateStaticRouteResponse", string(data)}, " ")
}
