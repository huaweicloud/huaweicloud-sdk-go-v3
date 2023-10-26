package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermRulesRequest Request Object
type ListPermRulesRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`
}

func (o ListPermRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermRulesRequest struct{}"
	}

	return strings.Join([]string{"ListPermRulesRequest", string(data)}, " ")
}
