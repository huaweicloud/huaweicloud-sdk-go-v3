package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainIp6SwitchResponse Response Object
type UpdateDomainIp6SwitchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDomainIp6SwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainIp6SwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainIp6SwitchResponse", string(data)}, " ")
}
