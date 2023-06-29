package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartnersServiceInfoResponse Response Object
type ShowPartnersServiceInfoResponse struct {
	PartnersServiceInfo *PartnersServiceInfo `json:"partners_service_info,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowPartnersServiceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartnersServiceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowPartnersServiceInfoResponse", string(data)}, " ")
}
