package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasInfoRequest Request Object
type ShowQuotasInfoRequest struct {
}

func (o ShowQuotasInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasInfoRequest", string(data)}, " ")
}
