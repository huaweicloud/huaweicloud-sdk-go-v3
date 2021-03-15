package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSignatureKeysV2Request struct {
	InstanceId    string  `json:"instance_id"`
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	Offset        *int64  `json:"offset,omitempty"`
	Limit         *int32  `json:"limit,omitempty"`
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListSignatureKeysV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSignatureKeysV2Request struct{}"
	}

	return strings.Join([]string{"ListSignatureKeysV2Request", string(data)}, " ")
}
