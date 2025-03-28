package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAliasResponse Response Object
type CreateAliasResponse struct {

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 密钥ID
	KeyId *string `json:"key_id,omitempty"`

	// 别名
	Alias *string `json:"alias,omitempty"`

	// 别名资源定位符
	AliasUrn *string `json:"alias_urn,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAliasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAliasResponse struct{}"
	}

	return strings.Join([]string{"CreateAliasResponse", string(data)}, " ")
}
