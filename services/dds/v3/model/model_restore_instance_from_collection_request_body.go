package model

import (
	"encoding/json"

	"strings"
)

type RestoreInstanceFromCollectionRequestBody struct {
	// 数据库信息。

	RestoreCollections []RestoreInstanceFromCollectionRequestBodyRestoreCollections `json:"restore_collections"`
}

func (o RestoreInstanceFromCollectionRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionRequestBody", string(data)}, " ")
}
