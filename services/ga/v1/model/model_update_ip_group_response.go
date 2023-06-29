package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupResponse Response Object
type UpdateIpGroupResponse struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	IpGroup        *IpGroupDetail `json:"ip_group,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupResponse", string(data)}, " ")
}
