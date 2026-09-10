package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAvailableVpcsResponse Response Object
type GetAvailableVpcsResponse struct {

	// 可用的VPC列表
	Vpcs *[]Vpc `json:"vpcs,omitempty"`

	XTraceId       *string `json:"X-TRACE-ID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetAvailableVpcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAvailableVpcsResponse struct{}"
	}

	return strings.Join([]string{"GetAvailableVpcsResponse", string(data)}, " ")
}
