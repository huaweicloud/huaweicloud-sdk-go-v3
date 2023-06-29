package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpGroupResponse Response Object
type CreateIpGroupResponse struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	IpGroup        *IpGroupDetail `json:"ip_group,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateIpGroupResponse", string(data)}, " ")
}
