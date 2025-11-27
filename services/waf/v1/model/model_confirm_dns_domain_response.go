package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmDnsDomainResponse Response Object
type ConfirmDnsDomainResponse struct {
	NextMarker *string `json:"next_marker,omitempty"`

	Items          *[]DnsDomain `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ConfirmDnsDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmDnsDomainResponse struct{}"
	}

	return strings.Join([]string{"ConfirmDnsDomainResponse", string(data)}, " ")
}
