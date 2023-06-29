package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoutesResponse Response Object
type CreateRoutesResponse struct {

	// 路由列表
	Routes         *[]Route `json:"routes,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o CreateRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutesResponse struct{}"
	}

	return strings.Join([]string{"CreateRoutesResponse", string(data)}, " ")
}
