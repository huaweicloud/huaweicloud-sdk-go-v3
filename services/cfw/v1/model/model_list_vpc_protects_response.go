package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListVpcProtectsResponse struct {

	// 调用链id
	TraceId *string `json:"trace_id,omitempty"`

	Data           *VpcProtectsVo `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListVpcProtectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcProtectsResponse struct{}"
	}

	return strings.Join([]string{"ListVpcProtectsResponse", string(data)}, " ")
}
