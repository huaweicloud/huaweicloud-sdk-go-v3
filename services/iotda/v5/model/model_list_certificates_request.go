package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCertificatesRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	AppId      *string `json:"app_id,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
	Marker     *string `json:"marker,omitempty"`
	Offset     *int32  `json:"offset,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}
