package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsSwitchStatusResponse Response Object
type ListIpsSwitchStatusResponse struct {
	Data           *IpsSwitchResponseDto `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListIpsSwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsSwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ListIpsSwitchStatusResponse", string(data)}, " ")
}
