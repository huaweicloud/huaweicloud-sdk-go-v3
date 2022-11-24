package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeIpsSwitchUsingPostRequest struct {
	Body *IpsSwitchDto `json:"body,omitempty"`
}

func (o ChangeIpsSwitchUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIpsSwitchUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ChangeIpsSwitchUsingPostRequest", string(data)}, " ")
}
