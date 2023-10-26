package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePermRuleRequest Request Object
type CreatePermRuleRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	Body *CreatePermRulesRequestBody `json:"body,omitempty"`
}

func (o CreatePermRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePermRuleRequest struct{}"
	}

	return strings.Join([]string{"CreatePermRuleRequest", string(data)}, " ")
}
