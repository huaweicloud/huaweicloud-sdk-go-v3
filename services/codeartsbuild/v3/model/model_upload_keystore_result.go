package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadKeystoreResult 返回结果
type UploadKeystoreResult struct {

	// 文件ID
	KeystoreId *string `json:"keystore_id,omitempty"`
}

func (o UploadKeystoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadKeystoreResult struct{}"
	}

	return strings.Join([]string{"UploadKeystoreResult", string(data)}, " ")
}
