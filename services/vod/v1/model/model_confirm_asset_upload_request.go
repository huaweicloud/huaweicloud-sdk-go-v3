package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ConfirmAssetUploadRequest struct {
	Body *ConfirmAssetUploadReq `json:"body,omitempty"`
}

func (o ConfirmAssetUploadRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConfirmAssetUploadRequest struct{}"
	}

	return strings.Join([]string{"ConfirmAssetUploadRequest", string(data)}, " ")
}
