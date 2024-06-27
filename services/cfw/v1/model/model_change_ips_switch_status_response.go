package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIpsSwitchStatusResponse Response Object
type ChangeIpsSwitchStatusResponse struct {
	Data           *CommonResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ChangeIpsSwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIpsSwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeIpsSwitchStatusResponse", string(data)}, " ")
}
