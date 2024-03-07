package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArtifactHashCode struct {

	// 哈希算法
	HashType *string `json:"hash_type,omitempty"`

	// 哈希值
	HashValue *string `json:"hash_value,omitempty"`
}

func (o ArtifactHashCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArtifactHashCode struct{}"
	}

	return strings.Join([]string{"ArtifactHashCode", string(data)}, " ")
}
