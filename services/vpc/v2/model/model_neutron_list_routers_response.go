package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronListRoutersResponse Response Object
type NeutronListRoutersResponse struct {

	// router对象列表
	Routers *[]NeutronRouter `json:"routers,omitempty"`

	// 分页信息
	RoutersLinks   *[]NeutronPageLink `json:"routers_links,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o NeutronListRoutersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListRoutersResponse struct{}"
	}

	return strings.Join([]string{"NeutronListRoutersResponse", string(data)}, " ")
}
