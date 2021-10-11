package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteGatewayResponseTypeV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGatewayResponseTypeV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteGatewayResponseTypeV2Response struct{}"
	}

	return strings.Join([]string{"DeleteGatewayResponseTypeV2Response", string(data)}, " ")
}
