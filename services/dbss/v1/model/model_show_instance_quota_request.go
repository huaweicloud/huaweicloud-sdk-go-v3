package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceQuotaRequest Request Object
type ShowInstanceQuotaRequest struct {
}

func (o ShowInstanceQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceQuotaRequest", string(data)}, " ")
}
