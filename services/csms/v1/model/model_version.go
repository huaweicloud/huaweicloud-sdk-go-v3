package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Version 凭据版本。
type Version struct {
	VersionMetadata *VersionMetadata `json:"version_metadata,omitempty"`

	// 二进制类型凭据在base64编码后的明文，凭据管理服务将其加密后，存入凭据的初始版本中。  类型：base64编码的二进制数据对象。
	SecretBinary *string `json:"secret_binary,omitempty"`

	// 文本类型凭据的明文，凭据管理服务将其加密后，存入凭据的初始版本中。
	SecretString *string `json:"secret_string,omitempty"`
}

func (o Version) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Version struct{}"
	}

	return strings.Join([]string{"Version", string(data)}, " ")
}
