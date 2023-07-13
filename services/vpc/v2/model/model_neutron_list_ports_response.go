package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronListPortsResponse Response Object
type NeutronListPortsResponse struct {

	// port对象列表
	Ports *[]NeutronPort `json:"ports,omitempty"`

	// 分页信息
	PortsLinks     *[]NeutronPageLink `json:"ports_links,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o NeutronListPortsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListPortsResponse struct{}"
	}

	return strings.Join([]string{"NeutronListPortsResponse", string(data)}, " ")
}
