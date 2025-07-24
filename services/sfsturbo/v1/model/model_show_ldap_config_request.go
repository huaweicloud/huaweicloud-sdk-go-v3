package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLdapConfigRequest Request Object
type ShowLdapConfigRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`
}

func (o ShowLdapConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLdapConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowLdapConfigRequest", string(data)}, " ")
}
