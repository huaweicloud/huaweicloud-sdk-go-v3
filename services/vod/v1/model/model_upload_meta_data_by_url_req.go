package model

import (
	"encoding/json"

	"strings"
)

type UploadMetaDataByUrlReq struct {
	// 待拉取创建的媒资元数据

	UploadMetadatas []UploadMetaDataByUrl `json:"upload_metadatas"`
}

func (o UploadMetaDataByUrlReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UploadMetaDataByUrlReq struct{}"
	}

	return strings.Join([]string{"UploadMetaDataByUrlReq", string(data)}, " ")
}
