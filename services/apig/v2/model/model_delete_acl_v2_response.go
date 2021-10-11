package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteAclV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAclV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAclV2Response struct{}"
	}

	return strings.Join([]string{"DeleteAclV2Response", string(data)}, " ")
}
