package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainRequestBody 更新防护域名的请求
type UpdateDomainRequestBody struct {

	// 防护状态:防护中-on;未防护-off
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateDomainRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainRequestBody", string(data)}, " ")
}
