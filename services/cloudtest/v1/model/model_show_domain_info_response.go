package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainInfoResponse Response Object
type ShowDomainInfoResponse struct {
	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDomainInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainInfoResponse", string(data)}, " ")
}
