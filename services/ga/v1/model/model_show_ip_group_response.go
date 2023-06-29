package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpGroupResponse Response Object
type ShowIpGroupResponse struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	IpGroup        *IpGroupDetail `json:"ip_group,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowIpGroupResponse", string(data)}, " ")
}
