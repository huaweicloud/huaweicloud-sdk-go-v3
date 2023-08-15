package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddIpGroupIpResponse Response Object
type AddIpGroupIpResponse struct {

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddIpGroupIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIpGroupIpResponse struct{}"
	}

	return strings.Join([]string{"AddIpGroupIpResponse", string(data)}, " ")
}
