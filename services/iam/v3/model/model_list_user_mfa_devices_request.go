package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListUserMfaDevicesRequest struct {
}

func (o ListUserMfaDevicesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListUserMfaDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListUserMfaDevicesRequest", string(data)}, " ")
}
