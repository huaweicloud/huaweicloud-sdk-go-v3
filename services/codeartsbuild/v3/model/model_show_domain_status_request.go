package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainStatusRequest Request Object
type ShowDomainStatusRequest struct {
}

func (o ShowDomainStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainStatusRequest", string(data)}, " ")
}
