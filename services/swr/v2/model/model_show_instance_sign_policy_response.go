package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceSignPolicyResponse Response Object
type ShowInstanceSignPolicyResponse struct {

	// 签名策略ID
	Id *int32 `json:"id,omitempty"`

	// 签名策略名称
	Name *string `json:"name,omitempty"`

	// 签名策略描述
	Description *string `json:"description,omitempty"`

	// 命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// 命名空间名
	NamespaceName *string `json:"namespace_name,omitempty"`

	Trigger *TriggerConfig `json:"trigger,omitempty"`

	// 创建者
	Creator *string `json:"creator,omitempty"`

	// 是否
	Enabled *bool `json:"enabled,omitempty"`

	// 作用范围规则
	ScopeRules *[]SignScopeRule `json:"scope_rules,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 加签算法，KMS的密钥算法EC_P256对应着ECDSA_SHA_256，EC_P384对应着ECDSA_SHA_384，SM2对应着SM2DSA_SM3
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 签名算法key ID
	SignatureKey *string `json:"signature_key,omitempty"`

	// 镜像签名方式
	SignatureMethod *string `json:"signature_method,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowInstanceSignPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceSignPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceSignPolicyResponse", string(data)}, " ")
}
