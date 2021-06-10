package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPartnersServiceInfoResponse struct {
	PartnersServiceInfo *PartnersServiceInfo `json:"partners_service_info,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowPartnersServiceInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPartnersServiceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowPartnersServiceInfoResponse", string(data)}, " ")
}
