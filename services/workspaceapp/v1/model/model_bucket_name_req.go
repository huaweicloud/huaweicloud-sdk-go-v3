package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BucketNameReq 桶名称
type BucketNameReq struct {

	// 桶名称,名称需满足如下规则: 3～63个字符，数字或字母开头，支持小写字母、数字、“-”、“.”。 禁止使用IP地址。 禁止以“-”或“.”开头及结尾。 禁止两个“.”相邻（如：“my..bucket”）。 禁止“.”和“-”相邻（如：“my-.bucket”和“my.-bucket”）
	BucketName *string `json:"bucket_name,omitempty"`
}

func (o BucketNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BucketNameReq struct{}"
	}

	return strings.Join([]string{"BucketNameReq", string(data)}, " ")
}
