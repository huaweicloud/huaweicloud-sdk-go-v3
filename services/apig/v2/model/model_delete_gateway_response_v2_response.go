package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteGatewayResponseV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGatewayResponseV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteGatewayResponseV2Response struct{}"
	}

	return strings.Join([]string{"DeleteGatewayResponseV2Response", string(data)}, " ")
}
