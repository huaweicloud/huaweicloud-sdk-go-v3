package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRegisterServiceRequest struct {
}

func (o ShowRegisterServiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRegisterServiceRequest struct{}"
	}

	return strings.Join([]string{"ShowRegisterServiceRequest", string(data)}, " ")
}
