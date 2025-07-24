package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLdapConfigRequest Request Object
type CreateLdapConfigRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *CreateLdapConfigRequestBody `json:"body,omitempty"`
}

func (o CreateLdapConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLdapConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateLdapConfigRequest", string(data)}, " ")
}
