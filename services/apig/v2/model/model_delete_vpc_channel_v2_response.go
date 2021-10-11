package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteVpcChannelV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcChannelV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVpcChannelV2Response struct{}"
	}

	return strings.Join([]string{"DeleteVpcChannelV2Response", string(data)}, " ")
}
