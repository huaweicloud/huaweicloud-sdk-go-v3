package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DownloadSecretBlobRequest struct {

	// 凭据的名称。
	SecretName string `json:"secret_name" xml:"secret_name"`
}

func (o DownloadSecretBlobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSecretBlobRequest struct{}"
	}

	return strings.Join([]string{"DownloadSecretBlobRequest", string(data)}, " ")
}
