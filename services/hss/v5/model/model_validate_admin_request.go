package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateAdminRequest Request Object
type ValidateAdminRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`
}

func (o ValidateAdminRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateAdminRequest struct{}"
	}

	return strings.Join([]string{"ValidateAdminRequest", string(data)}, " ")
}
