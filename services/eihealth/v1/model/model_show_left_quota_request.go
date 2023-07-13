package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLeftQuotaRequest Request Object
type ShowLeftQuotaRequest struct {
}

func (o ShowLeftQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLeftQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowLeftQuotaRequest", string(data)}, " ")
}
