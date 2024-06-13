package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainInfoRequest Request Object
type ShowDomainInfoRequest struct {
}

func (o ShowDomainInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainInfoRequest", string(data)}, " ")
}
