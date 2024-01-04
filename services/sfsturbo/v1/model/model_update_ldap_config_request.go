package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLdapConfigRequest Request Object
type UpdateLdapConfigRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	Body *UpdateLdapConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateLdapConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLdapConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateLdapConfigRequest", string(data)}, " ")
}
