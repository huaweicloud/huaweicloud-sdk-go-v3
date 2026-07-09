package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadDatabaseSslFileRequest struct {

	// 私钥文本内容
	PemKeyTxt string `json:"pem_key_txt"`
}

func (o UploadDatabaseSslFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadDatabaseSslFileRequest struct{}"
	}

	return strings.Join([]string{"UploadDatabaseSslFileRequest", string(data)}, " ")
}
