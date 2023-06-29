package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePersistentStorageReq 创建WKS存储
type CreatePersistentStorageReq struct {

	// WKS存储名称,名称需满足如下规则: 1. 名称允许可见字符或空格，不可为全空格 2. 长度1~128个字符
	Name string `json:"name"`

	StorageMetadata *Storage `json:"storage_metadata"`
}

func (o CreatePersistentStorageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePersistentStorageReq struct{}"
	}

	return strings.Join([]string{"CreatePersistentStorageReq", string(data)}, " ")
}
