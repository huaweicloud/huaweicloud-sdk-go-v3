package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisassociateEndpointIpaddressResponse struct {

	// 查询公网Zone的列表响应。
	Endpoints      *[]EndpointResp `json:"endpoints,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o DisassociateEndpointIpaddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateEndpointIpaddressResponse struct{}"
	}

	return strings.Join([]string{"DisassociateEndpointIpaddressResponse", string(data)}, " ")
}
