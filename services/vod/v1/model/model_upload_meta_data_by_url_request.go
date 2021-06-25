package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UploadMetaDataByUrlRequest struct {
	Body *UploadMetaDataByUrlReq `json:"body,omitempty"`
}

func (o UploadMetaDataByUrlRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UploadMetaDataByUrlRequest struct{}"
	}

	return strings.Join([]string{"UploadMetaDataByUrlRequest", string(data)}, " ")
}
