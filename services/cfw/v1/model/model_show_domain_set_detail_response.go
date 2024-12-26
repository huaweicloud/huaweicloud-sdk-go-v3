package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainSetDetailResponse Response Object
type ShowDomainSetDetailResponse struct {
	Data           *DomainSetVo `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDomainSetDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainSetDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainSetDetailResponse", string(data)}, " ")
}
