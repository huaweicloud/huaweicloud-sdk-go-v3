package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateIpGroupResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty" xml:"ipgroup"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateIpGroupResponse", string(data)}, " ")
}
