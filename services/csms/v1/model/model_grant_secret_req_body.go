package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GrantSecretReqBody 授权传入参数
type GrantSecretReqBody struct {

	// 资源id
	ResourceId string `json:"resource_id"`

	// 资源类型（SECRET、GROUP）
	Type string `json:"type"`

	// 被授权类型，（0：USER；2：GROUP）个人，群组
	GranteeType string `json:"grantee_type"`

	// 被授权id
	GranteeTargetId string `json:"grantee_target_id"`

	// 有效期截止时间
	ValidityTime *string `json:"validity_time,omitempty"`
}

func (o GrantSecretReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantSecretReqBody struct{}"
	}

	return strings.Join([]string{"GrantSecretReqBody", string(data)}, " ")
}
