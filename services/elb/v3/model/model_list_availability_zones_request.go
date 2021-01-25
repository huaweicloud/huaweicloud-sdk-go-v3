package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAvailabilityZonesRequest struct {
}

func (o ListAvailabilityZonesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesRequest", string(data)}, " ")
}
