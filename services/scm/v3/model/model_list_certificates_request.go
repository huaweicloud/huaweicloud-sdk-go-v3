package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCertificatesRequest struct {
	Limit   *int32  `json:"limit,omitempty"`
	Offset  *int32  `json:"offset,omitempty"`
	SortDir *string `json:"sort_dir,omitempty"`
	SortKey *string `json:"sort_key,omitempty"`
	Status  *string `json:"status,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}
