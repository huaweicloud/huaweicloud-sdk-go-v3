package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveIpGroupIpResponse Response Object
type RemoveIpGroupIpResponse struct {

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RemoveIpGroupIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveIpGroupIpResponse struct{}"
	}

	return strings.Join([]string{"RemoveIpGroupIpResponse", string(data)}, " ")
}
