package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 恢复凭据对象请求体
type UploadSecretBlobRequestBody struct {

	// 将指定凭据对象进行备份后得到的凭据备份文件，备份文件包含有凭据当前所有的凭据版本信息，备份文件经过加密与编码，内容不可直接读。
	SecretBlob string `json:"secret_blob"`
}

func (o UploadSecretBlobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSecretBlobRequestBody struct{}"
	}

	return strings.Join([]string{"UploadSecretBlobRequestBody", string(data)}, " ")
}
