package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteReservedInstanceByIdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteReservedInstanceByIdResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteReservedInstanceByIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteReservedInstanceByIdResponse", string(data)}, " ")
}
