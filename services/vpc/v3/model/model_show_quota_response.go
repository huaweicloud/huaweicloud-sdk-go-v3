package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaResponse Response Object
type ShowQuotaResponse struct {
	Quota *VpcQuotas `json:"quota,omitempty"`

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaResponse", string(data)}, " ")
}
