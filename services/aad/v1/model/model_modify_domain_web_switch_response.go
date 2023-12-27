package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDomainWebSwitchResponse Response Object
type ModifyDomainWebSwitchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyDomainWebSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDomainWebSwitchResponse struct{}"
	}

	return strings.Join([]string{"ModifyDomainWebSwitchResponse", string(data)}, " ")
}
