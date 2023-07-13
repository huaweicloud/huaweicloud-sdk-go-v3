package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSecurityGroupRequest Request Object
type ChangeSecurityGroupRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *ChangeSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o ChangeSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupRequest", string(data)}, " ")
}
