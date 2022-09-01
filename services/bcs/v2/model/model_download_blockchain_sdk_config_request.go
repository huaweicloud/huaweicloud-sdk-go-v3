package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DownloadBlockchainSdkConfigRequest struct {

	// blockchainID
	BlockchainId string `json:"blockchain_id" xml:"blockchain_id"`

	Body *CfgRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DownloadBlockchainSdkConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBlockchainSdkConfigRequest struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainSdkConfigRequest", string(data)}, " ")
}
