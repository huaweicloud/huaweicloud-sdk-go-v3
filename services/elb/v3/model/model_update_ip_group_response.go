package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupResponse Response Object
type UpdateIpGroupResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`

	// 参数解释：请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupResponse", string(data)}, " ")
}
