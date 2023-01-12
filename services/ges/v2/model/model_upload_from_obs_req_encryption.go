package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 元数据是否加密
type UploadFromObsReqEncryption struct {

	// 是否启动加密特性。
	Enable *bool `json:"enable,omitempty"`

	// 对应的project下，华为云数据加密服务创建的用户主密钥ID。
	MasterKeyId *string `json:"master_key_id,omitempty"`
}

func (o UploadFromObsReqEncryption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFromObsReqEncryption struct{}"
	}

	return strings.Join([]string{"UploadFromObsReqEncryption", string(data)}, " ")
}
