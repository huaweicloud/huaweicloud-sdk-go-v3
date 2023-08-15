package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronListNetworksResponse Response Object
type NeutronListNetworksResponse struct {

	// network对象列表
	Networks *[]NeutronNetwork `json:"networks,omitempty"`

	// 分页信息
	NetworksLinks  *[]NeutronPageLink `json:"networks_links,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o NeutronListNetworksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListNetworksResponse struct{}"
	}

	return strings.Join([]string{"NeutronListNetworksResponse", string(data)}, " ")
}
