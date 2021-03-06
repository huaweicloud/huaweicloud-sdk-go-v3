package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DownloadBlockchainCertResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadBlockchainCertResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DownloadBlockchainCertResponse struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainCertResponse", string(data)}, " ")
}
