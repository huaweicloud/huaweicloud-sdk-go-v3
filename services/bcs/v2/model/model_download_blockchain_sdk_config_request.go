package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DownloadBlockchainSdkConfigRequest struct {
	// blockchainID

	BlockchainId string `json:"blockchain_id"`

	Body *CfgRequestBody `json:"body,omitempty"`
}

func (o DownloadBlockchainSdkConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DownloadBlockchainSdkConfigRequest struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainSdkConfigRequest", string(data)}, " ")
}
