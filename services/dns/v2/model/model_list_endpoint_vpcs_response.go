package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEndpointVpcsResponse struct {

	// 查询公网Zone的列表响应。
	Vpcs *[]VpcsData `json:"vpcs,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListEndpointVpcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointVpcsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointVpcsResponse", string(data)}, " ")
}
