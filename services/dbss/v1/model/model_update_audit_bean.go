package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditBean 修改审计实例信息bean
type UpdateAuditBean struct {

	// 实例名称。只能由中文字符,英文字母,数字,下划线,中划线组成的字符串,长度小于等于64。不能为空字符串。
	Name *string `json:"name,omitempty"`

	// 实例描述信息，只能由中文字符,英文字母,数字,下划线,中划线,空格组成的字符串，长度小于等于255。可以为空字符串。
	Comment *string `json:"comment,omitempty"`
}

func (o UpdateAuditBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditBean struct{}"
	}

	return strings.Join([]string{"UpdateAuditBean", string(data)}, " ")
}
