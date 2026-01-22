package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigQuotaResponse Response Object
type ShowConfigQuotaResponse struct {
	Data           *ShowConfigQuotaDto `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowConfigQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigQuotaResponse", string(data)}, " ")
}
