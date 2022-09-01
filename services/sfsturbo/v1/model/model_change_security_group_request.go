package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeSecurityGroupRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id" xml:"share_id"`

	Body *ChangeSecurityGroupRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ChangeSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupRequest", string(data)}, " ")
}
