package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UploadMetaDataByUrlResponse struct {
	// 待拉取创建的媒资元数据

	UploadAssets   *[]UploadAsset `json:"upload_assets,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UploadMetaDataByUrlResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UploadMetaDataByUrlResponse struct{}"
	}

	return strings.Join([]string{"UploadMetaDataByUrlResponse", string(data)}, " ")
}
