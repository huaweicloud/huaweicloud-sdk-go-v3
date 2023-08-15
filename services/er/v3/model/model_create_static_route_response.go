package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStaticRouteResponse Response Object
type CreateStaticRouteResponse struct {
	Route *Route `json:"route,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	XClientToken   *string `json:"X-Client-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStaticRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStaticRouteResponse struct{}"
	}

	return strings.Join([]string{"CreateStaticRouteResponse", string(data)}, " ")
}
