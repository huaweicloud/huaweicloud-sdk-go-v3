package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcEndpointServiceResponse Response Object
type ListVpcEndpointServiceResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 数据
	Records        *[]VpcepServiceData `json:"records,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListVpcEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"ListVpcEndpointServiceResponse", string(data)}, " ")
}
