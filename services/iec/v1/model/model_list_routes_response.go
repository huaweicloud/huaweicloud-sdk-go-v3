package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRoutesResponse Response Object
type ListRoutesResponse struct {

	// 路由列表
	Routes         *[]Route `json:"routes,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoutesResponse struct{}"
	}

	return strings.Join([]string{"ListRoutesResponse", string(data)}, " ")
}
