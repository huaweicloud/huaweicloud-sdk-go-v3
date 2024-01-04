package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEngineQuotasRequest Request Object
type ShowEngineQuotasRequest struct {
}

func (o ShowEngineQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowEngineQuotasRequest", string(data)}, " ")
}
