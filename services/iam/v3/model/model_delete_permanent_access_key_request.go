package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePermanentAccessKeyRequest struct {
	// 待删除的指定AK。

	AccessKey string `json:"access_key"`
}

func (o DeletePermanentAccessKeyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePermanentAccessKeyRequest struct{}"
	}

	return strings.Join([]string{"DeletePermanentAccessKeyRequest", string(data)}, " ")
}
