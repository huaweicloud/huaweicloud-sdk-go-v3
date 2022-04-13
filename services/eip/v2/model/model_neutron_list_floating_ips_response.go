package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronListFloatingIpsResponse struct {
	// floatingip对象列表

	Floatingips *[]FloatingIpResp `json:"floatingips,omitempty"`
	// marker分页结构

	FloatingipsLinks *[]Pager `json:"floatingips_links,omitempty"`
	HttpStatusCode   int      `json:"-"`
}

func (o NeutronListFloatingIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListFloatingIpsResponse struct{}"
	}

	return strings.Join([]string{"NeutronListFloatingIpsResponse", string(data)}, " ")
}
