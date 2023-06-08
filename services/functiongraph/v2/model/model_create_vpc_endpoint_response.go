package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVpcEndpointResponse struct {

	// 依赖id列表
	State *[]string `json:"state,omitempty"`

	// 快照制作响应码
	Code           *string `json:"code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVpcEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcEndpointResponse", string(data)}, " ")
}
