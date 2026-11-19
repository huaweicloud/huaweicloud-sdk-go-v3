package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDelegatedAccountsRequest Request Object
type ListDelegatedAccountsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 是否强制从organization同步组织信息 **约束限制**: 不涉及 **取值范围**: - true：强制从organization同步组织信息 - false：使用本地缓存的组织信息  **默认取值**: false
	Refresh *bool `json:"refresh,omitempty"`
}

func (o ListDelegatedAccountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDelegatedAccountsRequest struct{}"
	}

	return strings.Join([]string{"ListDelegatedAccountsRequest", string(data)}, " ")
}
