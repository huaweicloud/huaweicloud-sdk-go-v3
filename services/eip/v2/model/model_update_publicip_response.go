package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePublicipResponse struct {
	Publicip       *PublicipUpdateResp `json:"publicip,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdatePublicipResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePublicipResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicipResponse", string(data)}, " ")
}
