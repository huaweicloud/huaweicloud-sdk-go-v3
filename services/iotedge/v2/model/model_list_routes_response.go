package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRoutesResponse struct {
	// 路由列表

	Routes *[]RouterDetailRespDto `json:"routes,omitempty"`
	// 最后一次修改时间

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoutesResponse struct{}"
	}

	return strings.Join([]string{"ListRoutesResponse", string(data)}, " ")
}
