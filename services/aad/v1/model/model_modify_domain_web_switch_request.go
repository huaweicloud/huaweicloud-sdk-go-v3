package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDomainWebSwitchRequest Request Object
type ModifyDomainWebSwitchRequest struct {
	Body *CadDomainSwitchRequest `json:"body,omitempty"`
}

func (o ModifyDomainWebSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDomainWebSwitchRequest struct{}"
	}

	return strings.Join([]string{"ModifyDomainWebSwitchRequest", string(data)}, " ")
}
