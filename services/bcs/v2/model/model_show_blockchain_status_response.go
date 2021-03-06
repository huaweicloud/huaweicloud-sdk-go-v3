package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowBlockchainStatusResponse struct {
	Bcs *Detail `json:"bcs,omitempty"`

	Eip *Detail `json:"eip,omitempty"`

	Sfs *Detail `json:"sfs,omitempty"`

	Obs *Detail `json:"obs,omitempty"`

	Kafka *Detail `json:"kafka,omitempty"`

	Cce            *ComCce `json:"cce,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBlockchainStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlockchainStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowBlockchainStatusResponse", string(data)}, " ")
}
