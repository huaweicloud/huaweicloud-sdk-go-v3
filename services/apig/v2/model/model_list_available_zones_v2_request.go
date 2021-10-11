package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAvailableZonesV2Request struct {
}

func (o ListAvailableZonesV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAvailableZonesV2Request struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesV2Request", string(data)}, " ")
}
