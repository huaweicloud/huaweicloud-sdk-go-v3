package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SseSpecification 用户数据静态加密设置。非必填项，默认不启用静态加密。
type SseSpecification struct {

	// 启用静态加密。
	Enable bool `bson:"enable"`

	// 加密类型，支持SSE-KMS-S和SSE-KMS-C。
	SseType *string `bson:"sse_type,omitempty"`

	// 加密算法，支持AES256_GCM。
	SseAlgorithm *string `bson:"sse_algorithm,omitempty"`

	// KMS中的用户主密钥ID。非必填项，仅当加密类型是SSE-KMS-C时需要填写该字段。
	CmkId *string `bson:"cmk_id,omitempty"`
}

func (o SseSpecification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SseSpecification struct{}"
	}

	return strings.Join([]string{"SseSpecification", string(data)}, " ")
}
