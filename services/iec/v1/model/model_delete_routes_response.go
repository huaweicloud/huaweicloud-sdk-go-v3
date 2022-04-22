package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteRoutesResponse struct {

	// 路由列表
	Routes         *[]Route `json:"routes,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o DeleteRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRoutesResponse struct{}"
	}

	return strings.Join([]string{"DeleteRoutesResponse", string(data)}, " ")
}
