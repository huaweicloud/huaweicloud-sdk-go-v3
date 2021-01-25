package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisassociateCertificateV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateCertificateV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateCertificateV2Response struct{}"
	}

	return strings.Join([]string{"DisassociateCertificateV2Response", string(data)}, " ")
}
