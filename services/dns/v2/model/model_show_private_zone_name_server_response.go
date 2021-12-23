package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPrivateZoneNameServerResponse struct {
	Nameservers    *[]PrivateNameServer `json:"nameservers,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowPrivateZoneNameServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneNameServerResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneNameServerResponse", string(data)}, " ")
}
