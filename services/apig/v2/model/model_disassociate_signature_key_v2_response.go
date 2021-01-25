package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisassociateSignatureKeyV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateSignatureKeyV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"DisassociateSignatureKeyV2Response", string(data)}, " ")
}
