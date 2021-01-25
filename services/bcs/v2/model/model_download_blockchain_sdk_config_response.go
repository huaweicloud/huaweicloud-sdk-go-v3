package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DownloadBlockchainSdkConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadBlockchainSdkConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DownloadBlockchainSdkConfigResponse struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainSdkConfigResponse", string(data)}, " ")
}
