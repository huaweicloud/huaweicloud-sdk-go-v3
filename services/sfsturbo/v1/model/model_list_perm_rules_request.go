package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermRulesRequest Request Object
type ListPermRulesRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	// 返回的权限规则个数
	Limit *int64 `json:"limit,omitempty"`

	// 返回的权限规则的偏移量
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListPermRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermRulesRequest struct{}"
	}

	return strings.Join([]string{"ListPermRulesRequest", string(data)}, " ")
}
