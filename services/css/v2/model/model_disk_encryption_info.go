package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiskEncryptionInfo struct {

	// 是否开启磁盘加密。通过磁盘加密对集群节点的数据盘进行KMS加密，保障数据存储的安全性。
	SystemEncrypted *string `json:"systemEncrypted,omitempty"`

	// 磁盘密钥ID。
	SystemCmkid *string `json:"systemCmkid,omitempty"`
}

func (o DiskEncryptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskEncryptionInfo struct{}"
	}

	return strings.Join([]string{"DiskEncryptionInfo", string(data)}, " ")
}
