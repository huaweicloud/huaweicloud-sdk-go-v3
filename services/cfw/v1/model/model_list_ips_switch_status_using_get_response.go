package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsSwitchStatusUsingGetResponse Response Object
type ListIpsSwitchStatusUsingGetResponse struct {
	Data           *IpsSwitchResponseDto `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListIpsSwitchStatusUsingGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsSwitchStatusUsingGetResponse struct{}"
	}

	return strings.Join([]string{"ListIpsSwitchStatusUsingGetResponse", string(data)}, " ")
}
