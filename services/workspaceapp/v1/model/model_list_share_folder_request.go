package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareFolderRequest Request Object
type ListShareFolderRequest struct {

	// 查询的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]
	Limit *int32 `json:"limit,omitempty"`

	// WKS存储ID
	StorageId string `json:"storage_id"`

	// WKS存储目录声明ID
	StorageClaimId *string `json:"storage_claim_id,omitempty"`

	// 查询名称需满足如下规则: 1. 可见字符+空格 2. 长度0~128个字符
	Path *string `json:"path,omitempty"`
}

func (o ListShareFolderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareFolderRequest struct{}"
	}

	return strings.Join([]string{"ListShareFolderRequest", string(data)}, " ")
}
